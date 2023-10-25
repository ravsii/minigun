package main

import (
	"log"

	"github.com/gdamore/tcell/v2"
	"github.com/ravsii/minigun/command"
	"github.com/ravsii/minigun/tab"
)

func main() {

	// Initialize screen
	s, err := tcell.NewScreen()
	if err != nil {
		log.Fatalf("%+v", err)
	}
	defer gracefulShotdown(s)
	if err := s.Init(); err != nil {
		log.Fatalf("%+v", err)
	}
	s.Clear()

	w, h := s.Size()

	root := tab.NewRootGroup(s, w, h-1)
	root.NewTab()
	if err := root.ActiveTab().FromPath("./main.go"); err != nil {
		log.Fatalf("%+v", err)
	}

	cmdLine := command.New(s)

	for {
		t := root.ActiveTab()
		t.Draw()
		cmdLine.Draw()

		// Update screen
		s.Show()

		// Process event
		switch event := s.PollEvent().(type) {
		case *tcell.EventResize:
			s.Sync()
		case *tcell.EventKey:
			if event.Rune() == ':' {
				cmdLine.HandleInput()
				continue
			}

			if t.HandleKey(s, event) {
				return
			}
		}
	}
}

func gracefulShotdown(s tcell.Screen) {
	// You have to catch panics in a defer, clean up, and
	// re-raise them - otherwise your application can
	// die without leaving any diagnostic trace.
	r := recover()
	s.Fini()
	if r != nil {
		panic(r)
	}
}
