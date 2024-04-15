package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"blog/model"
)

var gormDB *gorm.DB
var db *sql.DB

func Create_DB() {
	dsn := "root:root@tcp(localhost:3306)/blog?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}

	// Assign the gorm.DB instance to the global variable
	gormDB = db

	// Automatically migrate the schema
	err = gormDB.AutoMigrate(&model.User{})
	if err != nil {
		log.Fatal("Error migrating database:", err)
	}

	err = gormDB.AutoMigrate(&model.Blog{})
	if err != nil {
		log.Fatal("Error migrating database:", err)
	}
}

func Get_GormDB() *gorm.DB {
	return gormDB
}
