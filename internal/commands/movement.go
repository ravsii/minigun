package commands

import "github.com/ravsii/minigun/internal/mode"

func (h *CommandHandler) TabMoveUp(...string) {
	h.M.Tab.MoveUp()
	h.updateCursorSB()
}

func (h *CommandHandler) TabMoveDown(...string) {
	h.M.Tab.MoveDown()
	h.updateCursorSB()
}

func (h *CommandHandler) TabMoveLeft(...string) {
	switch mode.Current() {
	case mode.View:
		h.M.Tab.MoveLeft()
		h.updateCursorSB()
	case mode.Command:
		h.M.CommandLine.MoveLeft()
	}
}

func (h *CommandHandler) TabMoveRight(...string) {
	switch mode.Current() {
	case mode.View:
		h.M.Tab.MoveRight()
		h.updateCursorSB()
	case mode.Command:
		h.M.CommandLine.MoveRight()
	}
}

func (h *CommandHandler) TabJumpLineStart(...string) {
	h.M.Tab.MoveLeft()
	h.updateCursorSB()
}

func (h *CommandHandler) TabJumpLineEnd(...string) {
	h.M.Tab.MoveLeft()
	h.updateCursorSB()
}

func (h *CommandHandler) updateCursorSB() {
	c := h.M.Tab.Cursor()
	h.M.StatusBar.SetCursor(c.Line, c.Position)
	h.M.StatusBar.Draw()
}
