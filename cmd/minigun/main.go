package main

import (
	"flag"

	"github.com/gdamore/tcell/v2"
	"github.com/ravsii/minigun/internal/command"
	"github.com/ravsii/minigun/internal/config/binds"
	"github.com/ravsii/minigun/internal/config/log"
	"github.com/ravsii/minigun/internal/kbhandler"
	"github.com/ravsii/minigun/internal/minigun"
	"github.com/ravsii/minigun/internal/screen"
)

var logFilePath string

func main() {
	parseFlags()
	if err := binds.Load(); err != nil {
		panic(err)
	}
	log.Init(logFilePath)
	defer screen.Finish()

	mg := minigun.New()
	ch := command.New(&mg)
	kh := kbhandler.New(&ch)

	args := flag.Args()
	if len(args) > 0 {
		ch.OpenFile(args...)
	}

	mg.Draw()

	for {
		switch event := screen.Screen().PollEvent().(type) {
		case *tcell.EventResize:
			screen.Screen().Sync()
			mg.Tab.Resize()
		case *tcell.EventKey:
			kh.Handle(event)
		}
	}
}

func parseFlags() {
	flag.StringVar(&logFilePath, "logfile", "", "debug.log filepath")
	flag.Parse()
}
