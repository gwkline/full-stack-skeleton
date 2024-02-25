package types

import (
	"time"

	"gorm.io/gorm"
)

type BaseModel struct {
	ID        uint `gorm:"autoIncrement;primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type User struct {
	BaseModel
	Email        string `gorm:"unique"`
	PasswordHash string
	Role         Role

	Apples []Apple
}

type Apple struct {
	BaseModel

	Variety AppleVariety

	UserID uint
	User   User
}
