//go:generate protocol_builder $GOFILE Status Serverbound

package _340

// StatusRequest is sent by the client instantly after
// switching to the Status protocol state and is used
// to signal the server to send a StatusResponse to the
// client
//
// This is a Minecraft packet
type StatusRequest struct {
}

// StatusPing is sent by the client after recieving a
// StatusResponse. The client uses the time from sending
// the ping until the time of recieving a pong to measure
// the latency between the client and the server.
//
// This is a Minecraft packet
type StatusPing struct {
	// The time when the ping was sent
	Time int64
}
