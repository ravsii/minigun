package main

import (
	"os"

	"github.com/gdamore/tcell/v2"
	"github.com/ravsii/minigun/internal/command"
	"github.com/ravsii/minigun/internal/keybind"
	"github.com/ravsii/minigun/internal/minigun"
	"github.com/ravsii/minigun/internal/screen"
)

func main() {
	defer screen.Finish()

	mg := minigun.New()
	ch := command.New(&mg)
	kh := keybind.New(&ch)

	if len(os.Args) > 1 {
		ch.OpenFile(os.Args[1:]...)
	}

	mg.Draw()

	for {
		switch event := screen.Screen().PollEvent().(type) {
		case *tcell.EventResize:
			screen.Screen().Sync()
		case *tcell.EventKey:
			kh.Handle(event)
		}
	}

}
