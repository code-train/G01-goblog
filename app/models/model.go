package models

import "goblog/pkg/types"

// BaseModel struct
type BaseModel struct {
	ID uint64
}

// GetStringID models method
func (a BaseModel) GetStringID() string {
	return types.Uint64ToString(a.ID)
}
