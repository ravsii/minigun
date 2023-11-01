package config

import (
	"os"
	"path"
)

const configDirPerms = 0755

// DefaultProjectDir returns $HOME/.config/minigun, creating it if not exists,
// returning false if $HOME is empty
func DefaultProjectDir() (string, bool) {
	home := os.Getenv("HOME")
	if home == "" {
		return "", false
	}

	dir := path.Join(home, ".config", "minigun")
	err := os.MkdirAll(dir, configDirPerms)
	if err != nil {
		panic(err)
	}

	return dir, true
}

func ProjectSpecificDir() (string, error) {
	home, err := os.Getwd()
	if err != nil {
		return "", err
	}

	return path.Join(home, ".minigun"), nil
}
