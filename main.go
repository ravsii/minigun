package main

import (
	"log"

	"github.com/gdamore/tcell/v2"
	"github.com/ravsii/minigun/minigun"
)

func main() {
	s, err := tcell.NewScreen()
	if err != nil {
		log.Fatalf("%+v", err)
	}

	defer gracefulShotdown(s)

	if err := s.Init(); err != nil {
		log.Fatalf("%+v", err)
	}

	s.Clear()

	mg := minigun.New(s)
	mg.Run()
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
