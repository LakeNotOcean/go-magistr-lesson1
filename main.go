package main

import (
	"fmt"
	"log"
	"net/url"
	"path"
	"time"

	"github.com/LakeNotOcean/go-magistr-lesson1/config"
	"github.com/LakeNotOcean/go-magistr-lesson1/handlers"
	"github.com/go-resty/resty/v2"
	"github.com/joho/godotenv"
)

var appConfig *config.Config

func init() {
	err := godotenv.Load(path.Join("configs", "url.env"))
	if err != nil {
		log.Fatal("Errors loading .env file")
	}
	appConfig = config.NewConfig()
}

func monitorTask(client *resty.Client, url string) error {
	response, err := client.R().Get(url)
	if err != nil {
		return err
	}
	err = handlers.MonitoringHandler(response)
	return err
}
func main() {
	client := resty.New()
	metricsUrl, _ := url.Parse(appConfig.MetricUrl)
	metricsUrl.Scheme = appConfig.Scheme
	urlString := metricsUrl.String()

	errCount := 0
	for {
		err := monitorTask(client, urlString)
		if err != nil {
			errCount++
		}
		if errCount >= 3 {
			fmt.Println("Unable to fetch server statistic")
			errCount = 0
		}
		time.Sleep(1 * time.Second)
	}
}
