package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"blog/database"
	"blog/routes"

)

// var SqlDB *gorm.DB

func main() {

	fmt.Println("Making sql connection")

	routes.Routes()

	database.Create_DB()

	log.Println("Database schema migrated successfully")



	fmt.Println("Server listening on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Server error:", err)
	}
}
