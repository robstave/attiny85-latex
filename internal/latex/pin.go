package latex

import (
	"attiny85-latex/internal/data"
	"errors"
	"html/template"
	"strings"
)

type Pin struct {
	Ypos    float64 `json:"ypos"`
	Pin     int     `json:"pin"`
	PinText string  `json:"pintext"`
}

func generateFromTemplateString(tmplStr string, pin int, pinText string) (string, error) {

	if pin < 0 {
		return "", errors.New("ban pin")
	}
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
