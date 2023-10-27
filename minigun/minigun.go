package minigun

import (
	"log"

	"github.com/gdamore/tcell/v2"
	"github.com/ravsii/minigun/minigun/cmd"
	"github.com/ravsii/minigun/minigun/statusbar"
	"github.com/ravsii/minigun/minigun/tab"
)

// Minigun is the main "app" struct.
type Minigun struct {
	s           tcell.Screen
	commandLine *cmd.Command
	tab         *tab.Tab
	statusBar   *statusbar.StatusBar
}

func New(s tcell.Screen) Minigun {
	w, h := s.Size()

	commandLine := cmd.Init(s)
	statusBar := statusbar.Init(s)

	t := tab.New(s, w, h-2, 0, 0)
	if err := t.FromPath("./main.go"); err != nil {
		log.Fatalf("%+v", err)
	}

	return Minigun{
		s:           s,
		commandLine: commandLine,
		tab:         t,
		statusBar:   statusBar,
	}
}

func (m *Minigun) Run() {
	m.tab.Draw()
	m.statusBar.Draw()
	m.commandLine.Draw()

	for {
		// Update screen
		m.s.Show()

		switch event := m.s.PollEvent().(type) {
		case *tcell.EventResize:
			m.s.Sync()
		case *tcell.EventKey:
			if event.Key() == tcell.KeyCtrlC {
				return
			}

			if event.Rune() == ':' {
				m.commandLine.HandleInput()
				continue
			}

			if m.tab.HandleKey(m.s, event) {
				return
			}

		}
	}
}
