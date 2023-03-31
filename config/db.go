package config

import (
	"forum-app/helper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBSession struct {
	DB *gorm.DB
}

func NewDbSession() *DBSession {
	dsn := "host=localhost user=momodding password=mache123 dbname=gorm port=5432 sslmode=disable TimeZone=Asia/Jakerta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	helper.PanicIfError(err)
	defer func(db *gorm.DB) {
		sql, err := db.DB()
		helper.PanicIfError(err)
		sql.Close()
	}(db)

	return &DBSession{
		db,
	}
}
