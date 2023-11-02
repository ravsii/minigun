package commands

import "github.com/ravsii/minigun/internal/mode"

func (h *CommandHandler) CommandRemoveRune(...string) {
	if h.M.CommandLine.Input() == "" {
		h.changeMode(mode.View)
	}

	h.M.CommandLine.RemoveRune()
}

func (h *CommandHandler) CommandSubmit(...string) {
	cmd := h.M.CommandLine.Input()
	h.M.CommandLine.Reset()
	h.CmdExecute(cmd)
}
