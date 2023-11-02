package commands

import "github.com/ravsii/minigun/internal/mode"

func (h *CommandHandler) CommandRemoveRune(...string) {
	h.M.CommandLine.RemoveRune()
	if h.M.CommandLine.Input() == "" {
		h.changeMode(mode.View)
	}
}

func (h *CommandHandler) CommandSubmit(...string) {
	cmd := h.M.CommandLine.Input()
	if cmd == "" {
		h.changeMode(mode.View)
		return
	}

	h.M.CommandLine.Reset()
	h.CmdExecute(cmd)
	h.changeMode(mode.View)
}
