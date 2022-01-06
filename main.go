package main

import (
	"fmt"

	"github.com/alextwothousand/batteryman/batteryman"
)

func main() {
	device, err := batteryman.GetBatteryDevice()
	if err != nil {
		panic(err)
	}
	fmt.Printf("active battery device: %s\n", device)

	threshold, err := batteryman.GetThreshold()
	if err != nil {
		panic(err)
	}

	fmt.Printf("current battery threshold: %d\n", threshold)

	capacity, err := batteryman.GetCapacity()
	if err != nil {
		panic(err)
	}

	fmt.Printf("current charge levels: %d\n", capacity)

	status, err := batteryman.GetStatus()
	if err != nil {
		panic(err)
	}

	fmt.Printf("current battery status: %s\n", status)
}
