package config

import (
	"forum-app/helper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

type DBSession struct {
	DB *gorm.DB
}

func NewDbSession() *gorm.DB {
	dsn := "host=winhost user=momodding password=mache123 dbname=forum-app port=5432 sslmode=disable TimeZone=Asia/Jakarta"

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,        // Disable color
		},
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: newLogger, DisableForeignKeyConstraintWhenMigrating: true})
	helper.PanicIfError(err)
	//defer func(db *gorm.DB) {
	//	sql, err := db.DB()
	//	helper.PanicIfError(err)
	//	sql.Close()
	//}(db)

	return db
}
