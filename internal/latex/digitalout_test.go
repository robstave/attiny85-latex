package latex

import (
	"testing"
)

func TestGenerateWestDigitalOut(t *testing.T) {

	pin := 1
	text := "foo"
	result, err := generateWestDigitalOut(pin, text)
	if err != nil {
		t.Errorf("generateWestDigitalOut(%d, %s) returned error: %v", 1, "test.pinText", err)
	}

	if result == "" {
		t.Error("oops")
	}
}

func TestGenerateEastDigitalOut(t *testing.T) {

	pin := 1
	text := "foo"
	result, err := generateEastDigitalOut(pin, text)
	if err != nil {
		t.Errorf("generateEastDigitalOut(%d, %s) returned error: %v", 1, "test.pinText", err)
	}

	if result == "" {
		t.Error("oops")
	}
}
