//go:generate protocol_builder $GOFILE Login Serverbound

package _315

// LoginStart is sent immeditately after switching into the login
// state. The passed username is used by the server to authenticate
// the player in online mode.
//
// This is a Minecraft packet
type LoginStart struct {
	Username string
}

// EncryptionResponse is sent as a reply to EncryptionRequest. All
// packets following this one must be encrypted with AES/CFB8
// encryption.
//
// This is a Minecraft packet
type EncryptionResponse struct {
	// The key for the AES/CFB8 cipher encrypted with the
	// public key
	SharedSecret []byte `length:"lib.VarInt"`
	// The verify token from the request encrypted with the
	// public key
	VerifyToken []byte `length:"lib.VarInt"`
}
