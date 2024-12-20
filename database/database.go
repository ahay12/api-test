package database

import (
	"log"
	"time"

	"github.com/ahay12/api-test/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func InitDatabase() *gorm.DB {
	dsn := "root:ahay@tcp(localhost:3306)/api_test_u?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Filed to connect to database")
	}

	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatal("Failed to get sql.DB from gorm.DB")
	}
	// SetMaxOpenCons sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(100)
	// SetMaxIdleCons sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)
	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)

	err = DB.AutoMigrate(&model.Users{})
	err = DB.AutoMigrate(&model.Project{})
	if err != nil {
		log.Fatal("Failed to migrate database")
	}

	return DB
}
