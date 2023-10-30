package config

import (
	"errors"
	"io/fs"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

const keybindsLocalPath = ".minigun/keybinds.yaml"

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

func loadGlobal() {
	// TODO: implement
}

func loadLocal() error {
	if _, err := os.Stat(keybindsLocalPath); errors.Is(err, fs.ErrNotExist) {
		return nil
	}

	keybindsFile, err := os.Open(keybindsLocalPath)
	if err != nil {
		return err
	}

	var local keybindsConfig

	if err := yaml.NewDecoder(keybindsFile).Decode(&local); err != nil {
		return err
	}

	keybinds.Merge(&local)

	return nil
}
