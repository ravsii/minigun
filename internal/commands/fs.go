package commands

import (
	"io"
	"os"
	"strings"
)

func (h *CommandHandler) OpenFile(filePaths ...string) {
	// TODO: make multiple args open multiple tabs
	if len(filePaths) == 0 {
		h.M.CommandLine.Info(`use ":open <filepath>" to open a file`)
		return
	}

	if err := h.M.Tab.FromPath(filePaths[0]); err != nil {
		h.M.CommandLine.Error(err.Error())
		return
	}

	h.M.CommandLine.Infof("opened %s", filePaths[0])
}

// WriteFile saves current tab to a file. (0th arg)
// If no args present, saves file to the current tab path.
func (h *CommandHandler) WriteFile(filePath ...string) {
	var fPath string
	if len(filePath) == 0 {
		fPath = h.M.Tab.Path()
	} else {
		fPath = filePath[0]
	}

	if fPath == "" {
		h.M.CommandLine.Errorf(`use ":write <filepath>" to save a file`)
		return
	}

	f, err := os.Create(fPath)
	if err != nil {
		h.M.CommandLine.Errorf("can't save to %s: %s", fPath, err)
		return
	}

	r := strings.NewReader(h.M.Tab.AsString())

	if _, err := io.Copy(f, r); err != nil {
		h.M.CommandLine.Errorf("can't write to %s: %s", fPath, err)
		return
	}

	h.M.CommandLine.Infof("%s saved", fPath)
}
