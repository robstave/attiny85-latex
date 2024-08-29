package latex

import (
	"attiny85-latex/internal/data"
)

func generateBody() string {
	return data.ATTinyBody
}

func generateMain() string {
	return data.ATTiny85
}

func generateWestAnalogIn(pin int, pinText string) (string, error) {
	templateStr := data.WestAnalogInTemplate
	return generateFromTemplateString(templateStr, pin, pinText)
}

func generateEastAnalogIn(pin int, pinText string) (string, error) {
	templateStr := data.EastAnalogInTemplate
	return generateFromTemplateString(templateStr, pin, pinText)
}

func generateEastDigitalIn(pin int, pinText string) (string, error) {
	templateStr := data.EastDigitalInTemplate
	return generateFromTemplateString(templateStr, pin, pinText)
}

func generateWestDigitalIn(pin int, pinText string) (string, error) {
	templateStr := data.WestDigitalInTemplate
	return generateFromTemplateString(templateStr, pin, pinText)

}

func generateEastDigitalOut(pin int, pinText string) (string, error) {
	templateStr := data.EastDigitalOutTemplate
	return generateFromTemplateString(templateStr, pin, pinText)
}

func generateWestDigitalOut(pin int, pinText string) (string, error) {
	templateStr := data.WestDigitalOutTemplate
	return generateFromTemplateString(templateStr, pin, pinText)
}

func generateEastPMWOut(pin int, pinText string) (string, error) {
	templateStr := data.EastPwmOutTemplate

	return generateFromTemplateString(templateStr, pin, pinText)
}
