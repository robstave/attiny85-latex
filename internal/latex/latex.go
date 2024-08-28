package latex

import (
	"attiny85-latex/internal/data"
	"fmt"
	"os"
	"strings"
	"text/template"
)

func buildWestSubTemplate(pin int, allData data.Data) (string, error) {

	pd2, err := data.GetPinData(allData, 2)
	if err != nil {
		return "", err
	}

	p2String, err := buildWestSubPin(pd2)
	if err != nil {
		return "", err
	}

	fmt.Println("p2String", p2String)

	pd3, err := data.GetPinData(allData, 3)
	if err != nil {
		return "", err
	}

	p3String, err := buildWestSubPin(pd3)
	if err != nil {
		return "", err
	}

	fmt.Println("p3String", p3String)

	return "", nil

}

func buildWestSubPin(pd data.PinData) (string, error) {

	if pd.PinType == data.ANALOGIN {
		return generateWestAnalogIn(pd.Pin, pd.PinText)
	}

	if pd.PinType == data.DIGIN {
		return generateWestDigitalIn(pd.Pin, pd.PinText)

	}

	if pd.PinType == data.DIGOUT {
		return generateWestDigitalOut(pd.Pin, pd.PinText)

	}

	return "", nil
}

func buildEastSubPin(pd data.PinData) (string, error) {

	if pd.PinType == data.PWM {
		return generateEastPMWOut(pd.Pin, pd.PinText)
	}

	if pd.PinType == data.ANALOGIN {
		return generateEastAnalogIn(pd.Pin, pd.PinText)
	}

	if pd.PinType == data.DIGIN {
		return generateEastDigitalIn(pd.Pin, pd.PinText)

	}

	if pd.PinType == data.DIGOUT {
		return generateEastDigitalOut(pd.Pin, pd.PinText)

	}

	return "", nil
}

// GenerateLaTeXContent applies the parsed template with the data and returns the LaTeX content as a string.
func GenerateLaTeXContent(tmpl *template.Template, data data.Data) (string, error) {
	var result strings.Builder
	err := tmpl.Execute(&result, data)
	if err != nil {
		return "", err
	}
	return result.String(), nil
}

// ParseTemplate loads and parses the LaTeX template file.
func ParseTemplate(templatePath string) (*template.Template, error) {
	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		return nil, err
	}
	return tmpl, nil
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
func GenerateLaTeX(templatePath, outputPath string, data data.Data) error {
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
