package jsonparse

import (
	"attiny85-latex/internal/data"
	"testing"
)

func TestLoadJSON(t *testing.T) {
	// Test with a valid JSON file
	alldata, err := LoadJSON("../../examples/valid.json")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	title := "ACS123"
	subtitle := "a thing"
	tag := "stuff"
	// Replace the following with actual field checks in your Data struct
	if alldata.Title != title {
		t.Errorf("Expected title %s, got %s", title, alldata.Title)
	}
	if alldata.Subtitle != subtitle {
		t.Errorf("Expected subtitle %s, got %s", title, alldata.Subtitle)
	}
	if alldata.Tag != tag {
		t.Errorf("Expected tag %s, got %s", title, alldata.Tag)
	}

	checkPin(alldata.Pin2, 2, data.ANALOGIN, "pot1", t)
	checkPin(alldata.Pin3, 3, data.ANALOGIN, "pot2", t)
	checkPin(alldata.Pin5, 5, data.DIGOUT, "trigger", t)
	checkPin(alldata.Pin6, 6, data.PWM, "pwm", t)
	checkPin(alldata.Pin7, 7, data.DIGIN, "clock", t)

}

func checkPin(pd data.PinData, pinNumber int, pinType string, pinText string, t *testing.T) {
	if pd.Pin != pinNumber {
		t.Errorf("Expected pinNumber %d, got %d", pinNumber, pd.Pin)
	}
	if pd.PinType != pinType {
		t.Errorf("Expected pinNumber %s, got %s", pinType, pd.PinType)
	}
	if pd.PinText != pinText {
		t.Errorf("Expected PinText %s, got %s", pinText, pd.PinText)
	}

}
