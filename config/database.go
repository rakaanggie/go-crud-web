package config

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectDB() {
	db, err := sql.Open("mysql", "admin:12345678@tcp(database-1.clqqyhrtzbj3.us-east-1.rds.amazonaws.com:3306)/db?parseTime=true")
	if err != nil {
		panic(err)
	}

	DB = db

	log.Println("Database connected")
}
