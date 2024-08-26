package main

import (
	"flag"
	"fmt"

	"attiny85-latex/internal/jsonparse"
	"attiny85-latex/internal/latex"
	"os"
)

func main() {
	// Parse command line arguments

	fmt.Println("Attiny85 to Latex")
	jsonFile := flag.String("json", "", "Path to the JSON file")
	flag.Parse()

	if *jsonFile == "" {
		fmt.Println("Usage: attiny85 -json <path_to_json_file>")
		os.Exit(1)
	}

	// Load the JSON data
	data, err := jsonparse.LoadJSON(*jsonFile)
	if err != nil {
		fmt.Println("Error loading JSON:", err)
		os.Exit(1)
	}

	fmt.Println("data", data)
	// Load the LaTeX template
	templatePath := "templates/template.tex"
	err = latex.GenerateLaTeX(templatePath, "output/output.tex", data)
	if err != nil {
		fmt.Println("Error generating LaTeX:", err)
		os.Exit(1)
	}

	fmt.Println("LaTeX file generated: output/output.tex")
}
