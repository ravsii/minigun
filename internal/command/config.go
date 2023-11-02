package command

import (
	"github.com/ravsii/minigun/internal/config/binds"
)

func (h *CommandHandler) ReloadConfig(...string) {
	if err := binds.Load(); err != nil {
		h.m.CommandLine.Errorf("can't load config: %s", err)
	}
}
