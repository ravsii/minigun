package command

func (h *CommandHandler) OpenFile(args ...string) {
	// TODO: make multiple args open multiple tabs
	if len(args) == 0 {
		h.m.CommandLine.Info(`use ":open <filename>" to open a file`)
		return
	}

	if err := h.m.Tab.FromPath(args[0]); err != nil {
		h.m.CommandLine.Errorf("can't open %s: %s", args[0], err)
		return
	}

	h.m.Tab.Draw()

	h.m.CommandLine.Infof("opened %s", args[0])
}
