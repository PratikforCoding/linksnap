package main

import (
	"github.com/PratikforCoding/linksnap/database"
	"github.com/PratikforCoding/linksnap/routers"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {

	err := godotenv.Load(".env")
    if err != nil {
        log.Fatalf("Error loading .env file: %v", err)
    }

	// Access environment variable to confirm it's loaded
	dbConnection := os.Getenv("DB_CONNECTION_LINK")
	if dbConnection == "" {
		log.Fatal("DB_CONNECTION_LINK is not set in the environment variables")
	}
	log.Println("DB Connection Link loaded successfully")

	// Connect to the database
	database.ConnectDB()
	defer database.DisconnectDB()

	// inittialize the router
	r := routers.SetUpRouter()

	//start the server
	r.Run(":8080")
}