package latex

func generateEastDigitalIn(pin int, pinText string) (string, error) {
	templatePath := "../../templates/east-digitalin.tex"
	return generateFromTemplate(templatePath, pin, pinText)
}

func generateWestDigitalIn(pin int, pinText string) (string, error) {
	templatePath := "../../templates/west-digitalin.tex"
	return generateFromTemplate(templatePath, pin, pinText)

}
