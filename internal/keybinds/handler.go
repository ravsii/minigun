package keybinds

import (
	"unicode"

	"github.com/gdamore/tcell/v2"
	"github.com/ravsii/minigun/internal/command"
	"github.com/ravsii/minigun/internal/mode"
	"github.com/ravsii/minigun/internal/screen"
)

type KeybindHandler struct {
	c *command.CommandHandler
}

func New(c *command.CommandHandler) *KeybindHandler {
	return &KeybindHandler{c}
}

func (v *KeybindHandler) Handle(e tcell.Event) {
	switch mode.Current() {
	case mode.View:
		v.handleView(e)
	case mode.Console:
		v.handleCommand(e)
	case mode.Replace:
		v.handleReplace(e)
	default:
		v.c.Info("unknown mode ", mode.Current().String())
	}

	screen.Show()
}

func (v *KeybindHandler) handleView(event tcell.Event) {
	key, ok := event.(*tcell.EventKey)
	if !ok {
		return
	}

	switch {
	case key.Key() == tcell.KeyCtrlC:
		v.c.Quit()
	case key.Rune() == ':':
		v.c.EnterCommandMode()
	case key.Key() == tcell.KeyBackspace2:
		v.c.ClearCommandLine()
	case key.Rune() == 'H' || key.Rune() == 'h':
		v.c.MoveLeft()
	case key.Rune() == 'J' || key.Rune() == 'j':
		v.c.MoveDown()
	case key.Rune() == 'K' || key.Rune() == 'k':
		v.c.MoveUp()
	case key.Rune() == 'L' || key.Rune() == 'l':
		v.c.MoveRight()
	case key.Rune() == 'R' || key.Rune() == 'r':
		v.c.EnterReplaceMode()
	}
}

func (v *KeybindHandler) handleCommand(event tcell.Event) {
	key, ok := event.(*tcell.EventKey)
	if !ok {
		return
	}

	switch {
	case key.Key() == tcell.KeyCtrlC:
		v.c.Quit()
	case key.Rune() == ':':
		v.c.EnterCommandMode()
	case key.Key() == tcell.KeyBackspace2:
		v.c.ClearCommandLine()
	case key.Rune() == 'H' || key.Rune() == 'h':
		v.c.MoveLeft()
	case key.Rune() == 'J' || key.Rune() == 'j':
		v.c.MoveDown()
	case key.Rune() == 'K' || key.Rune() == 'k':
		v.c.MoveUp()
	case key.Rune() == 'L' || key.Rune() == 'l':
		v.c.MoveRight()

	}
}

func (v *KeybindHandler) handleReplace(event tcell.Event) {
	key, ok := event.(*tcell.EventKey)
	if !ok {
		return
	}

	for {
		r := key.Rune()
		if !unicode.IsGraphic(r) {
			continue
		}

		v.c.ReplaceSelected(string(r))

		break
	}

	v.c.EnterViewMode()
}
