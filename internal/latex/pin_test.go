package latex

import (
	"attiny85-latex/internal/data"

	"fmt"
	"strings"
	"testing"
)

// TestGenerateFromTemplateString_Valid tests the basic valid template usage.
func TestGenerateFromTemplateString_Valid(t *testing.T) {
	// Define a simple template that outputs the pin, its Ypos, and text.
	tmplStr := "Pin: {{.Pin}}, Ypos: {{.Ypos}}, Text: {{.PinText}}"

	// Use a sample pin and text.
	pin := 3
	text := "TestText"
	expectedYpos := data.PinValue(pin) // expected Ypos value from your data package

	result, err := generateFromTemplateString(tmplStr, pin, text)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	// Verify that the output contains the expected values.
	if !strings.Contains(result, fmt.Sprintf("Pin: %d", pin)) {
		t.Errorf("Result %q missing expected pin value", result)
	}
	// For the float, we format it with a similar precision as in the output.
	expectedYposStr := fmt.Sprintf("Ypos: %g", expectedYpos)
	if !strings.Contains(result, expectedYposStr) {
		t.Errorf("Result %q missing expected Ypos value %q", result, expectedYposStr)
	}

	if !strings.Contains(result, fmt.Sprintf("Text: %s", text)) {
		t.Errorf("Result %q missing expected text %q", result, text)
	}
}

// TestGenerateFromTemplateString_Multiline tests a multi-line template.
func TestGenerateFromTemplateString_Multiline(t *testing.T) {
	// Define a template spanning several lines.
	tmplStr := `Line 1: Pin {{.Pin}}
Line 2: Ypos {{.Ypos}}
Line 3: Text {{.PinText}}`

	pin := 2
	text := "HelloWorld"
	result, err := generateFromTemplateString(tmplStr, pin, text)
	if err != nil {
		t.Fatalf("Unexpected error in multi-line template: %v", err)
	}

	if !strings.Contains(result, "Pin 2") {
		t.Errorf("Expected 'Pin 2' in result, got: %s", result)
	}
	if !strings.Contains(result, "Ypos") {
		t.Errorf("Expected 'Ypos' in result, got: %s", result)
	}
	if !strings.Contains(result, "Text HelloWorld") {
		t.Errorf("Expected 'Text HelloWorld' in result, got: %s", result)
	}
}

// TestGenerateFromTemplateString_EmptyText checks that an empty PinText is handled.
func TestGenerateFromTemplateString_EmptyText(t *testing.T) {
	tmplStr := "Pin: {{.Pin}}, Ypos: {{.Ypos}}, Text: [{{.PinText}}]"
	pin := 5
	text := ""
	result, err := generateFromTemplateString(tmplStr, pin, text)
	if err != nil {
		t.Fatalf("Unexpected error for empty text: %v", err)
	}

	// Expect the text section to display empty brackets.
	if !strings.Contains(result, "Text: []") {
		t.Errorf("Expected output to contain 'Text: []', got: %s", result)
	}
}

// TestGenerateFromTemplateString_InvalidTemplate ensures that an invalid template panics.
// Since your current implementation calls panic on parse error, we must recover from it.
func TestGenerateFromTemplateString_InvalidTemplate(t *testing.T) {
	// Create an intentionally invalid template (for example, a missing closing brace).
	tmplStr := "Pin: {{.Pin, Ypos: {{.Ypos}}, Text: {{.PinText}}"

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic for invalid template syntax, but no panic occurred")
		}
		// Optionally, you can check that r contains an expected substring.
	}()

	// This call should panic due to the invalid template.
	_, _ = generateFromTemplateString(tmplStr, 1, "bad")
}
