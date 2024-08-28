package latex

import (
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
		return "", err
	}

	pinData := Pin{
		Ypos:    PinValue(pin),
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
