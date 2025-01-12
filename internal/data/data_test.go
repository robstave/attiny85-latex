package data

import (
	"fmt"
	"testing"
)

func createTestData() Data {
	// Create a Data instance with distinct PinData for each pin.
	return Data{
		Title:    "Test Title",
		Subtitle: "Test Subtitle",
		Tag:      "Test Tag",
		Pin1:     PinData{Pin: 1, PinText: "Pin One", PinType: "Analog"},
		Pin2:     PinData{Pin: 2, PinText: "Pin Two", PinType: "Digital"},
		Pin3:     PinData{Pin: 3, PinText: "Pin Three", PinType: "Analog"},
		Pin4:     PinData{Pin: 4, PinText: "Pin Four", PinType: "Digital"},
		Pin5:     PinData{Pin: 5, PinText: "Pin Five", PinType: "PWM"},
		Pin6:     PinData{Pin: 6, PinText: "Pin Six", PinType: "Digital"},
		Pin7:     PinData{Pin: 7, PinText: "Pin Seven", PinType: "Analog"},
		Body:     "Test Body",
	}
}

func TestGetPinData_ValidPins(t *testing.T) {
	testData := createTestData()

	testCases := []struct {
		pin          int
		expectedPin  int
		expectedText string
		expectedType string
	}{
		{pin: 1, expectedPin: 1, expectedText: "Pin One", expectedType: "Analog"},
		{pin: 2, expectedPin: 2, expectedText: "Pin Two", expectedType: "Digital"},
		{pin: 3, expectedPin: 3, expectedText: "Pin Three", expectedType: "Analog"},
		{pin: 4, expectedPin: 4, expectedText: "Pin Four", expectedType: "Digital"},
		{pin: 5, expectedPin: 5, expectedText: "Pin Five", expectedType: "PWM"},
		{pin: 6, expectedPin: 6, expectedText: "Pin Six", expectedType: "Digital"},
		{pin: 7, expectedPin: 7, expectedText: "Pin Seven", expectedType: "Analog"},
	}

	for _, tc := range testCases {
		tc := tc // capture range variable
		t.Run(fmt.Sprintf("Pin%d", tc.pin), func(t *testing.T) {
			result, err := GetPinData(testData, tc.pin)
			if err != nil {
				t.Fatalf("Unexpected error for pin %d: %v", tc.pin, err)
			}
			if result.Pin != tc.expectedPin {
				t.Errorf("For pin %d, expected Pin %d, got %d", tc.pin, tc.expectedPin, result.Pin)
			}
			if result.PinText != tc.expectedText {
				t.Errorf("For pin %d, expected PinText %q, got %q", tc.pin, tc.expectedText, result.PinText)
			}
			if result.PinType != tc.expectedType {
				t.Errorf("For pin %d, expected PinType %q, got %q", tc.pin, tc.expectedType, result.PinType)
			}
		})
	}
}

func TestGetPinData_InvalidPin(t *testing.T) {
	testData := createTestData()

	// Choose a pin number that is not valid
	invalidPins := []int{-1, 0, 8, 99}
	for _, pin := range invalidPins {
		pin := pin // capture range variable
		t.Run(fmt.Sprintf("InvalidPin%d", pin), func(t *testing.T) {
			_, err := GetPinData(testData, pin)
			if err == nil {
				t.Errorf("Expected error for invalid pin %d but got none", pin)
			}
		})
	}
}
