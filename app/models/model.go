package models

import (
	"goblog/pkg/types"
	"time"
)

// BaseModel struct
type BaseModel struct {
	ID        uint64    `gorm:"column:id;primaryKey;autoIncrement;not null"`
	CreatedAt time.Time `gorm:"column:created_at;index"`
	UpdatedAt time.Time `gorm:"column:updated_at;index"`
}

// GetStringID models method
func (a BaseModel) GetStringID() string {
	return types.Uint64ToString(a.ID)
}
