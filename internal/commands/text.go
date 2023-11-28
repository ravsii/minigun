package commands

import "unicode/utf8"

func (h *CommandHandler) ReplaceRune(s ...string) {
	if len(s) == 0 {
		return
	}

	r, _ := utf8.DecodeRuneInString(s[0])
	h.M.Tab.ReplaceRune(r)
}

func (h *CommandHandler) TabDeleteRune(s ...string) {
	h.M.Tab.DeleteRune()
}

func (h *CommandHandler) InsertRune(s ...string) {
	if len(s) == 0 {
		return
	}

	r, _ := utf8.DecodeRuneInString(s[0])
	h.M.Tab.InsertRune(r)
}
