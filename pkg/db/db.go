package db

import (
	"fmt"
	"log"
	"time"

	"github.com/pndwrzk/taskhub-service/config"
	"github.com/pndwrzk/taskhub-service/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func Connect() {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		config.App.DBHost,
		config.App.DBUser,
		config.App.DBPassword,
		config.App.DBName,
		config.App.DBPort,
	)

	newLogger := logger.New(
		log.New(log.Writer(), "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: true,
			Colorful:                  true,
		},
	)

	gdb, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic("failed to connect database: " + err.Error())
	}

	sqlDB, err := gdb.DB()
	if err != nil {
		panic(err)
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	err = gdb.AutoMigrate(
		&model.User{},
		&model.Task{},
	)
	if err != nil {
		panic("failed to migrate: " + err.Error())
	}

	DB = gdb
	log.Println("Database connected & migrated successfully")
}
