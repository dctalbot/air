package main

import (
	"fmt"
	"log"
	"os"

	"github.com/dctalbot/air/awair"
)

func main() {
	if os.Getenv("AWAIR_TOKEN") == "" {
		log.Fatal("AWAIR_TOKEN environment variable not set")
		return
	}

	deviceID := awair.GetDeviceID()

	fmt.Println("deviceID:", deviceID)
}
