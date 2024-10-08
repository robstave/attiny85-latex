package latex

import (
	"strings"
	"testing"
)

func TestGenerateWestAnalogIn(t *testing.T) {

	pin := 1
	text := "foo"
	result, err := generateWestAnalogIn(pin, text)
	if err != nil {
		t.Errorf("generateEastAnalogIn(%d, %s) returned error: %v", pin, text, err)
	}

	substring := "{ foo }"
	if !strings.Contains(result, substring) {
		t.Error("oops")
	}

}

func TestGenerateEastAnalogIn(t *testing.T) {

	pin := 1
	text := "foo2"
	result, err := generateEastAnalogIn(pin, text)
	if err != nil {
		t.Errorf("generateEastAnalogIn(%d, %s) returned error: %v", pin, text, err)
	}

	substring := "{ foo2 }"
	if !strings.Contains(result, substring) {
		t.Error("oops")
	}
}

func TestGenerateWestDigitalIn(t *testing.T) {

	pin := 1
	text := "foo"
	result, err := generateWestDigitalIn(pin, text)
	if err != nil {
		t.Errorf("generateWestDigitalIn(%d, %s) returned error: %v", pin, text, err)
	}

	substring := "{ foo }"
	if !strings.Contains(result, substring) {
		t.Error("oops")
	}
}

func TestGenerateEastDigitalIn(t *testing.T) {

	pin := 1
	text := "foo"
	result, err := generateEastDigitalIn(pin, text)
	if err != nil {
		t.Errorf("generateEastDigitalIn(%d, %s) returned error: %v", pin, text, err)
	}

	substring := "{ foo }"
	if !strings.Contains(result, substring) {
		t.Error("oops")
	}
}

func TestGenerateWestDigitalOut(t *testing.T) {

	pin := 1
	text := "foo"
	result, err := generateWestDigitalOut(pin, text)
	if err != nil {
		t.Errorf("generateWestDigitalOut(%d, %s) returned error: %v", pin, text, err)
	}

	substring := "{ foo }"
	if !strings.Contains(result, substring) {
		t.Error("oops")
	}
}

func TestGenerateEastDigitalOut(t *testing.T) {

	pin := 1
	text := "foo"
	result, err := generateEastDigitalOut(pin, text)
	if err != nil {
		t.Errorf("generateEastDigitalOut(%d, %s) returned error: %v", pin, text, err)
	}

	substring := "{ foo }"
	if !strings.Contains(result, substring) {
		t.Error("oops")
	}
}

func TestGenerateEastPWMOut(t *testing.T) {

	pin := 1
	text := "foo"
	result, err := generateEastPMWOut(pin, text)
	if err != nil {
		t.Errorf("generateEastPMWOut(%d, %s) returned error: %v", 1, "test.pinText", err)
	}

	substring := "{ foo }"
	if !strings.Contains(result, substring) {
		t.Error("oops")
	}
}
