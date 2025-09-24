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
		fmt.Printf("Load Average is too high: %d\n", loadAverage)
	}
}

func monitorMemoryUsage(memorySize int, memoryUsage int) {
	percent := int(math.Ceil(float64(memoryUsage) / float64(memorySize) * 100))
	if percent > 80 {
		fmt.Printf("Memory usage too high: %d%%\n", percent)
	}
}

const BYTES_IN_MEGABYTE = 1_000_000

func monitorDiskUsage(diskSize int, diskUsage int) {
	percent := int(math.Ceil(float64(diskUsage) / float64(diskSize) * 100))
	if percent > 90 {
		freeSpace := diskSize - diskUsage
		fmt.Printf("Free disk space is too low: %d Mb left\n", freeSpace/BYTES_IN_MEGABYTE)
	}
}

const BYTES_IN_MEGABITS = 125_000

func monitorNetworkBandwidth(networkBandwith int, networkBandwithUsage int) {
	percent := int(math.Ceil(float64(networkBandwithUsage) / float64(networkBandwith) * 100))
	if percent > 90 {
		freeBandwidth := networkBandwith - networkBandwithUsage
		fmt.Printf("Network bandwidth usage high: %d Mbit/s available\n", freeBandwidth/BYTES_IN_MEGABITS)
	}
}
