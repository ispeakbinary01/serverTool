package user

import "gopkg.in/go-playground/validator.v9"

// User ...
type User struct {
	ID       int    `json:"id"`
	Username string `json:"username" validate:"required,max=10,min=5"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,max=20,min=6"`
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
		return err
	}
	return nil
}