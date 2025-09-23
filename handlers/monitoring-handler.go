package handlers

import (
	"fmt"
	"math"
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
	networkBandwith
	networkBandwithUsage
)

func MonitoringHandler(response *resty.Response) error {
	responseStrings := strings.Split(string(response.Body()), ",")
	if len(responseStrings) != 7 {
		return fmt.Errorf("invalid parameters")
	}
	metrics := make([]int, len(responseStrings))
	for i, stringNumber := range responseStrings {
		metric, err := strconv.Atoi(stringNumber)
		if err != nil {
			return err
		}
		metrics[i] = metric
	}
	monitorLoadAverage(metrics[loadAverage])
	monitorMemoryUsage(metrics[memorySize], metrics[memoryUsage])
	monitorDiskUsage(metrics[diskSize], metrics[diskUsage])
	monitorNetworkBandwidth(metrics[networkBandwith], metrics[networkBandwithUsage])
	return nil
}

func monitorLoadAverage(loadAverage int) {
	if loadAverage > 30 {
		fmt.Println("Load Average is too high:", loadAverage)
	}
}

func monitorMemoryUsage(memorySize int, memoryUsage int) {
	percent := math.Ceil(float64(memoryUsage) / float64(memorySize) * 100)
	if percent > 80 {
		fmt.Println("Memory usage too high:", percent, "%")
	}
}

const BYTES_IN_MEGABYTE = 1_000_000

func monitorDiskUsage(diskSize int, diskUsage int) {
	percent := math.Ceil(float64(memoryUsage) / float64(memorySize) * 100)
	if percent > 90 {
		freeSpace := diskSize - diskUsage
		fmt.Println("Free disk space is too low:", freeSpace/BYTES_IN_MEGABYTE, "Mb left")
	}
}

const BYTES_IN_MEGABITS = 125_000

func monitorNetworkBandwidth(networkBandwith int, networkBandwithUsage int) {
	percent := math.Ceil(float64(networkBandwith) / float64(networkBandwithUsage) * 100)
	if percent > 90 {
		freeBandwidth := networkBandwith - networkBandwithUsage
		fmt.Println("Network bandwidth usage high:", freeBandwidth/BYTES_IN_MEGABITS, "Mbit/s available")
	}
}
