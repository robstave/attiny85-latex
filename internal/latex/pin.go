package latex

import (
	"attiny85-latex/internal/data"
	"fmt"
	"html/template"
	"strings"
)

type Pin struct {
	Ypos    float64 `json:"ypos"`
	Pin     int     `json:"pin"`
	PinText string  `json:"pintext"`
}

func generateFromTemplate(templatePath string, pin int, pinText string) (string, error) {
	tmpl, err := ParseTemplate(templatePath)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	pinData := Pin{
		Ypos:    data.PinValue(pin),
		Pin:     pin,
		PinText: pinText,
	}

	var result strings.Builder
	err = tmpl.Execute(&result, pinData)
	if err != nil {
		return "", err
	}
	return result.String(), nil
}

func generateFromTemplateString(tmplStr string, pin int, pinText string) (string, error) {

	tmpl, err := template.New("example").Parse(tmplStr)
	if err != nil {
		panic(err)
	}

	pinData := Pin{
		Ypos:    data.PinValue(pin),
		Pin:     pin,
		PinText: pinText,
	}

	var result strings.Builder
	err = tmpl.Execute(&result, pinData)
	if err != nil {
		return "", err
	}
	return result.String(), nil
}
