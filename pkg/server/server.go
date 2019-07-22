package server

// Server ...
type Server struct {
	IP string    `json:"ip"`
	OS string `json:"os"`
}

func NewServer() *Server {
	server := &Server{}
	return server
}
