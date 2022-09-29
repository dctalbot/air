package awair

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

const BaseURL = "https://developer-apis.awair.is"
const DevicesURL = BaseURL + "/v1/users/self/devices"

func GetDevices() []Device {
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

	return respObj.Devices
}

func GetLatest(d Device) (result AirDatum) {
	httpClient := &http.Client{Timeout: 10 * time.Second}
	url := fmt.Sprintf("%s/v1/users/self/devices/%s/%d/air-data/latest?fahrenheit=true", BaseURL, d.DeviceType, d.DeviceID)
	req, err := http.NewRequest("GET", url, nil)
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

	var respObj AirDataResponse
	err = json.Unmarshal(data, &respObj)
	if err != nil {
		log.Println(err)
	}

	if len(respObj.Data) != 1 {
		log.Println("Expected 1 result, got", len(respObj.Data))
		return result
	}

	return respObj.Data[0]
}
