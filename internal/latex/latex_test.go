package latex

import (
	"attiny85-latex/internal/jsonparse"
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
	tmpl := template.Must(template.New("test").Parse("{{.Title}} - {{.Color}} {{.Fruit}}"))

	data := jsonparse.Data{
		Title: "Yum",
		Color: "red",
		Fruit: "apple",
	}

	content, err := GenerateLaTeXContent(tmpl, data)
	if err != nil {
		t.Fatalf("Failed to generate LaTeX content: %v", err)
	}

	expected := "Yum - red apple"
	if strings.TrimSpace(content) != expected {
		t.Fatalf("Expected '%s', got '%s'", expected, content)
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
