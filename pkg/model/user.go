package model

type User struct {
	Model
	Username  string    `json:"name" gorm:"column:username"`
	Password  string    `json:"addr" gorm:"column:password"`
	Role      string    `json:"birth" gorm:"column:role"`
}

func (f *User) TableName() string {
	return "user"
}

