package latex

import (
	"attiny85-latex/internal/data"
)

func generateBody() string {

	return data.ATTinyBody
}

func generateWestAnalogIn(pin int, pinText string) (string, error) {

	templatePath := data.WestAnalogInTemplate
	//templatePath := "../../templates/west-analogin.tex"
	return generateFromTemplateString(templatePath, pin, pinText)
}

func generateEastAnalogIn(pin int, pinText string) (string, error) {
	//templatePath := "../../templates/east-analogin.tex"
	templatePath := data.EastAnalogInTemplate
	return generateFromTemplateString(templatePath, pin, pinText)
}

func generateEastDigitalIn(pin int, pinText string) (string, error) {
	//templatePath := "../../templates/east-digitalin.tex"
	templatePath := data.EastDigitalInTemplate
	return generateFromTemplateString(templatePath, pin, pinText)
}

func generateWestDigitalIn(pin int, pinText string) (string, error) {
	//templatePath := "../../templates/west-digitalin.tex"
	templatePath := data.WestDigitalInTemplate
	return generateFromTemplateString(templatePath, pin, pinText)

}

func generateEastDigitalOut(pin int, pinText string) (string, error) {
	//templatePath := "../../templates/east-digitalout.tex"
	templatePath := data.EastDigitalOutTemplate
	return generateFromTemplateString(templatePath, pin, pinText)
}

func generateWestDigitalOut(pin int, pinText string) (string, error) {
	//templatePath := "../../templates/west-digitalout.tex"
	templatePath := data.WestDigitalOutTemplate
	return generateFromTemplateString(templatePath, pin, pinText)
}

func generateEastPMWOut(pin int, pinText string) (string, error) {
	//templatePath := "../../templates/east-pwmout.tex"
	templatePath := data.EastPwmOutTemplate

	return generateFromTemplateString(templatePath, pin, pinText)
}
