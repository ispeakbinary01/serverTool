package server

// Server ...
type Server struct {
	IP int    `json:"ip"`
	OS string `json:"os"`
}

func NewServer() *Server {
	server := &Server{}
	return server
}
