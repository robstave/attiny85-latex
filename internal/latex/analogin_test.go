package latex

import (
	"testing"
)

func TestGenerateWestAnalogIn(t *testing.T) {

	pin := 1
	text := "foo"
	result, err := generateWestAnalogIn(pin, text)
	if err != nil {
		t.Errorf("generateWestAnalogIn(%d, %s) returned error: %v", 1, "test.pinText", err)
	}

	if result == "" {
		t.Error("oops")
	}
}

func TestGenerateEastAnalogIn(t *testing.T) {

	pin := 1
	text := "foo"
	result, err := generateEastAnalogIn(pin, text)
	if err != nil {
		t.Errorf("generateEastAnalogIn(%d, %s) returned error: %v", 1, "test.pinText", err)
	}

	if result == "" {
		t.Error("oops")
	}
}
