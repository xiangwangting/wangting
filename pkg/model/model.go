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
}