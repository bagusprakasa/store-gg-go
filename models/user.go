package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name            string
	Email           string    `gorm:"unique"`
	EmailVerifiedAt time.Time `gorm:"default:null"`
	RememberToken   string
	Password        string
}
