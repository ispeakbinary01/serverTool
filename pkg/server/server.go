package server

import (
	"github.com/ispeakbinary01/serverTool/pkg/server/software"
	"github.com/ispeakbinary01/serverTool/pkg/server/ssh"
)

// Server ...
type Server struct {
	IP       int               `json:"id"`
	OS       string            `json:"os"`
	Software software.Software `json:"software"`
	SSH      ssh.SSH           `json:"ssh"`
}

func NewServer() *Server {
	server := &Server{}
	return server
}