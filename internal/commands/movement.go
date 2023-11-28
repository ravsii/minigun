package commands

import "github.com/ravsii/minigun/internal/mode"

func (h *CommandHandler) MoveUp(...string) {
	h.M.Tab.MoveUp()
	h.updateCursorSB()
}

func (h *CommandHandler) MoveDown(...string) {
	h.M.Tab.MoveDown()
	h.updateCursorSB()
}

func (h *CommandHandler) MoveLeft(...string) {
	switch mode.Current() {
	case mode.View:
		h.M.Tab.MoveLeft()
		h.updateCursorSB()
	case mode.Command:
		h.M.CommandLine.MoveLeft()
	}
}

func (h *CommandHandler) MoveRight(...string) {
	switch mode.Current() {
	case mode.View:
		h.M.Tab.MoveRight()
		h.updateCursorSB()
	case mode.Command:
		h.M.CommandLine.MoveRight()
	}
}

func (h *CommandHandler) JumpLineStart(...string) {
	h.M.Tab.MoveLeft()
	h.updateCursorSB()
}

func (h *CommandHandler) JumpLineEnd(...string) {
	h.M.Tab.MoveLeft()
	h.updateCursorSB()
}

func (h *CommandHandler) updateCursorSB() {
	c := h.M.Tab.Cursor()
	h.M.StatusBar.SetCursor(c.Line, c.Position)
	h.M.StatusBar.Draw()
}
