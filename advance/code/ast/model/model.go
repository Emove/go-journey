package model

import (
	"time"
)

type Model struct {
	ID         int       `gorm:"primaryKey"`
	CreateTime time.Time `gorm:"type:timestamp;autoCreateTime;not null"`
	UpdateTime time.Time `gorm:"type:timestamp;autoUpdateTime;not null"`
	IsDeleted  int       `gorm:"default:2;not null"`
	changes    map[string]interface{}
}

func (model *Model) Update(filedName string, value interface{}) {
	if model.changes == nil {
		model.changes = make(map[string]interface{})
	}
	model.changes[filedName] = value
}

func (model *Model) GetChanges() map[string]interface{} {
	if model.changes == nil {
		model.changes = make(map[string]interface{})
	}
	return model.changes
}
