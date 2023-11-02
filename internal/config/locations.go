package config

import (
	"os"
	"path"
)

const configPerms = 0755

// DefaultProjectDir returns $HOME/.config/minigun, creating it if not exists,
// returning false if $HOME is empty
func DefaultProjectDir() (string, bool) {
	home := os.Getenv("HOME")
	if home == "" {
		return "", false
	}

	dir := path.Join(home, ".config", "minigun")
	err := os.MkdirAll(dir, configPerms)
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

func DefaultProjectDirFile(filePath ...string) (*os.File, bool, error) {
	dir, ok := DefaultProjectDir()
	if !ok {
		return nil, false, nil
	}

	return open(dir, filePath...)
}

func ProjectSpecificDirFile(filePath ...string) (*os.File, bool, error) {
	dir, err := ProjectSpecificDir()
	if err != nil {
		return nil, false, err
	}

	return open(dir, filePath...)
}

func open(dir string, filePath ...string) (*os.File, bool, error) {
	p := path.Join(append([]string{dir}, filePath...)...)

	_, err := os.Stat(p)
	if err != nil {
		return nil, false, err
	}

	f, err := os.OpenFile(p, os.O_RDONLY, configPerms)
	if err != nil {
		return nil, false, err
	}

	return f, true, nil
}
