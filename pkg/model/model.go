package model

import (
	"gorm.io/gorm"
	"time"
)

// Model base model definition, including fields `id`, `created_at`, `updated_at`, `deleted_at`, which could be embedded in your models
//    type User struct {
//      gorm.Model
//    }
//type Model struct {
//
//	Id        uint       `json:"id" gorm:"column:id"`
//	UpdatedAt time.Time `json:"update_at" gorm:"column:update_at"" description:"更新时间"`
//	CreatedAt time.Time `json:"create_at" gorm:"column:create_at" description:"创建时间"`
//	DeletedAt time.Time `json:"delete_at" gorm:"column:delete_at" description:"删除时间"`
//}

// gorm.Model 的定义
type Model struct {
	ID        uint           `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}