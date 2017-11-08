package protocol

import (
	"github.com/ShadowJonathan/mopher/lib"
	"github.com/ShadowJonathan/mopher/versions/315"
	"github.com/ShadowJonathan/mopher/versions/335"
	"github.com/ShadowJonathan/mopher/versions/338"
	"github.com/ShadowJonathan/mopher/versions/340"
	"fmt"
)

type Protocol interface {
	Packets() [4][2][lib.MaxPacketCount]func() lib.Packet
	Translate(interface{}) (lib.Packet, error)
	Back(i interface{}) (lib.MetaPacket, error)

	InitTranslate(func(interface{}) (lib.Packet, error), func(i interface{}) (lib.MetaPacket, error))
}

var supported []int
var protocols map[int]Protocol

func init() {
	protocols = make(map[int]Protocol)

	// 1.11.0
	_315.P.InitTranslate(Translate_315, Back_315)
	protocols[_335.Version()] = _315.P

	// 1.12.0
	_335.P.InitTranslate(Translate_335, Back_335)
	protocols[_335.Version()] = _335.P

	// 1.12.1
	_338.P.InitTranslate(Translate_338, Back_338)
	protocols[_338.Version()] = _338.P

	// 1.12.2
	_340.P.InitTranslate(Translate_340, Back_340)
	protocols[_340.Version()] = _340.P

	for i := range protocols {
		supported = append(supported, i)
	}
}

func SupportedProtocol(p int) (bool, error) {
	for _, i := range supported {
		if i == p {
			return true, nil
		}
	}
	return false, fmt.Errorf("unsupported version: %d, supported: %s", p, supported)
}

func defaultProtocol() Protocol {
	return protocols[_335.Version()]
}
