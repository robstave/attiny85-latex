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
