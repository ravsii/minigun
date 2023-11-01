package log

import (
	"log/slog"
	"os"
	"path"

	"github.com/ravsii/minigun/internal/config"
)

const defaultName = "debug.log"

// Init creates and write to the file at a given filepath.
// if fPath is empty, DefaultProjectDir()/debug.log is used.
func Init(fPath string) {
	if fPath == "" {
		var ok bool
		fPath, ok = config.DefaultProjectDir()
		if !ok {
			panic("can't create default project dir")
		}

		fPath = path.Join(fPath, defaultName)
	}

	logF, err := os.OpenFile(fPath, os.O_CREATE|os.O_RDWR, 0755)
	if err != nil {
		panic(err)
	}

	slog.SetDefault(slog.New(slog.NewTextHandler(
		logF,
		&slog.HandlerOptions{Level: slog.LevelDebug},
	)))
}
