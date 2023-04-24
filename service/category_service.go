package service

import (
	"forum-app/entity"
	"forum-app/helper"
	"gorm.io/gorm"
	"log"
)

type CategoryService interface {
	Save(category entity.Category) entity.Category
	FindAll() []entity.Category
}

type CategoryServiceImpl struct {
	db *gorm.DB
}

func NewCategoryService(db *gorm.DB) *CategoryServiceImpl {
	return &CategoryServiceImpl{db: db}
}

func (s *CategoryServiceImpl) Save(category entity.Category) entity.Category {
	tx := s.db.Model(&entity.Category{}).Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			log.Fatal(r)
		}
	}()

	err := tx.Create(&category)

	if err.Error != nil {
		helper.PanicIfError(err.Error)
	}

	tx.Commit()

	return category
}

func (s *CategoryServiceImpl) FindAll() []entity.Category {
	var categories []entity.Category
	s.db.Model(&entity.Category{}).Find(&categories)

	return categories
}
