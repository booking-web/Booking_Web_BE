package db

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func ConnectPostgres() (*sql.DB, error) {
	err := godotenv.Load("./internal/.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	DB_Path := os.Getenv("DB_PATH")

	db, err := sql.Open("postgres", DB_Path)
	if err != nil {
		return db, err
	}

	// err = db.Ping()

	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println("Postgresql Connected!")

	return db, nil
}
