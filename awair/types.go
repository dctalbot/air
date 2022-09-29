package awair

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

type Datapoint struct {
	Comp  string
	Value float32
}

type AirDatum struct {
	Timestamp string
	Score     int
	Sensors   []Datapoint
	Indices   []Datapoint
}

type AirDataResponse struct {
	Data []AirDatum
}
