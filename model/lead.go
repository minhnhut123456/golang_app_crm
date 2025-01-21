package model

import "gorm.io/gorm"

type Lead struct {
	ID   int64       					`json:"id"`
	Name string								`json:"name"`
	Company string						`json:"company"`
	DeletedAt gorm.DeletedAt	`json:"deleted_at"`
}