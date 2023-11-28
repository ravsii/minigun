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

func (h *CommandHandler) InsertRuneCmd(s ...string) {
	if len(s) == 0 {
		return
	}

	runes := make([]rune, 0, len(s))
	for _, str := range s {
		runes = append(runes, []rune(str)...)
	}

	h.M.Tab.InsertRune(runes...)
}

func (h *CommandHandler) InsertRune(r rune) {
	h.M.Tab.InsertRune(r)
}

func (h *CommandHandler) InsertNewLine() {
	h.M.Tab.InsertNewLine()
}
