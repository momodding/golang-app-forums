package entity

import "time"

type CommonEntity struct {
	ID        uint64     `gorm:"column:id;primary_key;auto_increment" json:"id"`
	CreatedAt *time.Time `gorm:"column:created_at;autoCreateTime" json:"-"`
	UpdatedAt *time.Time `gorm:"column:updated_at;autoUpdateTime" json:"-"`
	DeletedAt *time.Time `gorm:"column:deleted_at" json:"-"`
}
