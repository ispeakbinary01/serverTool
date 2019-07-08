package ssh

// SSH ...
type SSH struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Key      int    `json:"key"`
	ServerID int `json:"server_id"`
}

// NewSSH ...
func NewSSH() *SSH {
	SSH := &SSH{}
	return SSH
}
