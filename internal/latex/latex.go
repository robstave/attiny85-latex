package latex

import (
	"attiny85-latex/internal/jsonparse"
	"os"
	"strings"
	"text/template"
)

// ParseTemplate loads and parses the LaTeX template file.
func ParseTemplate(templatePath string) (*template.Template, error) {
	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		return nil, err
	}
	return tmpl, nil
}

// GenerateLaTeXContent applies the parsed template with the data and returns the LaTeX content as a string.
func GenerateLaTeXContent(tmpl *template.Template, data jsonparse.Data) (string, error) {
	var result strings.Builder
	err := tmpl.Execute(&result, data)
	if err != nil {
		return "", err
	}
	return result.String(), nil
}

// WriteToFile writes the LaTeX content to the specified file.
func WriteToFile(outputPath, content string) error {
	file, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(content)
	if err != nil {
		return err
	}

	return nil
}

// GenerateLaTeX combines the steps to generate the LaTeX file from template and data.
func GenerateLaTeX(templatePath, outputPath string, data jsonparse.Data) error {
	tmpl, err := ParseTemplate(templatePath)
	if err != nil {
		return err
	}

	content, err := GenerateLaTeXContent(tmpl, data)
	if err != nil {
		return err
	}

	return WriteToFile(outputPath, content)
}
