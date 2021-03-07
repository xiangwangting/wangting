package model

import (
	"time"
)

type User struct {
	Id        int       `json:"id" gorm:"column:id"`
	Username  string    `json:"name" gorm:"column:username"`
	Password  string    `json:"addr" gorm:"column:password"`
	Role      string    `json:"birth" gorm:"column:role"`
	UpdatedAt time.Time `json:"update_at" gorm:"column:update_at"" description:"更新时间"`
	CreatedAt time.Time `json:"create_at" gorm:"column:create_at" description:"创建时间"`
	DeletedAt time.Time `json:"delete_at" gorm:"column:delete_at" description:"删除时间"`
}

func (f *User) TableName() string {
	return "user"
}

