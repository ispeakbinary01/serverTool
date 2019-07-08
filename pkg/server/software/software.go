package software

// Software ...
type Software struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Version int    `json:"version"`
	ServerID  int    `json:"server_id"`
}

// NewSoftware ...
func NewSoftware() *Software {
	software := &Software{}
	return software
}
