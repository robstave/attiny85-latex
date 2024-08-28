package latex

func generateEastDigitalOut(pin int, pinText string) (string, error) {
	templatePath := "../../templates/east-digitalout.tex"
	return generateFromTemplate(templatePath, pin, pinText)
}

func generateWestDigitalOut(pin int, pinText string) (string, error) {
	templatePath := "../../templates/west-digitalout.tex"
	return generateFromTemplate(templatePath, pin, pinText)
}
