package commands

import (
	"github.com/ravsii/minigun/internal/mode"
)

func (h *CommandHandler) EnterViewMode(...string) {
	if mode.Current() == mode.Command {
		h.ClearCommandLine()
	}

	h.changeMode(mode.View)
}

func (h *CommandHandler) EnterCommandMode(...string) {
	h.changeMode(mode.Command)
}

func (h *CommandHandler) EnterReplaceMode(...string) {
	h.changeMode(mode.Replace)
}

func (h *CommandHandler) changeMode(m mode.Mode) {
	mode.Set(m)
	h.M.StatusBar.Draw()
}
