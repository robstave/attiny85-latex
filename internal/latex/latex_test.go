package latex

import (
	"attiny85-latex/internal/data"
	"os"
	"strings"
	"testing"
	"text/template"
)

func TestParseTemplate(t *testing.T) {
	tmpl, err := ParseTemplate("../../templates/template.tex")
	if err != nil {
		t.Fatalf("Failed to parse template: %v", err)
	}

	if tmpl == nil {
		t.Fatal("Expected non-nil template")
	}
}

func TestGenerateLaTeXContent(t *testing.T) {
	tmpl := template.Must(template.New("test").Parse("{{.Title}}  "))

	data := data.Data{
		Title: "Yum",
	}

	_, err := GenerateLaTeXContent(tmpl, data)
	if err != nil {
		t.Fatalf("Failed to generate LaTeX content: %v", err)
	}

}

func TestWriteToFile(t *testing.T) {
	content := "Sample LaTeX content"
	outputPath := "../../output/test_output.tex"

	err := WriteToFile(outputPath, content)
	if err != nil {
		t.Fatalf("Failed to write to file: %v", err)
	}

	fileContent, err := os.ReadFile(outputPath)
	if err != nil {
		t.Fatalf("Failed to read the output file: %v", err)
	}

	if strings.TrimSpace(string(fileContent)) != content {
		t.Fatalf("Expected '%s', got '%s'", content, string(fileContent))
	}

	// Clean up the file after testing
	os.Remove(outputPath)
}

func TestBuildWestSubPin(t *testing.T) {
	pinData := data.PinData{Pin: 2, PinText: "in2", PinType: data.ANALOGIN}

	result, err := buildWestSubPin(pinData)

	if err != nil {
		t.Fatalf("Failed to write to file: %v", err)
	}

	if result == "" {
		t.Fatalf("%s", result)
	}

}

func TestBuildWestSubPin3(t *testing.T) {
	pinData := data.PinData{Pin: 3, PinText: "in3", PinType: data.DIGIN}
	result, err := buildWestSubPin(pinData)

	if err != nil {
		t.Fatalf("Failed to write to file: %v", err)
	}

	if result == "" {
		t.Fatalf("%s", result)
	}

}
