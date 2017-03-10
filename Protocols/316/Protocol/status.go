package protocol

import (
	"fmt"
	"time"

	"github.com/ShadowJonathan/MOpher/format"
)

// StatusReply is the reply retrieved from a server when pinging
// it.
type StatusReply struct {
	Version struct {
		Name     string `json:"name"`
		Protocol int    `json:"protocol"`
	} `json:"version"`
	Players struct {
		Max    int            `json:"max"`
		Online int            `json:"online"`
		Sample []StatusPlayer `json:"sample,omitempty"`
	} `json:"players"`
	Description format.AnyComponent `json:"description"`
	Favicon     string              `json:"favicon"`
}

// StatusPlayer is one of the sample players in a StatusReply
type StatusPlayer struct {
	Name string `json:"name"`
	ID   string `json:"id"`
}

// RequestStatus starts a status request to the server and
// returns the results of the request. The connection will
// be closed after this request.
func (c *Conn) RequestStatus() (response StatusReply, ping time.Duration, err error) {
	defer c.Close()

	err = c.WritePacket(&Handshake{
		ProtocolVersion: SupportedProtocolVersion,
		Host:            c.host,
		Port:            c.port,
		Next:            VarInt(Status - 1),
	})
	if err != nil {
		return
	}
	c.State = Status
	if err = c.WritePacket(&StatusRequest{}); err != nil {
		return
	}

	// Get the reply
	var packet Packet
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
