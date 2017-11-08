// Generated by protocol_builder
// Do not edit

package _335

import (
	"github.com/ShadowJonathan/mopher/Protocol/lib"
	"io"
)

func (s *StatusRequest) Id() int { return 0 }
func (s *StatusRequest) Write(ww io.Writer) (err error) {
	return
}
func (s *StatusRequest) Read(rr io.Reader) (err error) {
	return
}

func (s *StatusPing) Id() int { return 1 }
func (s *StatusPing) Write(ww io.Writer) (err error) {
	var tmp [8]byte
	tmp[0] = byte(s.Time >> 56)
	tmp[1] = byte(s.Time >> 48)
	tmp[2] = byte(s.Time >> 40)
	tmp[3] = byte(s.Time >> 32)
	tmp[4] = byte(s.Time >> 24)
	tmp[5] = byte(s.Time >> 16)
	tmp[6] = byte(s.Time >> 8)
	tmp[7] = byte(s.Time >> 0)
	if _, err = ww.Write(tmp[:8]); err != nil {
		return
	}
	return
}
func (s *StatusPing) Read(rr io.Reader) (err error) {
	var tmp [8]byte
	if _, err = rr.Read(tmp[:8]); err != nil {
		return
	}
	s.Time = int64((uint64(tmp[7]) << 0) | (uint64(tmp[6]) << 8) | (uint64(tmp[5]) << 16) | (uint64(tmp[4]) << 24) | (uint64(tmp[3]) << 32) | (uint64(tmp[2]) << 40) | (uint64(tmp[1]) << 48) | (uint64(tmp[0]) << 56))
	return
}

func init() {
	packets[lib.Status][lib.Serverbound][0] = func() lib.Packet { return &StatusRequest{} }
	packets[lib.Status][lib.Serverbound][1] = func() lib.Packet { return &StatusPing{} }
}
