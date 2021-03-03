package model

import (
	"gorm.io/gorm"
	"time"
)

// Model base model definition, including fields `id`, `created_at`, `updated_at`, `deleted_at`, which could be embedded in your models
//    type User struct {
//      gorm.Model
//    }
type Model struct {
	ID        uint           `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	//id        uint `gorm:"primary_key"`
	//created_at time.Time
	//updated_at time.Time
	//deleted_at *time.Time `sql:"index"`
}