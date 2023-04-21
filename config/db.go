package config

import (
	"fmt"
	"forum-app/helper"
	"github.com/spf13/viper"
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
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		viper.GetString("app.db.host"), viper.GetString("app.db.user"), viper.GetString("app.db.password"), viper.GetString("app.db.name"),
		viper.GetString("app.db.port"))

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
