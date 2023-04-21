package entity

import "time"

type PostComment struct {
	ID              uint64     `gorm:"column:id;primary_key;auto_increment" json:"id"`
	PostID          uint64     `gorm:"column:post_id" json:"post_id"`
	ParentCommentID uint64     `gorm:"column:parent_comment_id" json:"parent_comment_id"`
	Content         string     `gorm:"column:content" json:"content"`
	PostedBy        uint64     `gorm:"column:posted_by" json:"posted_by"`
	CreatedAt       *time.Time `gorm:"column:created_at;autoCreateTime" json:"-"`
	UpdatedAt       *time.Time `gorm:"column:updated_at;autoUpdateTime" json:"-"`
	DeletedAt       *time.Time `gorm:"column:deleted_at" json:"-"`
}

type PostCommentTable interface {
	TableName() string
}

func (table *PostComment) TableName() string {
	return "post_comment"
}
