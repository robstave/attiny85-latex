package latex

import (
	"strings"
	"testing"
)

// Existing tests, unchanged:

func TestGenerateWestAnalogIn(t *testing.T) {
	pin := 1
	text := "foo"
	result, err := generateWestAnalogIn(pin, text)
	if err != nil {
		t.Errorf("generateWestAnalogIn(%d, %s) returned error: %v", pin, text, err)
	}

	substring := "{ foo }"
	if !strings.Contains(result, substring) {
		t.Errorf("Expected substring %q not found in result: %s", substring, result)
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
		t.Errorf("Expected substring %q not found in result: %s", substring, result)
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
		t.Errorf("Expected substring %q not found in result: %s", substring, result)
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
		t.Errorf("Expected substring %q not found in result: %s", substring, result)
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
		t.Errorf("Expected substring %q not found in result: %s", substring, result)
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
		t.Errorf("Expected substring %q not found in result: %s", substring, result)
	}
}

func TestGenerateEastPWMOut(t *testing.T) {
	pin := 1
	text := "foo"
	result, err := generateEastPMWOut(pin, text)
	if err != nil {
		t.Errorf("generateEastPMWOut(%d, %s) returned error: %v", pin, text, err)
	}

	substring := "{ foo }"
	if !strings.Contains(result, substring) {
		t.Errorf("Expected substring %q not found in result: %s", substring, result)
	}
}

// Additional tests:

// Test empty text input to verify that the functions handle it gracefully.
func TestGenerateWithEmptyText(t *testing.T) {
	pin := 2
	emptyText := ""

	// Testing one of the functions, repeat similarly for others if needed.
	result, err := generateWestAnalogIn(pin, emptyText)
	if err != nil {
		t.Errorf("generateWestAnalogIn(%d, %q) returned error: %v", pin, emptyText, err)
	}

	// You might decide what the expected behavior is. For example, maybe the
	// function should include an empty set of braces, like "{}"
	expectedSubstring := "{}"
	if !strings.Contains(result, expectedSubstring) {
		t.Errorf("Expected substring %q not found in result: %s", expectedSubstring, result)
	}
}

// Test for invalid pin values if your functions are expected to validate the pin range.
func TestGenerateInvalidPin(t *testing.T) {
	invalidPin := -1
	text := "invalid"

	// Check one function; add similar tests for the others if they should all validate the pin.
	_, err := generateEastDigitalOut(invalidPin, text)
	if err == nil {
		t.Errorf("Expected error for invalid pin %d, but got none", invalidPin)
	}
}

// Test for complete command structure – checking that the LaTeX command begins correctly.
func TestCommandStructure(t *testing.T) {
	pin := 3
	text := "cmdTest"
	result, err := generateWestDigitalIn(pin, text)
	if err != nil {
		t.Errorf("generateWestDigitalIn(%d, %s) returned error: %v", pin, text, err)
	}

	// Check that the string begins with a backslash (indicating a LaTeX command)
	if !strings.HasPrefix(strings.TrimSpace(result), "\\") {
		t.Errorf("Resulting LaTeX command does not appear to begin with a backslash: %s", result)
	}
}
