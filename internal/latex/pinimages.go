package latex

// GenerateLaTeXContent applies the parsed template with the data and returns the LaTeX content as a string.
func generateWestAnalogIn(pin int, pinText string) (string, error) {
	templatePath := "../../templates/west-analogin.tex"
	return generateFromTemplate(templatePath, pin, pinText)
}

func generateEastAnalogIn(pin int, pinText string) (string, error) {
	templatePath := "../../templates/east-analogin.tex"
	return generateFromTemplate(templatePath, pin, pinText)
}

func generateEastDigitalIn(pin int, pinText string) (string, error) {
	templatePath := "../../templates/east-digitalin.tex"
	return generateFromTemplate(templatePath, pin, pinText)
}

func generateWestDigitalIn(pin int, pinText string) (string, error) {
	templatePath := "../../templates/west-digitalin.tex"
	return generateFromTemplate(templatePath, pin, pinText)

}

func generateEastDigitalOut(pin int, pinText string) (string, error) {
	templatePath := "../../templates/east-digitalout.tex"
	return generateFromTemplate(templatePath, pin, pinText)
}

func generateWestDigitalOut(pin int, pinText string) (string, error) {
	templatePath := "../../templates/west-digitalout.tex"
	return generateFromTemplate(templatePath, pin, pinText)
}

func generateEastPMWOut(pin int, pinText string) (string, error) {
	templatePath := "../../templates/east-pwmout.tex"
	return generateFromTemplate(templatePath, pin, pinText)
}
