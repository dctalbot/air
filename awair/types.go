package awair

import "time"

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
	Timestamp time.Time
	Score     int
	Sensors   []Datapoint
	Indices   []Datapoint
}

func (d AirDatum) Sensor(comp string) float32 {
	for _, s := range d.Sensors {
		if s.Comp == comp {
			return s.Value
		}
	}
	return -1
}

func (d AirDatum) Index(comp string) float32 {
	for _, s := range d.Indices {
		if s.Comp == comp {
			return s.Value
		}
	}
	return -1
}

type AirDataResponse struct {
	Data []AirDatum
}
