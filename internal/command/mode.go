package command

import (
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/ravsii/minigun/internal/mode"
	"github.com/ravsii/minigun/internal/screen"
)

func (h *CommandHandler) EnterViewMode(...string) {
	if mode.Current() == mode.Command {
		h.ClearCommandLine()
	}

	h.changeMode(mode.View)
}

func (h *CommandHandler) EnterCommandMode(...string) {
	h.changeMode(mode.Command)
	screen.Screen().SetCursorStyle(tcell.CursorStyleBlinkingBar)
	defer func() {
		screen.Screen().SetCursorStyle(tcell.CursorStyleBlinkingBlock)
		screen.Screen().HideCursor()
		h.changeMode(mode.View)
	}()

	input := h.m.CommandLine.HandleUserInput()
	if input != "" {
		h.CmdExecute(strings.Split(input, " ")...)
	}
}

func (h *CommandHandler) EnterReplaceMode(...string) {
	h.changeMode(mode.Replace)
}

func (h *CommandHandler) changeMode(m mode.Mode) {
	mode.Set(m)
	h.m.StatusBar.Draw()
}
