package latex

import (
	"testing"
)

func TestGenerateWestDigitalIn(t *testing.T) {

	pin := 1
	text := "foo"
	result, err := generateWestDigitalIn(pin, text)
	if err != nil {
		t.Errorf("generateWestDigitalIn(%d, %s) returned error: %v", 1, "test.pinText", err)
	}

	if result == "" {
		t.Error("oops")
	}
}

func TestGenerateEastDigitalIn(t *testing.T) {

	pin := 1
	text := "foo"
	result, err := generateEastDigitalIn(pin, text)
	if err != nil {
		t.Errorf("generateEastDigitalIn(%d, %s) returned error: %v", 1, "test.pinText", err)
	}

	if result == "" {
		t.Error("oops")
	}
}
