package keybind

import (
	"strconv"

	"github.com/gdamore/tcell/v2"
	"github.com/ravsii/minigun/internal/command"
	"github.com/ravsii/minigun/internal/screen"
)

type KeybindHandler struct {
	c *command.CommandHandler
}

func New(c *command.CommandHandler) KeybindHandler {
	return KeybindHandler{c: c}
}

func (h *KeybindHandler) Handle(e tcell.Event) {
	key, ok := e.(*tcell.EventKey)
	if !ok {
		return
	}

	switch {
	case key.Key() == tcell.KeyCtrlC:
		h.c.Quit()
	case key.Rune() == ':':
		h.c.GoIntoCommandMode()
	case key.Key() == tcell.KeyBackspace2:
		h.c.ClearCommandLine()
	case key.Rune() == 'H' || key.Rune() == 'h':
		h.c.MoveLeft()
	case key.Rune() == 'J' || key.Rune() == 'j':
		h.c.MoveDown()
	case key.Rune() == 'K' || key.Rune() == 'k':
		h.c.MoveUp()
	case key.Rune() == 'L' || key.Rune() == 'l':
		h.c.MoveRight()
	}

	h.c.Info("events", strconv.Itoa(screen.Updates))
	screen.Show()
}
