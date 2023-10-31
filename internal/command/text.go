package command

import "unicode/utf8"

func (h *CommandHandler) ReplaceSelected(s ...string) {
	if len(s) == 0 {
		return
	}

	r, _ := utf8.DecodeRuneInString(s[0])
	h.m.Tab.ReplaceSelected(r)
}
