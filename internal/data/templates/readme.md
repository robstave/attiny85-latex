# LaTeX Templates for Attiny85

This directory contains the LaTeX template files used to generate the final circuit diagrams and documents for the Attiny85 microcontroller. The current implementation supports the Attiny85, with the possibility of adding support for additional microcontrollers in the future.

---

## Template Files

Below is a description of each file contained in this directory:

- **attiny85.tex**  
  The main template for the Attiny85. This file:
  - Sets up the document class and includes required packages (e.g., `circuitikz`, `tikz`, `geometry`).
  - Contains placeholders for document metadata (such as `Title`, `Subtitle`, and `Tag`), as well as the overall circuit layout (`Body`).
  - Serves as the container into which chip-specific sub-templates are injected.

- **body.tex**  
  Defines the circuit diagram of the Attiny85 chip:
  - Lays out the chip using `circuitikz`.
  - Specifies the positions and labels of the pins.
  - Acts as the base for inserting additional pin-specific diagram elements.

- **east-analogin.tex**  
  Provides the drawing instructions for an analog input pin on the east (right) side of the chip:
  - Uses placeholders such as `Ypos` to determine the vertical position.
  - Displays the label (`PinText`) associated with the pin.

- **east-digitalin.tex**  
  Contains the template for a digital input pin on the east side:
  - Configures the arrow and label positioning to clearly indicate a digital input.
  - Uses a similar layout to the analog input template but with modifications that reflect digital signal conventions.

- **east-digitalout.tex**  
  Used for rendering a digital output pin on the east side:
  - Features a graphical style that distinguishes output functionality.
  - Adjusts the path and labeling compared to the input templates.

- **east-pwmout.tex**  
  Provides the layout for a PWM (Pulse-Width Modulation) output pin on the east side:
  - Incorporates additional graphic elements that help visualize PWM-specific behavior.
  - Uses the same dynamic positioning placeholders as the other east-side templates.

- **west-analogin.tex**  
  Defines the template for an analog input pin on the west (left) side of the chip:
  - Mirrors the east analog input style to ensure visual consistency.
  - Adjusts the drawing path and label alignment to the left side of the chip.

- **west-digitalin.tex**  
  Contains the drawing instructions for a digital input pin on the west side:
  - Similar in structure to the east digital input template, with adaptations for left-side placement.

- **west-digitalout.tex**  
  Used for rendering a digital output pin on the west side:
  - Follows the style of the other output templates but with positioning tailored for the west side.

---

## How It Works

1. **Embedding Templates**  
   These template files are embedded into the Go binary using the [`go:embed`](https://pkg.go.dev/embed) directive. This allows the application to access the template contents at runtime without needing separate files.

2. **Assembling the LaTeX Document**  
   The LaTeX file is generated in several steps:
   - **Main Template:**  
     The `attiny85.tex` file acts as the container for the entire document.
   - **Chip Body:**  
     The `body.tex` file defines the overall chip layout (including the position of pins).
   - **Pin Sub-Templates:**  
     Depending on the pin type specified in the input JSON, the application dynamically calls functions to process the corresponding sub-template (e.g., `east-analogin.tex`, `west-digitalin.tex`, etc.). Placeholders like `{{.Ypos}}` and `{{.PinText}}` are populated based on the pin configuration.
   - **Document Assembly:**  
     The resulting pin renderings are concatenated with the chip body, and then the completed LaTeX content is injected into the main template to produce the final document.

3. **Output**  
   The final LaTeX document is written to an output file (e.g., `output/output.tex`), which can then be processed with a LaTeX engine to generate the corresponding PDF or other document formats.

---
 