package structs

// SSH ...
type SSH struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Key      int    `json:"key"`
}
