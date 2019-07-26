package software

import "gopkg.in/go-playground/validator.v9"

// Software ...
type Software struct {
	ID      int    `json:"id"validate:"required"`
	Name    string `json:"name"validate:"required,alphanum"`
	Version string    `json:"version"validate:"required,alphanum"`
	ServerID  int    `json:"server_id"validate:"required,alphanum"`
}

// NewSoftware ...
func NewSoftware() *Software {
	software := &Software{}
	return software
}

func (sw *Software) Validate() error {
	validate := validator.New()
	err := validate.Struct(sw)
	if err != nil{
		return err
	}
	return nil
}
