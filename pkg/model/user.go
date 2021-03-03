package model

//用户mode
type User struct {
	Model
	Username string
	Password string
	Role     string
}
