package binds

import (
	"fmt"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/ravsii/minigun/internal/config"
	"github.com/ravsii/minigun/internal/mode"
)

const keybindsFilename = "keybinds.toml"

var b Keybinds

type Keybinds struct {
	View    map[string]string `toml:"view"`
	Command map[string]string `toml:"command"`
	Replace map[string]string `toml:"replace"`
}

func newB() Keybinds {
	return Keybinds{
		View:    make(map[string]string),
		Command: make(map[string]string),
		Replace: make(map[string]string),
	}
}

func (c *Keybinds) Merge(new *Keybinds) {
	if new == nil {
		return
	}

	for key, cmd := range new.View {
		c.View[strings.ToLower(key)] = cmd
	}
	for key, cmd := range new.Command {
		c.Command[strings.ToLower(key)] = cmd
	}
	for key, cmd := range new.Replace {
		c.Replace[strings.ToLower(key)] = cmd
	}
}

func Load() error {
	newBinds := newB()

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

	b = newBinds
	return nil
}

func CommandFor(m mode.Mode, key string) (string, bool) {
	var cmd string
	switch m {
	case mode.View:
		if bind, ok := b.View[key]; ok {
			cmd = bind
		}
	case mode.Command:
		if bind, ok := b.Command[key]; ok {
			cmd = bind
		}
	case mode.Replace:
		if bind, ok := b.Replace[key]; ok {
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
	f, exists, err := config.DefaultProjectDirFile(keybindsFilename)
	if !exists {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	var global Keybinds

	if _, err := toml.NewDecoder(f).Decode(&global); err != nil {
		return nil, err
	}

	return &global, nil
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

	return &local, nil
}
