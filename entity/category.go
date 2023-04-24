package entity

type Category struct {
	CommonEntity CommonEntity `gorm:"embedded"`
	Name         string       `gorm:"column:name;size:100;not null" json:"name"`
	Description  string       `gorm:"column:description;text;not null;unique" json:"description"`
}

type CategoryTable interface {
	TableName() string
}

func (table *Category) TableName() string {
	return "category"
}
