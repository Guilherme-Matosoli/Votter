package database

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	"os"

	_ "github.com/lib/pq"
)

func Connection() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	DB_HOST := os.Getenv("DB_HOST")
	DB_PORT := os.Getenv("DB_PORT")
	DB_NAME := os.Getenv("DB-NAME")
	DB_PASS := os.Getenv("DB_PASS")
	DB_USER := os.Getenv("DB_USER")

	config := fmt.Sprintf("host=%s port=%s user=%s password=%s db_name=%s sslmode=disable",
		DB_HOST,
		DB_PORT,
		DB_USER,
		DB_PASS,
		DB_NAME,
	)

	conn, err := sql.Open("postgres", config)
	if err != nil {
		panic(err)
	}

	err = conn.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Database conneciton successfully!")
}
