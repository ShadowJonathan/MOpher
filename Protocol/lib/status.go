package lib

import (
	"../../format"
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
