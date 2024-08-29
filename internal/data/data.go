package data

import "fmt"

type Data struct {
	Title    string  `json:"title"`
	Subtitle string  `json:"subtitle"`
	Tag      string  `json:"tag"`
	Pin1     PinData `json:"pin1"`
	Pin2     PinData `json:"pin2"`
	Pin3     PinData `json:"pin3"`
	Pin4     PinData `json:"pin4"`
	Pin5     PinData `json:"pin5"`
	Pin6     PinData `json:"pin6"`
	Pin7     PinData `json:"pin7"`
	Body     string  `json:"body"`
}

type PinData struct {
	Pin     int    `json:"pin"`
	PinText string `json:"pintext"`
	PinType string `json:"pintype"`
}

// GetPinData takes a Data struct and a pin number and returns the corresponding PinData.
func GetPinData(allData Data, pin int) (PinData, error) {
	switch pin {
	case 1:
		return allData.Pin1, nil
	case 2:
		return allData.Pin2, nil
	case 3:
		return allData.Pin3, nil
	case 4:
		return allData.Pin4, nil
	case 5:
		return allData.Pin5, nil
	case 6:
		return allData.Pin6, nil
	case 7:
		return allData.Pin7, nil
	default:
		return PinData{}, fmt.Errorf("invalid pin number: %d", pin)
	}
}
