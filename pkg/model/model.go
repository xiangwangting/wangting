package model

import "time"

//所有表必须有的公共字段
// Model base model definition, including fields `id`, `created_at`, `updated_at`, `deleted_at`, which could be embedded in your models
//    type User struct {
//      gorm.Model
//    }
type Model struct {
	Id        uint      `json:"id" gorm:"column:id"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at" description:"创建时间"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at"" description:"更新时间"`
	DeletedAt time.Time `json:"deleted_at" gorm:"column:deleted_at" description:"删除时间"`
}
