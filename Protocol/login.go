package protocol

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"errors"
	"fmt"

	"github.com/ShadowJonathan/mopher/Protocol/lib"
	"github.com/ShadowJonathan/mopher/Protocol/mojang"
)

// BUG(Think) LoginToServer doesn't support offline mode. Call it a feature?

// LoginToServer sends the necessary packets to join a server. This
// also authenticates the request with mojang for online mode connections.
// This stops before LoginSuccess (or any other preceding packets).
func (c *Conn) LoginToServer(profile mojang.Profile) (err error, success *LoginSuccess) {
	err = c.WritePacket(&Handshake{
		ProtocolVersion: lib.VarInt(c.ProtocolVersion),
		Host:            c.host,
		Port:            c.port,
		Next:            lib.VarInt(lib.Login - 1),
	})
	if err != nil {
		return
	}
	c.State = lib.Login
	if err = c.WritePacket(&LoginStart{
		Username: profile.Username,
	}); err != nil {
		return
	}

	var packet lib.MetaPacket
	if packet, err = c.ReadPacket(); err != nil {
		return
	}

	req, err, ls := checkLoginPacket(c, packet)
	if err != nil {
		if err == OFFLINE_ERR {
			return nil, ls
		}
		return err, nil
	} else {
		var p interface{}
		if p, err = x509.ParsePKIXPublicKey(req.PublicKey); err != nil {
			return
		}
		pub := p.(*rsa.PublicKey)

		key := make([]byte, 16)
		n, err := rand.Read(key)
		if n != 16 || err != nil {
			return errors.New("crypto error"), nil
		}

		sharedKey, err := rsa.EncryptPKCS1v15(rand.Reader, pub, key)
		if err != nil {
			return err, nil
		}
		verifyToken, err := rsa.EncryptPKCS1v15(rand.Reader, pub, req.VerifyToken)
		if err != nil {
			return err, nil
		}

		err = mojang.JoinServer(profile, []byte(req.ServerID), key, req.PublicKey)
		if err != nil {
			return err, nil
		}

		err = c.WritePacket(&EncryptionResponse{
			SharedSecret: sharedKey,
			VerifyToken:  verifyToken,
		})
		if err != nil {
			return err, nil
		}

		err = c.EnableEncryption(key)
	}
	return
}

var OFFLINE_ERR = errors.New("server is in offline mode which is currently unsupported")

func checkLoginPacket(c *Conn, p lib.MetaPacket) (*EncryptionRequest, error, *LoginSuccess) {
	switch p := p.(type) {
	case *EncryptionRequest:
		return p, nil, nil
	case *LoginDisconnect:
		return nil, errors.New(fmt.Sprintf("DISCONNECT: %s", p.Reason)), nil
	case *LoginSuccess:
		return nil, OFFLINE_ERR, p
	case *SetInitialCompression:
		fmt.Println("Handled compression inside checkloginpacket:", int(p.Threshold))
		c.SetCompression(int(p.Threshold))
		p2, err := c.ReadPacket()
		if err != nil {
			return nil, err, nil
		}
		return checkLoginPacket(c, p2)
	default:
		return nil, fmt.Errorf("unexpected packet %#v", p), nil
	}
}
