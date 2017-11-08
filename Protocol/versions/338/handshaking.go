//go:generate protocol_builder $GOFILE Handshaking Serverbound

package _338

import (
	"github.com/ShadowJonathan/mopher/Protocol/lib"
)

// Handshake is the first packet sent in the protocol.
// Its used for deciding if the request is a client
// is requesting status information about the server
// (MOTD, players etc) or trying to login to the server.
//
// The host and port fields are not used by the vanilla
// server but are there for virtual server hosting to
// be able to redirect a client to a target server with
// a single address + port.
//
// Some modified servers/proxies use the handshake field
// differently, packing information into the field other
// than the hostname due to the protocol not providing
// any system for custom information to be transfered
// by the client to the server until after login.
//
// This is a Minecraft packet
type Handshake struct {
	// The protocol version of the connecting client
	ProtocolVersion lib.VarInt
	// The hostname the client connected to
	Host string
	// The port the client connected to
	Port uint16
	// The next protocol state the client wants
	Next lib.VarInt
}
