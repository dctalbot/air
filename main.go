package main

import (
	"log"
	"os"

	"github.com/dctalbot/air/awair"
)

func main() {
	if os.Getenv("AWAIR_TOKEN") == "" {
		log.Fatal("AWAIR_TOKEN environment variable not set")
		return
	}

	devices := awair.GetDevices()
	if len(devices) == 0 {
		log.Fatal("No devices found")
		return
	}

	latest := awair.GetLatest(devices[0])

	log.Printf("Latest: %+v", latest)

}
