package server

import (
	"gopkg.in/go-playground/validator.v9"
	"log"
)

// Server ...
type Server struct {
	ID int `json:"id"validate:"required"`
	IP string    `json:"ip"validate:"required"`
	OS string `json:"os"validate:"required"`
}

func NewServer() *Server {
	server := &Server{}
	return server
}

func (s *Server) Validate() error {
	validate := validator.New()
	err := validate.Struct(s)
	if err != nil {
		log.Printf("%s", err)
		return err
	}
	return nil
}
