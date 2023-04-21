package entity

import "time"

type Post struct {
	ID        uint64     `gorm:"column:id;primary_key;auto_increment" json:"id"`
	Title     string     `gorm:"column:title;size:50" json:"title"`
	Post      string     `gorm:"column:post" json:"content"`
	PostedBy  uint64     `gorm:"column:posted_by" json:"posted_by"`
	CreatedAt *time.Time `gorm:"column:created_at;autoCreateTime" json:"-"`
	UpdatedAt *time.Time `gorm:"column:updated_at;autoUpdateTime" json:"-"`
	DeletedAt *time.Time `gorm:"column:deleted_at" json:"-"`
}

type PostTable interface {
	TableName() string
}

func (table *Post) TableName() string {
	return "post"
}
