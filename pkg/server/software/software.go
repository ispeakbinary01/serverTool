package software

// Software ...
type Software struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Version int    `json:"version"`
}

// NewSoftware ...
func NewSoftware() *Software {
	software := &Software{}
	return software
}
