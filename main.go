package main

import (
	"log"
	"path"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(path.Join("configs", ".env"))
	if err != nil {
		log.Fatal("Errors loading .env file")
	}
	log.Println("Initialization is completed")
}
func main() {}
