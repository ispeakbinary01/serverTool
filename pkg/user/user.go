package user

import (
	"errors"
	"gopkg.in/go-playground/validator.v9"
	"log"
	"strings"
)

var roles = []string{"admin", "moderator", "user"}

// User ...
type User struct {
	ID       int    `json:"id"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,max=50,min=6"`
	Role     string `json:"role" validate:"required"`
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
		log.Printf("%s", err)
		return err
	}
	u.Role = strings.ToLower(u.Role)
	rerr := false
	for _, r := range roles {
		if u.Role == r {
			rerr = true
			break
		}
	}
	if !rerr {
		return errors.New("role not available")
	}
	return nil
}