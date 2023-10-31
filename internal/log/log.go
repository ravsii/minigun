package log

import (
	"log/slog"
	"os"
	"path"
)

// Init creates and write to the file at a given filepath.
// if fPath is empty, $HOME/.config/minigun/debug.log is used.
func Init(fPath string) {
	if fPath == "" {
		fPath = defaultPath()
	}

	fDir := path.Dir(fPath)

	err := os.MkdirAll(fDir, 0755)
	if err != nil {
		panic(err)
	}

	logF, err := os.OpenFile(fPath, os.O_CREATE|os.O_RDWR, 0755)
	if err != nil {
		panic(err)
	}

	l := slog.New(slog.NewTextHandler(logF, nil))
	slog.SetDefault(l)
}

func defaultPath() string {
	home := os.Getenv("HOME")
	if home == "" {
		var err error
		home, err = os.Getwd()
		if err != nil {
			panic(err)
		}
	}

	return path.Join(home, ".config", "minigun", "debug.log")
}
