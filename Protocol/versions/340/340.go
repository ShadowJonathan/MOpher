//go:generate protocol_translator

package _340

import (
	"../../lib"
)

type This struct {
	t func(interface{}) (lib.Packet, error)
	b func(i interface{}) (lib.MetaPacket, error)
}

var P *This

func (*This) Packets() [4][2][lib.MaxPacketCount]func() lib.Packet {
	return packets
}

func (t *This) Translate(i interface{}) (lib.Packet, error) {
	return t.t(i)
}

func (t *This) Back(i interface{}) (lib.MetaPacket, error) {
	return t.b(i)
}

func (t *This) InitTranslate(T func(interface{}) (lib.Packet, error), B func(i interface{}) (lib.MetaPacket, error)) {
	t.t = T
	t.b = B
}

func Version() int {
	return 340
}

var packets [4][2][lib.MaxPacketCount]func() lib.Packet

func init() {
	P = &This{}
}
