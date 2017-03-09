package mojang

// Profile contains information about the player required
// to connect to a server
type Profile struct {
	Username    string
	ID          string
	AccessToken string
}

// IsComplete returns whether the profile is enough to connect
// with.
func (p Profile) IsComplete() bool {
	return p.Username != "" && p.ID != "" && p.AccessToken != ""
}
