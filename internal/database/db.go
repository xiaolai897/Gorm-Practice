package database

import (
	"log"
	"os"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

func init() {
	var err error
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Error,
			IgnoreRecordNotFoundError: true,
			ParameterizedQueries:      false,
			Colorful:                  false,
		},
	)
	db, err = gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic("数据库连接失败, error=" + err.Error())
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetMaxIdleConns(20)
}

func GetDB() *gorm.DB {
	return db
}
