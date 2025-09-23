package main

import (
	"log"
	"path"

	"github.com/LakeNotOcean/go-magistr-lesson1/config"
	"github.com/LakeNotOcean/go-magistr-lesson1/handlers"
	"github.com/go-resty/resty/v2"
	"github.com/joho/godotenv"
)

var appConfig *config.Config

func init() {
	err := godotenv.Load(path.Join("configs", ".env"))
	if err != nil {
		log.Fatal("Errors loading .env file")
	}
	appConfig = config.NewConfig()
	log.Println("Initialization is completed")
}
func main() {
	client := resty.New()
	response, err := client.R().Get(appConfig.MetricUrl)
	if err != nil {
		log.Fatalln(err)
	}
	handlers.MonitoringHandler(response)
}
