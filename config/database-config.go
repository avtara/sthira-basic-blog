package config

import (
	"fmt"
	"log"

	"github.com/avtara/sthira-simple-blog/entity"
	"github.com/avtara/sthira-simple-blog/helper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

//SetupDatabaseConnection is creating new connection to DB
func SetupDatabaseConnection() *gorm.DB {

	dbUser := helper.GetEnv("DB_USER", "root")
	dbPass := helper.GetEnv("DB_PASS", "")
	dbHost := helper.GetEnv("DB_HOST", "127.0.0.1")
	dbPort := helper.GetEnv("DB_PORT", "5432")
	dbName := helper.GetEnv("DB_NAME", "postgres")
	dbSslMode := helper.GetEnv("DB_SSLMODE", "disable")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=Asia/Shanghai", dbHost, dbUser, dbPass, dbName, dbPort, dbSslMode)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to create a connection to database")
	}

	db.AutoMigrate(&entity.User{}, &entity.Blog{}, &entity.ContactUs{})
	return db
}

//CloseDatabaseConnection is closing connection between app and db
func CloseDatabaseConnection(db *gorm.DB) {
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalln(err)
	}
	defer sqlDB.Close()
}
