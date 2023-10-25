package mode

import "github.com/gdamore/tcell/v2"

type Mode int

const (
	View Mode = iota
	Console
)

var modeString = map[Mode]string{
	View:    "View",
	Console: "Console",
}

var modeColor = map[Mode]tcell.Color{
	View:    tcell.NewHexColor(0xFF0000),
	Console: tcell.NewHexColor(0x00FF00),
}

var current = View

func Set(m Mode) {
	current = m
}

func Current() Mode {
	return current
}

func String() string {
	return modeString[current]
}

func Color() tcell.Color {
	return modeColor[current]
}
