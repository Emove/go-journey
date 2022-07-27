package model

import (
	"context"
	"gorm.io/gorm"
)

type User struct {
	ID uint `gorm:"primaryKey"`
}

func (User) TableName() string {
	return "user"
}

func CreateUser(ctx *context.Context, db *gorm.DB, user *User) error {
	res := db.Create(user)
	return res.Error
}
