package server

// Server ...
type Server struct {
	ID			int `json:"id"`
	IP       int               `json:"ip"`
	OS       string            `json:"os"`
}

func NewServer() *Server {
	server := &Server{}
	return server
}