package data

import "time"

type Model struct {
	ID         int       `gorm:"primaryKey"`
	CreateTime time.Time `gorm:"type:timestamp;autoCreateTime;not null"`
	UpdateTime time.Time `gorm:"type:timestamp;autoUpdateTime;not null"`
	IsDeleted  int       `gorm:"default:2;not null"`
	changes    map[string]interface{}
}
