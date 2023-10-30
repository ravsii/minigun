package command

import "github.com/ravsii/minigun/internal/config"

func (h *CommandHandler) ReloadConfig(...string) {
	if err := config.Load(); err != nil {
		h.m.CommandLine.Errorf("can't reload config: %s", err)
	}
}
