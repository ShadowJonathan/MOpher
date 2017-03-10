// Package protocol provides definitions for Minecraft packets as well as
// methods for reading and writing them
package protocol

import (
	"bytes"
	"compress/zlib"
	"crypto/aes"
	"crypto/cipher"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"strconv"
	"strings"
	"time"
)

// Conn is a connection from or to a Minecraft client.
//
// The Minecraft protocol as multiple states that it
// switches between during login/status pinging, the
// state may be set using the State field.
type Conn struct {
	r                    io.Reader
	w                    io.Writer
	net                  net.Conn
	direction            int
	State                State
	compressionThreshold int

	Logger func(read bool, packet Packet)

	host string
	port uint16

	zlibReader io.ReadCloser
	zlibWriter *zlib.Writer
}

// Dial creates a connection to a Minecraft server at
// the passed address. The address is in same format as
// the vanilla client takes it
//     host:port
//     // or
//     host
// If the port isn't provided then a SRV lookup is
// performed and if successful it will continue
// connecting using the returned address. If the lookup
// fails then the port is assumed to be 25565.
func Dial(address string) (*Conn, error) {
	var toTry []string
	if !strings.ContainsRune(address, ':') {
		// Attempt a srv lookup first (like vanilla)
		_, srvs, err := net.LookupSRV("minecraft", "tcp", address)
		if err == nil && len(srvs) > 0 {
			for _, srv := range srvs {
				toTry = append(toTry, fmt.Sprintf("%s:%d", srv.Target, srv.Port))
			}
		}
		toTry = append(toTry, address+":25565")
	} else {
		toTry = append(toTry, address)
	}
	lastErr := errors.New("Unable to connect to server")
	for _, address := range toTry {
		host, portStr, err := net.SplitHostPort(address)
		if err != nil {
			lastErr = err
			continue
		}
		port, err := strconv.Atoi(portStr)
		if err != nil {
			lastErr = err
			continue
		}
		c, err := net.DialTimeout("tcp", address, 10*time.Second)
		if err != nil {
			lastErr = err
			continue
		}
		return &Conn{
			r:                    c,
			w:                    c,
			net:                  c,
			direction:            serverbound,
			host:                 host,
			port:                 uint16(port),
			compressionThreshold: -1,
		}, nil
	}
	return nil, lastErr
}

// WritePacket serializes the packet to the underlying
// connection, optionally encrypting and/or compressing
func (c *Conn) WritePacket(packet Packet) error {
	// 15 second timeout
	c.net.SetWriteDeadline(time.Now().Add(15 * time.Second))

	buf := &bytes.Buffer{}

	// Contents of the packet (ID + Data)
	if err := WriteVarInt(buf, VarInt(packet.id())); err != nil {
		return err
	}
	if err := packet.write(buf); err != nil {
		return err
	}

	uncompessedSize := 0
	extra := 0
	// Only compress if compression is enabled and the packet is large enough
	if c.compressionThreshold >= 0 && buf.Len() > c.compressionThreshold {
		var err error
		nBuf := &bytes.Buffer{}
		if c.zlibWriter == nil {
			c.zlibWriter, _ = zlib.NewWriterLevel(nBuf, zlib.BestSpeed)
		} else {
			c.zlibWriter.Reset(nBuf)
		}
		uncompessedSize = buf.Len()

		if _, err = buf.WriteTo(c.zlibWriter); err != nil {
			return err
		}
		if err = c.zlibWriter.Close(); err != nil {
			return err
		}
		buf = nBuf
	}

	// Account for the compression header if enabled
	if c.compressionThreshold >= 0 {
		extra = varIntSize(VarInt(uncompessedSize))
	}

	// Write the length prefix followed by the buffer
	if err := WriteVarInt(c.w, VarInt(buf.Len()+extra)); err != nil {
		return err
	}

	// Write the uncompressed packet size
	if c.compressionThreshold >= 0 {
		if err := WriteVarInt(c.w, VarInt(uncompessedSize)); err != nil {
			return err
		}
	}

	_, err := buf.WriteTo(c.w)
	if c.Logger != nil {
		c.Logger(false, packet)
	}
	return err
}

// ReadPacket deserializes a packet from the underlying
// connection, optionally decrypting and/or decompressing
func (c *Conn) ReadPacket() (Packet, error) {
	// 15 second timeout
	c.net.SetReadDeadline(time.Now().Add(15 * time.Second))
	return c.readPacket()
}

var (
	errNegativeLength = errors.New("invalid length: negative")
)

func (c *Conn) readPacket() (Packet, error) {
	// Length prefix
	size, err := ReadVarInt(c.r)
	if err != nil {
		return nil, err
	}
	if size < 0 {
		return nil, errNegativeLength
	}
	buf := make([]byte, size)
	if _, err := io.ReadFull(c.r, buf); err != nil {
		return nil, err
	}

	var r *bytes.Reader
	r = bytes.NewReader(buf)

	// If compression is enabled then we may need to decompress the packet
	if c.compressionThreshold >= 0 {
		// With compression enabled an extra length prefix is added
		// which is the length of the packet when uncompressed.
		uncompSize, err := ReadVarInt(r)
		if err != nil {
			return nil, err
		}
		// A uncompressed size of 0 means the packet wasn't compressed
		// and when can continue normally.
		if uncompSize != 0 {
			// Reuse the old reader to save on allocations
			if c.zlibReader == nil {
				c.zlibReader, err = zlib.NewReader(r)
				if err != nil {
					return nil, err
				}
			} else {
				err = c.zlibReader.(zlib.Resetter).Reset(r, nil)
				if err != nil {
					return nil, err
				}
			}

			// Read the whole packet at once instead of in tiny steps
			data := make([]byte, uncompSize)
			_, err := io.ReadFull(c.zlibReader, data)
			if err != nil {
				return nil, err
			}
			r = bytes.NewReader(data)
		}
	}

	// Packet ID
	id, err := ReadVarInt(r)
	if err != nil {
		return nil, err
	}
	//fmt.Printf("Received packet: %02X;%d\n", id, id)
	// Direction is swapped as this is coming from the other way
	packets := packetCreator[c.State][(c.direction+1)&1]
	if id < 0 || int(id) >= len(packets) || packets[id] == nil {
		b, _ := ioutil.ReadAll(r)
		var bs string
		for i := 0; i < len(b); i++ {
			bs += MakeByte(b[i]) + "\n"
		}
		return nil, fmt.Errorf("Unknown packet %s:%02X\nBytes:\n%s", c.State, id, bs)
	}
	packet := packets[id]()
	if err := packet.read(r); err != nil {
		return packet, fmt.Errorf("packet(%s:%02X): %s", c.State, id, err)
	}
	// If we haven't fully read the whole buffer then something went wrong.
	// Mostly likely our packet definitions are out of date or incorrect
	if r.Len() > 0 {
		lb := r.Len()

		var bs string
		if lb < 100 {
			b, _ := ioutil.ReadAll(r)
			for i := 0; i < lb; i++ {
				bs += MakeByte(b[i]) + "\n"
			}
			return packet, fmt.Errorf("Didn't finish reading packet %s:%d/0x%02X, have %d bytes left\nLeft bytes:\n%s", c.State, id, id, lb, bs)
		} else {
			return packet, fmt.Errorf("Didn't finish reading packet %s:%d/0x%02X, have %d bytes left", c.State, id, id, lb)
		}
	}
	if c.Logger != nil {
		c.Logger(true, packet)
	}
	return packet, nil
}

func MakeByte(b byte) string {
	bs := strconv.FormatUint(uint64(uint8(b)), 2)
	if len(bs) < 8 {
		l := len(bs)
		for i:=0;i<8-l;i++ {
			bs = "0" + bs
		}
	}
	return bs
}

// EnableEncryption enables cfb8 encryption on the protocol using the passed
// key.
func (c *Conn) EnableEncryption(key []byte) error {
	cip, err := aes.NewCipher(key)
	if err != nil {
		return err
	}

	c.r = cipher.StreamReader{
		R: c.net,
		S: newCFB8(cip, key, true),
	}

	c.w = cipher.StreamWriter{
		W: c.net,
		S: newCFB8(cip, key, false),
	}
	return nil
}

// SetCompression changes the threshold at which packets are compressed.
func (c *Conn) SetCompression(threshold int) {
	c.compressionThreshold = threshold
}

// Close closes the underlying connection
func (c *Conn) Close() error {
	if c.net == nil {
		return nil
	}
	return c.net.Close()
}
