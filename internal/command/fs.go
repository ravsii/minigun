package command

import (
	"io"
	"os"
	"strings"
)

func (h *CommandHandler) OpenFile(filePaths ...string) {
	// TODO: make multiple args open multiple tabs
	if len(filePaths) == 0 {
		h.m.CommandLine.Info(`use ":open <filepath>" to open a file`)
		return
	}

	if err := h.m.Tab.FromPath(filePaths[0]); err != nil {
		h.m.CommandLine.Errorf("can't open %s: %s", filePaths[0], err)
		return
	}

	h.m.CommandLine.Infof("opened %s", filePaths[0])
}

// WriteFile saves current tab to a file. (0th arg)
// If no args present, saves file to the current tab path.
func (h *CommandHandler) WriteFile(filePath ...string) {
	var fPath string
	if len(filePath) == 0 {
		fPath = h.m.Tab.Path()
	} else {
		fPath = filePath[0]
	}

	if fPath == "" {
		h.m.CommandLine.Errorf(`use ":write <filepath>" to save a file`)
		return
	}

	f, err := os.Create(fPath)
	if err != nil {
		h.m.CommandLine.Errorf("can't save to %s: %s", fPath, err)
		return
	}

	r := strings.NewReader(h.m.Tab.AsString())

	if _, err := io.Copy(f, r); err != nil {
		h.m.CommandLine.Errorf("can't write to %s: %s", fPath, err)
		return
	}

	h.m.CommandLine.Infof("%s saved", fPath)
}
