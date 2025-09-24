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
	percent := getPercent(memoryUsage, memorySize)
	if percent > 80 {
		fmt.Printf("Memory usage too high: %d%%\n", percent)
	}
}

const BytesInMegabytes = 1_048_576

func monitorDiskUsage(diskSize int, diskUsage int) {
	percent := getPercent(diskUsage, diskSize)
	if percent > 90 {
		freeSpace := diskSize - diskUsage
		fmt.Printf("Free disk space is too low: %d Mb left\n", freeSpace/BytesInMegabytes)
	}
}

const bytesInMegaBits = 1_000_000

func monitorNetworkBandwidth(networkBandwith int, networkBandwithUsage int) {
	percent := getPercent(networkBandwithUsage, networkBandwith)
	if percent > 90 {
		freeBandwidth := networkBandwith - networkBandwithUsage
		fmt.Printf("Network bandwidth usage high: %d Mbit/s available\n", freeBandwidth/bytesInMegaBits)
	}
}

func getPercent(a int, b int) int {
	return int(math.Trunc(float64(a) / float64(b) * 100))
}
