package latex

import (
	"attiny85-latex/internal/data"
	"errors"
	"fmt"
	"os"
	"strings"
	"text/template"
)

// BuildSubTemplate constructs the template for a given set of pin numbers and their corresponding build function.
func BuildSubTemplate(allData data.Data, pinNumbers []int, buildPinFunc func(data.PinData) (string, error)) (string, error) {
	var result string
	for _, pin := range pinNumbers {
		pinData, err := data.GetPinData(allData, pin)
		if err != nil {
			return "", fmt.Errorf("GetPinData pin %d: %w", pin, err)
		}

		pinString, err := buildPinFunc(pinData)
		if err != nil {
			return "", fmt.Errorf("buildSubPin pin %d: %w", pin, err)
		}

		result += pinString
	}
	return result, nil
}

// BuildWestSubTemplate constructs the West sub-template by invoking BuildSubTemplate with West-specific pin numbers and function.
func BuildWestSubTemplate(allData data.Data) (string, error) {
	return BuildSubTemplate(allData, []int{2, 3}, buildWestSubPin)
}

// BuildEastSubTemplate constructs the East sub-template by invoking BuildSubTemplate with East-specific pin numbers and function.
func BuildEastSubTemplate(allData data.Data) (string, error) {
	return BuildSubTemplate(allData, []int{5, 6, 7}, buildEastSubPin)
}

// buildWestSubPin generates the string for a West sub-pin based on its type.
func buildWestSubPin(pd data.PinData) (string, error) {
	switch pd.PinType {
	case data.ANALOGIN:
		return generateWestAnalogIn(pd.Pin, pd.PinText)
	case data.DIGIN:
		return generateWestDigitalIn(pd.Pin, pd.PinText)
	case data.DIGOUT:
		return generateWestDigitalOut(pd.Pin, pd.PinText)
	default:
		return "", errors.New("unsupported pin type for BuildWestSubPin")
	}
}

// buildEastSubPin generates the string for an East sub-pin based on its type.
func buildEastSubPin(pd data.PinData) (string, error) {
	switch pd.PinType {
	case data.PWM:
		return generateEastPMWOut(pd.Pin, pd.PinText)
	case data.ANALOGIN:
		return generateEastAnalogIn(pd.Pin, pd.PinText)
	case data.DIGIN:
		return generateEastDigitalIn(pd.Pin, pd.PinText)
	case data.DIGOUT:
		return generateEastDigitalOut(pd.Pin, pd.PinText)
	default:
		return "", errors.New("unsupported pin type for BuildEastSubPin")
	}
}

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
func GenerateLaTeX(outputPath string, data data.Data) error {

	west, err := BuildWestSubTemplate(data)
	if err != nil {
		fmt.Println("error building west")

		return err
	}
	//fmt.Println(west)
	east, err := BuildEastSubTemplate(data)
	if err != nil {
		fmt.Println("error building east")

		return err
	}

	data.Body = generateBody() + west + east

	tmpl, err := template.New("example").Parse(generateMain())
	if err != nil {
		panic(err)
	}

	content, err := GenerateLaTeXContent(tmpl, data)
	if err != nil {
		return err
	}

	return WriteToFile(outputPath, content)
}
