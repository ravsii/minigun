package binds

import (
	"fmt"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/ravsii/minigun/internal/config"
	"github.com/ravsii/minigun/internal/mode"
)

const keybindsFilename = "keybinds.toml"

var binds Keybinds

type Keybinds struct {
	View    map[string]string `toml:"view"`
	Command map[string]string `toml:"command"`
	Replace map[string]string `toml:"replace"`
}

func (c *Keybinds) Merge(new *Keybinds) {
	for key, cmd := range new.View {
		c.View[strings.ToLower(key)] = cmd
	}
}

func Load() error {
	var newBinds Keybinds

	global, err := loadGlobal()
	if err != nil {
		return fmt.Errorf("global: %s", err)
	}
	newBinds.Merge(global)

	local, err := loadLocal()
	if err != nil {
		return fmt.Errorf("local: %s", err)
	}
	newBinds.Merge(local)

	binds = newBinds
	return nil
}

func CommandFor(m mode.Mode, key string) (string, bool) {
	var cmd string
	switch m {
	case mode.View:
		if bind, ok := binds.View[key]; ok {
			cmd = bind
		}
	default:
		return "", false
	}

	if cmd == "" {
		return "", false
	}

	return cmd, true
}

func loadGlobal() (*Keybinds, error) {
	return nil, nil
}

func loadLocal() (*Keybinds, error) {
	f, exists, err := config.ProjectSpecificDirFile(keybindsFilename)
	if !exists {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	var local Keybinds

	if _, err := toml.NewDecoder(f).Decode(&local); err != nil {
		return nil, err
	}

	return &binds, nil
}
