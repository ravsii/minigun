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
	h.m.StatusBar.SetCursor(h.m.Tab.Cursor().Line, h.m.Tab.Cursor().Position)
	h.m.StatusBar.Draw()
}
