package command

func (h *CommandHandler) MoveUp(...string) {
	h.m.Tab.MoveUp()
	h.updateCursorSB()
}

func (h *CommandHandler) MoveDown(...string) {
	h.m.Tab.MoveDown()
	h.updateCursorSB()
}

func (h *CommandHandler) MoveRight(...string) {
	h.m.Tab.MoveRight()
	h.updateCursorSB()
}

func (h *CommandHandler) MoveLeft(...string) {
	h.m.Tab.MoveLeft()
	h.updateCursorSB()
}

func (h *CommandHandler) updateCursorSB() {
	c := h.m.Tab.Cursor()
	h.m.StatusBar.SetCursor(c.Line, c.Position)
	h.m.StatusBar.Draw()
}
