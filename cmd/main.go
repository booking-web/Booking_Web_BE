package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/billzayy/Booking_Web_BE/internal/routes"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("./internal/.env")
	if err != nil {
		// log.Fatal("Error loading .env file")
		fmt.Println("Error loading .env file")
	}

	PORT := os.Getenv("PORT")
	routes := routes.NewRouter()

	fmt.Printf("Server is running on port %s\n", PORT)
	http.ListenAndServe(":"+PORT, routes)
}
