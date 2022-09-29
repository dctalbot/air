package awair

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

const BaseURL = "https://developer-apis.awair.is"
const DevicesURL = BaseURL + "/v1/users/self/devices"

type Device struct {
	Name         string
	DeviceID     int
	DeviceType   string
	DeviceUUID   string
	Latitude     float32
	Longitude    float32
	Preference   string
	LocationName string
	RoomType     string
	SpaceType    string
	MacAddress   string
	Timezone     string
}

type DevicesResponse struct {
	Devices []Device
}

type DevicesURLConfig struct {
	DeviceType string
	DeviceID   string
	Fahrenheit bool
}

func MakeDevicesURL(cfg DevicesURLConfig) string {
	return BaseURL + "/v1/users/self/devices/{{device_type}}/{{device_id}}/air-data/latest?fahrenheit={{fahrenheit}}"
}

func GetDeviceID() int {
	httpClient := &http.Client{Timeout: 10 * time.Second}

	req, err := http.NewRequest("GET", DevicesURL, nil)
	req.Header.Add("Authorization", "Bearer "+os.Getenv("AWAIR_TOKEN"))
	if err != nil {
		log.Println(err)
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}

	var respObj DevicesResponse
	err = json.Unmarshal(data, &respObj)
	if err != nil {
		log.Println(err)
	}

	if len(respObj.Devices) == 0 {
		log.Fatal("No devices found")
		return 0
	}

	return respObj.Devices[0].DeviceID
}
