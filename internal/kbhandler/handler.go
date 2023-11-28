package kbhandler

import (
	"strings"
	"unicode"

	"github.com/gdamore/tcell/v2"
	"github.com/ravsii/minigun/internal/commands"
	"github.com/ravsii/minigun/internal/config/binds"
	"github.com/ravsii/minigun/internal/mode"
	"github.com/ravsii/minigun/internal/screen"
)

type KeybindHandler struct {
	c *commands.CommandHandler

	aliases map[string]string
}

func New(c *commands.CommandHandler) *KeybindHandler {
	return &KeybindHandler{
		c: c,
		aliases: map[string]string{
			"backspace2": "backspace",
		},
	}
}

func (h *KeybindHandler) Handle(e tcell.Event) {
	switch mode.Current() {
	case mode.View:
		h.handleView(e)
	case mode.Command:
		h.handleCommand(e)
	case mode.Replace:
		h.handleReplace(e)
	case mode.Edit:
		h.handleEdit(e)
	default:
		h.c.Info("unknown mode ", mode.Current().String())
	}

	screen.Show()
}

func (h *KeybindHandler) handleView(e tcell.Event) {
	cmd, found := h.cmdFromEvent(e)
	if !found {
		return
	}

	h.c.CmdExecute(cmd)
}

func (h *KeybindHandler) handleCommand(e tcell.Event) {
	key, ok := e.(*tcell.EventKey)
	if !ok {
		return
	}

	cmd, found := h.cmdFromEvent(e)
	if found {
		h.c.CmdExecute(cmd)
		return
	}

	r := key.Rune()
	if !unicode.IsGraphic(r) {
		return
	}

	h.c.M.CommandLine.AddRune(r)
}

func (h *KeybindHandler) handleReplace(event tcell.Event) {
	key, ok := event.(*tcell.EventKey)
	if !ok {
		return
	}

	for {
		r := key.Rune()
		if !unicode.IsGraphic(r) {
			continue
		}

		h.c.ReplaceRune(string(r))
		break
	}

	h.c.EnterViewMode()
}

func (h *KeybindHandler) handleEdit(event tcell.Event) {
	key, ok := event.(*tcell.EventKey)
	if !ok {
		return
	}

	for {
		r := key.Rune()
		if !unicode.IsGraphic(r) {
			continue
		}

		h.c.ReplaceRune(string(r))
		break
	}

	h.c.EnterViewMode()
}

func (h *KeybindHandler) cmdFromEvent(e tcell.Event) (string, bool) {
	key, ok := e.(*tcell.EventKey)
	if !ok {
		return "", false
	}

	var k string

	if key.Key() == tcell.KeyRune {
		k = string(key.Rune())
	} else {
		k = key.Name()
	}

	k = strings.ToLower(k)
	if alias, ok := h.aliases[k]; ok {
		k = alias
	}

	cmd, found := binds.CommandFor(mode.Current(), k)
	if !found {
		return "", false
	}

	return cmd, true
}
