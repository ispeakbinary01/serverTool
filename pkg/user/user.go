package user

// User ...
type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// NewUser ...
func NewUser() *User {
	u := &User{}
	return u
}