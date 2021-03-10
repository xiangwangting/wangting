package model

type User struct {
	Username  string    `json:"name" gorm:"column:username"`
	Password  string    `json:"password" gorm:"column:password"`
	Role      string    `json:"role" gorm:"column:role"`
	Model
}

func (u *User) TableName() string {
	return "user"
}

