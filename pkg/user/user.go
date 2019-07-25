package user

import (
	"gopkg.in/go-playground/validator.v9"
	"log"
)

// User ...
type User struct {
	ID       int    `json:"id"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,max=50,min=6"`
	Position string `json:"position" validate:"required"`
}

// NewUser ...
func NewUser() *User {
	u := &User{}
	return u
}

func (u *User) Validate() error {
	validate := validator.New()
	err := validate.Struct(u)
	if err != nil{
		log.Printf("%s", err.Error())
		return err
	}
	return nil
}