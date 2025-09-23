package handlers

import (
	"log"
	"strconv"
	"strings"

	"github.com/go-resty/resty/v2"
)

const (
	loadAverage = iota
	memorySize
	memoryUsage
	diskSize
	diskUsage
	networBandwith
	networkBandwithUsage
)

func MonitoringHandler(response *resty.Response) {
	responseStrings := strings.Split(string(response.Body()), ",")
	if len(responseStrings) != 7 {
		log.Fatalln("invalid data length", responseStrings)
		return
	}
	metrics := make([]int, len(responseStrings))
	for i, stringNumber := range responseStrings {
		metric, err := strconv.Atoi(stringNumber)
		if err != nil {
			log.Fatalln("invalid data format", stringNumber)
			return
		}
		metrics[i] = metric
	}
}
