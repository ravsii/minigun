package commands

import (
	"strings"

	"github.com/ravsii/minigun/internal/mode"
)

func (h *CommandHandler) CommandRemoveRune(...string) {
	h.M.CommandLine.RemoveRune()
	if h.M.CommandLine.Input() == "" {
		h.changeMode(mode.View)
	}
}

func (h *CommandHandler) CommandSubmit(...string) {
	input := h.M.CommandLine.Input()
	if input == "" {
		h.changeMode(mode.View)
		return
	}

	cmdWithArgs := strings.Fields(input)

	h.M.CommandLine.Reset()
	h.CmdExecute(cmdWithArgs...)
	h.changeMode(mode.View)
}
