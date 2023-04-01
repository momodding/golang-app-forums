package entity

import "time"

type Category struct {
	ID          uint64     `gorm:"column:id;primary_key;auto_increment" json:"id"`
	Name        string     `gorm:"column:name;size:100;not null" json:"name"`
	Description string     `gorm:"column:description;text;not null;unique" json:"description"`
	CreatedAt   *time.Time `gorm:"column:created_at;autoCreateTime" json:"-"`
	UpdatedAt   *time.Time `gorm:"column:updated_at;autoUpdateTime" json:"-"`
	DeletedAt   *time.Time `gorm:"column:deleted_at" json:"-"`
}

type CategoryTable interface {
	TableName() string
}

func (table *Category) TableName() string {
	return "category"
}
