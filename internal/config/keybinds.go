package config

import (
	"strings"

	"github.com/ravsii/minigun/internal/mode"
	"gopkg.in/yaml.v3"
)

const keybindsFilename = "keybinds.yaml"

var keybinds keybindsConfig

type keybindsConfig struct {
	// View stores view-mode keybinds.
	View map[string]string `yaml:"view"`
}

func (c *keybindsConfig) Merge(new *keybindsConfig) {
	for key, cmd := range new.View {
		c.View[strings.ToLower(key)] = cmd
	}
}

func Load() error {
	keybinds.View = make(map[string]string)

	loadGlobal()
	return loadLocal()
}

func CommandFor(m mode.Mode, key string) (string, bool) {
	var cmd string
	switch m {
	case mode.View:
		if bind, ok := keybinds.View[key]; ok {
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

func loadGlobal() {
	// TODO: implement
}

func loadLocal() error {
	binds, exists, err := ProjectSpecificDirFile(keybindsFilename)
	if !exists {
		return nil
	}

	if err != nil {
		return err
	}

	var local keybindsConfig

	if err := yaml.NewDecoder(binds).Decode(&local); err != nil {
		return err
	}

	keybinds.Merge(&local)

	return nil
}
