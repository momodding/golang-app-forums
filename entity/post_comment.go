package entity

type PostComment struct {
	CommonEntity    `gorm:"embedded"`
	PostID          uint64 `gorm:"column:post_id" json:"post_id"`
	ParentCommentID uint64 `gorm:"column:parent_comment_id" json:"parent_comment_id"`
	Content         string `gorm:"column:content" json:"content"`
	PostedBy        uint64 `gorm:"column:posted_by" json:"posted_by"`
}

type PostCommentTable interface {
	TableName() string
}

func (table *PostComment) TableName() string {
	return "post_comment"
}
