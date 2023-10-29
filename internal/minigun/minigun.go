package minigun

import (
	"github.com/ravsii/minigun/internal/cmdline"
	"github.com/ravsii/minigun/internal/screen"
	"github.com/ravsii/minigun/internal/statusbar"
	"github.com/ravsii/minigun/internal/tab"
)

// Minigun is the main "app" struct.
type Minigun struct {
	CommandLine cmdline.CommandLine
	Tab         tab.Tab
	StatusBar   statusbar.StatusBar
}

func New() Minigun {
	w, h := screen.Screen().Size()

	t := tab.New(w, h-2, 0, 0)
	commandLine := cmdline.New()
	statusBar := statusbar.New()

	return Minigun{
		CommandLine: commandLine,
		Tab:         *t,
		StatusBar:   statusBar,
	}
}

func (m *Minigun) Draw() {
	m.Tab.Draw()
	m.StatusBar.Draw()
	m.CommandLine.Draw()
}
