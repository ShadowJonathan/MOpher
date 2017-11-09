package lib

import (
	"errors"
	"io"
	"math"

	"github.com/ShadowJonathan/mopher/encoding/nbt"
)

const VarPart = uint32(0x7F)
const VarPartLong = uint64(0x7F)

var (
	// ErrVarIntTooLarge is returned when a read varint was too large
	// (more than 5 bytes)
	ErrVarIntTooLarge = errors.New("VarInt too large")
	// ErrVarLongTooLarge is returned when a read varint was too large
	// (more than 10 bytes)
	ErrVarLongTooLarge = errors.New("VarLong too large")
)

func VarIntSize(i VarInt) int {
	size := 0
	ui := uint32(i)
	for {
		size++
		if (ui & ^VarPart) == 0 {
			return size
		}
		ui >>= 7
	}

}

// WriteVarInt encodes the passed VarInt into the writer.
func WriteVarInt(w io.Writer, i VarInt) error {
	ui := uint32(i)
	for {
		if (ui & ^VarPart) == 0 {
			err := WriteByte(w, byte(ui))
			return err
		}
		err := WriteByte(w, byte((ui&VarPart)|0x80))
		if err != nil {
			return err
		}
		ui >>= 7
	}
}

// ReadVarInt reads a VarInt encoded integer from the reader.
func ReadVarInt(r io.Reader) (VarInt, error) {
	var size uint
	var val uint32
	for {
		b, err := ReadByte(r)
		if err != nil {
			return VarInt(val), err
		}

		val |= (uint32(b) & VarPart) << (size * 7)
		size++
		if size > 5 {
			return VarInt(val), ErrVarIntTooLarge
		}

		if (b & 0x80) == 0 {
			break
		}
	}
	return VarInt(val), nil
}

// WriteVarLong encodes the passed VarLong into the writer.
func WriteVarLong(w io.Writer, i VarLong) error {
	ui := uint64(i)
	for {
		if (ui & ^VarPartLong) == 0 {
			err := WriteByte(w, byte(ui))
			return err
		}
		err := WriteByte(w, byte((ui&VarPartLong)|0x80))
		if err != nil {
			return err
		}
		ui >>= 7
	}
}

// ReadVarLong reads a VarLong encoded 64 bit integer from the reader.
func ReadVarLong(r io.Reader) (VarLong, error) {
	var size uint
	var val uint64
	for {
		b, err := ReadByte(r)
		if err != nil {
			return VarLong(val), err
		}

		val |= (uint64(b) & VarPartLong) << (size * 7)
		size++
		if size > 10 {
			return VarLong(val), ErrVarLongTooLarge
		}

		if (b & 0x80) == 0 {
			break
		}
	}
	return VarLong(val), nil
}

// WriteString writes a VarInt prefixed utf-8 string to the
// writer.
func WriteString(w io.Writer, str string) error {
	b := []byte(str)
	err := WriteVarInt(w, VarInt(len(b)))
	if err != nil {
		return err
	}
	_, err = w.Write(b)
	return err
}

// ReadString reads a VarInt prefixed utf-8 string to the
// reader.
func ReadString(r io.Reader) (string, error) {
	l, err := ReadVarInt(r)
	if err != nil {
		return "", nil
	}
	if l < 0 || l > math.MaxInt16 {
		return "", errors.New("string length out of bounds")
	}
	buf := make([]byte, int(l))
	_, err = io.ReadFull(r, buf)
	return string(buf), err
}

// WriteBool writes a bool to the writer as a single byte.
func WriteBool(w io.Writer, b bool) error {
	if b {
		return WriteByte(w, 1)
	}
	return WriteByte(w, 0)
}

// ReadBool reads a single byte from the reader as a bool.
func ReadBool(r io.Reader) (bool, error) {
	b, err := ReadByte(r)
	if b == 0 {
		return false, err
	}
	return true, err
}

// WriteByte writes a single byte to the writer. If the
// Writer is a ByteWriter then that will be used instead.
func WriteByte(w io.Writer, b byte) error {
	if bw, ok := w.(io.ByteWriter); ok {
		return bw.WriteByte(b)
	}
	var buf [1]byte
	buf[0] = b
	_, err := w.Write(buf[:1])
	return err
}

// ReadByte reads a single byte from the Reader. If the
// Reader is a ByteReader then that will be used instead.
func ReadByte(r io.Reader) (byte, error) {
	if br, ok := r.(io.ByteReader); ok {
		return br.ReadByte()
	}
	var buf [1]byte
	_, err := r.Read(buf[:1])
	return buf[0], err
}

// ReadNBT reads an nbt tag from the reader.
// Returns nil if there is no tag.
func ReadNBT(r io.Reader) (*nbt.Compound, error) {
	b, err := ReadByte(r)
	if err != nil || b == 0 { // 0 == No tag
		return nil, err
	}
	n := nbt.NewCompound()
	err = n.Deserialize(r)
	return n, err
}

// WriteNBT writes an nbt tag to the wrtier.
// nil can be used to specify that there isn't a tag.
func WriteNBT(w io.Writer, n *nbt.Compound) error {
	if n == nil {
		return WriteByte(w, 0)
	}
	WriteByte(w, byte(nbt.TagCompound))
	return n.Serialize(w)
}
