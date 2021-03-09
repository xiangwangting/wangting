package model

type User struct {
	Id        uint       `json:"id" gorm:"column:id"`
	Username  string    `json:"name" gorm:"column:username"`
	Password  string    `json:"password" gorm:"column:password"`
	Role      string    `json:"role" gorm:"column:role"`
	Model
}

func (u *User) TableName() string {
	return "user"
}

