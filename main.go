package main

import (
	"log"

	"github.com/gdamore/tcell/v2"
	"github.com/ravsii/minigun/tab"
)

func main() {
	tab, err := tab.FromPath("./main.go")
	if err != nil {
		log.Fatalf("%+v", err)
	}

	// Initialize screen
	s, err := tcell.NewScreen()
	if err != nil {
		log.Fatalf("%+v", err)
	}
	if err := s.Init(); err != nil {
		log.Fatalf("%+v", err)
	}
	s.Clear()

	defer func() {
		// You have to catch panics in a defer, clean up, and
		// re-raise them - otherwise your application can
		// die without leaving any diagnostic trace.
		maybePanic := recover()
		s.Fini()
		if maybePanic != nil {
			panic(maybePanic)
		}
	}()

	for {
		tab.Draw(s)

		// Update screen
		s.Show()

		// Process event
		switch event := s.PollEvent().(type) {
		case *tcell.EventResize:
			s.Sync()
		case *tcell.EventKey:
			if tab.HandleKey(s, event) {
				return
			}
		}
	}
}
