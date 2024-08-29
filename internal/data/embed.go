package data

import _ "embed"

//go:embed templates/body.tex
var ATTinyBody string

//go:embed templates/west-analogin.tex
var WestAnalogInTemplate string

//go:embed templates/west-digitalin.tex
var WestDigitalInTemplate string

//go:embed templates/west-digitalout.tex
var WestDigitalOutTemplate string

//go:embed templates/east-analogin.tex
var EastAnalogInTemplate string

//go:embed templates/east-digitalin.tex
var EastDigitalInTemplate string

//go:embed templates/east-digitalout.tex
var EastDigitalOutTemplate string

//go:embed templates/east-pwmout.tex
var EastPwmOutTemplate string
