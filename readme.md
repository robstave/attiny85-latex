
# Attiny85 Latex generator

A small utility that creates a fairly opinionated, but decent ATTINY85 pinout latex file.

It takes a json file that specifies how to draw the image and translates it to a .tex file.

You still have to turn that to pdf or whatever.  I use Overleaf to inspect.

I suppose in the long run, it would be nice to just have the json files in github and have a service that batch generates the images as needed.

## Problem

I love this image ![example](assets/attiny85-example.png)

I have been using inkscape to draw a similar deal for years.  It tedious and error prone.

Latex is AWESOME.  So I created an app that takes in a json file like this

```json
{
    "title": "Cowbell",
    "subtitle": "General pattern",
    "tag": "github.com/robstave/ArduinoComponentSketches",
    "pin2": {
        "pin":2,
        "pintype": "AIN",
        "pintext": "Pot/CV 1"
    },
    "pin3": {
        "pin":3,
        "pintype": "AIN",
        "pintext": "Pot/CV 2"
    },
    "pin5": {
        "pin":5,
        "pintype": "DOUT",
        "pintext": "Out 0"
    },
    "pin6": {
        "pin":6,
        "pintype": "PWM",
        "pintext": "PWM out"
    },
    "pin7": {
        "pin":7,
        "pintype": "DIN",
        "pintext": "Clock/Trigger"
    }
}


```

Resulting in a tex file  ( link tbd)
that looks like this

TBD



## Json formatting

The json has

 - title
 - subtitle
 - tag

 and pins

 Pin 4 and 8 are reserved

     "pin7": {
        "pin":7,
        "pintype": "DIN",
        "pintext": "Clock/Trigger"
    }

    Pintype is
    - DIN  Digital in
    - DOUT Digital out
    - AIN Analog in
    - PWM - Pulse Width modulation



## Running  the project

` go run cmd/attiny85-latex/main.go -json examples/cowbell.json`