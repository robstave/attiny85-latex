package latex

func generateEastPMWOut(pin int, pinText string) (string, error) {
	templatePath := "../../templates/east-pwmout.tex"
	return generateFromTemplate(templatePath, pin, pinText)
}
