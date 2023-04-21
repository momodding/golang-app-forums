package entity

type Post struct {
	CommonEntity `gorm:"embedded"`
	Title        string `gorm:"column:title;size:50" json:"title"`
	Post         string `gorm:"column:post" json:"content"`
	PostedBy     uint64 `gorm:"column:posted_by" json:"posted_by"`
}

type PostTable interface {
	TableName() string
}

func (table *Post) TableName() string {
	return "post"
}
