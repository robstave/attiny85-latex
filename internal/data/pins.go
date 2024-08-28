package data

const ANALOGIN = "AIN"
const DIGIN = "DIN"
const DIGOUT = "DOUT"
const PWM = "PWM"

// PinValue returns the corresponding value based on the pin number
func PinValue(pin int) float64 {
	switch pin {
	case 8, 1:
		return 1.5
	case 7, 2:
		return 0.5
	case 6, 3:
		return -0.5
	case 5, 4:
		return -1.5
	default:
		return 0
	}
}
