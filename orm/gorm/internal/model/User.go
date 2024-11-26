package model

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"gorm/internal/data"
)

type User struct {
	*data.Model
}

type UserRepository struct {
}

func (User) TableName() string {
	return "user"
}

func CreateUser(ctx context.Context, db *gorm.DB, user *User) error {
	res := db.Create(user)
	return res.Error
}

func FindUser(ctx context.Context, db *gorm.DB, id uint) (*User, error) {
	user := User{}
	res := db.First(&user, id)
	if errors.Is(gorm.ErrRecordNotFound, res.Error) {
		return nil, nil
	}
	return &user, res.Error
}
