// Package protocol provides definitions for Minecraft packets as well as
// methods for reading and writing them
package protocol

import (
	"./lib"
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
	State                lib.State
	compressionThreshold int

	ProtocolVersion int
	CP              Protocol

	Logger func(read bool, packet lib.MetaPacket, id int, state lib.State)

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
			direction:            lib.Serverbound,
			host:                 host,
			port:                 uint16(port),
			compressionThreshold: -1,
		}, nil
	}
	return nil, lastErr
}

// WritePacket serializes the packet to the underlying
// connection, optionally encrypting and/or compressing
func (c *Conn) WritePacket(i interface{}) error {
	packet, err := c.CP.Translate(i)
	if err != nil {
		return err
	}
	// 15 second timeout
	c.net.SetWriteDeadline(time.Now().Add(15 * time.Second))

	buf := &bytes.Buffer{}

	// Contents of the packet (ID + Data)
	if err := lib.WriteVarInt(buf, lib.VarInt(packet.Id())); err != nil {
		return err
	}
	if err := packet.Write(buf); err != nil {
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
		extra = lib.VarIntSize(lib.VarInt(uncompessedSize))
	}

	// Write the length prefix followed by the buffer
	if err := lib.WriteVarInt(c.w, lib.VarInt(buf.Len()+extra)); err != nil {
		return err
	}

	// Write the uncompressed packet size
	if c.compressionThreshold >= 0 {
		if err := lib.WriteVarInt(c.w, lib.VarInt(uncompessedSize)); err != nil {
			return err
		}
	}

	_, err = buf.WriteTo(c.w)
	if c.Logger != nil {
		c.Logger(false, packet, packet.Id(), c.State)
	}
	return err
}

// ReadPacket deserializes a packet from the underlying
// connection, optionally decrypting and/or decompressing
func (c *Conn) ReadPacket() (lib.MetaPacket, error) {
	// 15 second timeout
	c.net.SetReadDeadline(time.Now().Add(15 * time.Second))
	return c.readPacket()
}

var (
	errNegativeLength = errors.New("invalid length: negative")
)

func (c *Conn) readPacket() (lib.MetaPacket, error) {
	// Length prefix
	size, err := lib.ReadVarInt(c.r)
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
		uncompSize, err := lib.ReadVarInt(r)
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
	id, err := lib.ReadVarInt(r)
	if err != nil {
		return nil, err
	}
	//fmt.Printf("Received packet: %02X;%d\n", id, id)
	// Direction is swapped as this is coming from the other way
	packets := c.CP.Packets()[c.State][(c.direction+1)&1]
	if id < 0 || int(id) >= len(packets) || packets[id] == nil {
		b, _ := ioutil.ReadAll(r)
		var bs string
		for i := 0; i < len(b); i++ {
			bs += MakeByte(b[i]) + "\n"
		}
		return nil, fmt.Errorf("Unknown packet %s:%02X\nBytes:\n%s", c.State, id, bs)
	}
	packet := packets[id]()
	if err := packet.Read(r); err != nil {
		p, _ := c.CP.Back(packet)
		return p, fmt.Errorf("packet(%s:%02X): %s", c.State, id, err)
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
			p, _ := c.CP.Back(packet)
			return p, fmt.Errorf("didn't finish reading packet %s:%d/0x%02X, have %d bytes left\nLeft bytes:\n%s", c.State, id, id, lb, bs)
		} else {
			p, _ := c.CP.Back(packet)
			return p, fmt.Errorf("didn't finish reading packet %s:%d/0x%02X, have %d bytes left", c.State, id, id, lb)
		}
	}
	if c.Logger != nil {
		c.Logger(true, packet, packet.Id(), c.State)
	}
	return c.CP.Back(packet)
}

func MakeByte(b byte) string {
	bs := strconv.FormatUint(uint64(uint8(b)), 2)
	if len(bs) < 8 {
		l := len(bs)
		for i := 0; i < 8-l; i++ {
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

// RequestStatus starts a status request to the server and
// returns the results of the request. The connection will
// be closed after this request.
func (c *Conn) RequestStatus() (response lib.StatusReply, ping time.Duration, err error) {
	defer c.Close()

	err = c.WritePacket(&Handshake{
		ProtocolVersion: 315,
		Host:            c.host,
		Port:            c.port,
		Next:            lib.VarInt(lib.Status - 1),
	})
	if err != nil {
		return
	}
	c.State = lib.Status
	if err = c.WritePacket(&StatusRequest{}); err != nil {
		return
	}

	// Get the reply
	var packet interface{}
	if packet, err = c.ReadPacket(); err != nil {
		return
	}

	resp, ok := packet.(*StatusResponse)
	if !ok {
		err = fmt.Errorf("unexpected packet %#v", packet)
		return
	}
	response = resp.Status

	t := time.Now()
	if err = c.WritePacket(&StatusPing{
		Time: t.UnixNano(),
	}); err != nil {
		return
	}

	// Get the pong reply
	packet, err = c.ReadPacket()
	if err != nil {
		return
	}

	_, ok = packet.(*StatusPong)
	if !ok {
		err = fmt.Errorf("unexpected packet %#v", packet)
	}
	ping = time.Now().Sub(t)
	return
}

func (c *Conn) ResolveConnectable() (bool, error, int, Protocol) {
	c.CP = defaultProtocol()

	response, _, err := c.RequestStatus()
	if err != nil {
		panic(err)
	}

	s, err := SupportedProtocol(response.Version.Protocol)

	return s, err, response.Version.Protocol, protocols[response.Version.Protocol]
}
