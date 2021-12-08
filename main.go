package main

import (
	"StanCodingChallenge/app"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func main() {
	port := goDotEnvVariable("PORT")
	if port == "" {
		port = "3000" // Default port if not specified
	}
	println("Service starting on port: " + port)
	app := &app.App{}
	app.Initialize()
	app.Run(":" + port)
}
