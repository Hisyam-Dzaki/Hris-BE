package database

import (
	"fmt"
	"log"

	"hris-BE/internal/config"
	"hris-BE/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func Connect(cfg *config.Config) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBSSLMode,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	DB = db
	log.Println("Database connected")
}

func Migrate() {
	err := DB.AutoMigrate(
		&models.Company{},
		&models.Role{},
		&models.User{},
		&models.Department{},
		&models.Location{},
		&models.Employee{},
	)
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
	log.Println("Database migrated")
}

func GetDB() *gorm.DB {
	return DB
}
