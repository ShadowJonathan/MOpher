// Generated by protocol_builder
// Do not edit

package _340

import (
	"github.com/ShadowJonathan/mopher/Protocol/lib"
	"encoding/json"
	"fmt"
	"io"
	"math"
)

func (l *LoginDisconnect) Id() int { return 0 }
func (l *LoginDisconnect) Write(ww io.Writer) (err error) {
	var tmp0 []byte
	if tmp0, err = json.Marshal(&l.Reason); err != nil {
		return
	}
	tmp1 := string(tmp0)
	if err = lib.WriteString(ww, tmp1); err != nil {
		return
	}
	return
}
func (l *LoginDisconnect) Read(rr io.Reader) (err error) {
	var tmp0 string
	if tmp0, err = lib.ReadString(rr); err != nil {
		return err
	}
	if err = json.Unmarshal([]byte(tmp0), &l.Reason); err != nil {
		return
	}
	return
}

func (e *EncryptionRequest) Id() int { return 1 }
func (e *EncryptionRequest) Write(ww io.Writer) (err error) {
	if err = lib.WriteString(ww, e.ServerID); err != nil {
		return
	}
	if err = lib.WriteVarInt(ww, lib.VarInt(len(e.PublicKey))); err != nil {
		return
	}
	if _, err = ww.Write(e.PublicKey); err != nil {
		return
	}
	if err = lib.WriteVarInt(ww, lib.VarInt(len(e.VerifyToken))); err != nil {
		return
	}
	if _, err = ww.Write(e.VerifyToken); err != nil {
		return
	}
	return
}
func (e *EncryptionRequest) Read(rr io.Reader) (err error) {
	if e.ServerID, err = lib.ReadString(rr); err != nil {
		return
	}
	var tmp0 lib.VarInt
	if tmp0, err = lib.ReadVarInt(rr); err != nil {
		return
	}
	if tmp0 > math.MaxInt16 {
		return fmt.Errorf("array larger than max value: %d > %d", tmp0, math.MaxInt16)
	}
	if tmp0 < 0 {
		return fmt.Errorf("negative array size: %d < 0", tmp0)
	}
	e.PublicKey = make([]byte, tmp0)
	if _, err = rr.Read(e.PublicKey); err != nil {
		return
	}
	var tmp1 lib.VarInt
	if tmp1, err = lib.ReadVarInt(rr); err != nil {
		return
	}
	if tmp1 > math.MaxInt16 {
		return fmt.Errorf("array larger than max value: %d > %d", tmp1, math.MaxInt16)
	}
	if tmp1 < 0 {
		return fmt.Errorf("negative array size: %d < 0", tmp1)
	}
	e.VerifyToken = make([]byte, tmp1)
	if _, err = rr.Read(e.VerifyToken); err != nil {
		return
	}
	return
}

func (l *LoginSuccess) Id() int { return 2 }
func (l *LoginSuccess) Write(ww io.Writer) (err error) {
	if err = lib.WriteString(ww, l.UUID); err != nil {
		return
	}
	if err = lib.WriteString(ww, l.Username); err != nil {
		return
	}
	return
}
func (l *LoginSuccess) Read(rr io.Reader) (err error) {
	if l.UUID, err = lib.ReadString(rr); err != nil {
		return
	}
	if l.Username, err = lib.ReadString(rr); err != nil {
		return
	}
	return
}

func (s *SetInitialCompression) Id() int { return 3 }
func (s *SetInitialCompression) Write(ww io.Writer) (err error) {
	if err = lib.WriteVarInt(ww, s.Threshold); err != nil {
		return
	}
	return
}
func (s *SetInitialCompression) Read(rr io.Reader) (err error) {
	if s.Threshold, err = lib.ReadVarInt(rr); err != nil {
		return
	}
	return
}

func init() {
	packets[lib.Login][lib.Clientbound][0] = func() lib.Packet { return &LoginDisconnect{} }
	packets[lib.Login][lib.Clientbound][1] = func() lib.Packet { return &EncryptionRequest{} }
	packets[lib.Login][lib.Clientbound][2] = func() lib.Packet { return &LoginSuccess{} }
	packets[lib.Login][lib.Clientbound][3] = func() lib.Packet { return &SetInitialCompression{} }
}
