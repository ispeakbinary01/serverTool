package handlers

type handler struct{}

//// Signin ...
//func (h *handler) Signin(c echo.Context) error {
//	username := c.FormValue("username")
//	stmt, err := db.Get().Prepare(userCheck)
//	if err != nil {
//		return err
//	}
//	res, err := stmt.Exec(username)
//	if err != nil {
//		return err
//	}
//	if res != nil {
//		token := jwt.New(jwt.SigningMethodHS256)
//		claims := token.Claims.(jwt.MapClaims)
//		claims["username"] = username
//		claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
//	}
//
//
//
//}



const userCheck = `
SELECT username, password FROM user WHERE username = ? LIMIT 1
`