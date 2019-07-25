package ssh

import (
	"gopkg.in/go-playground/validator.v9"
	"log"
)

// SSH ...
type SSH struct {
	ID       int    `json:"id"`
	Key      string    `json:"key"validate:"required"`
	ServerID int `json:"server_id"validate:"required"`
}

// NewSSH ...
func NewSSH() *SSH {
	SSH := &SSH{}
	return SSH
}

func (ssh *SSH) Validate() error {
	validate := validator.New()
	err := validate.Struct(ssh)
	if err != nil{
		log.Printf("%s", err)
		return err
	}
	return nil
}
