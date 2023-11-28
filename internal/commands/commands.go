package commands

import (
	"os"
	"strings"

	"github.com/ravsii/minigun/internal/minigun"
)

// Cmd a function type for ANY command possible inside the minigun.
type Cmd func(args ...string)

// var commands = map[string]Cmd{
// 	"q":    Quit,
// 	"quit": Quit,
// }

type CommandHandler struct {
	aliases map[string]string
	cmds    map[string]Cmd

	M *minigun.Minigun
}

func New(mg *minigun.Minigun) CommandHandler {
	handler := CommandHandler{M: mg}
	handler.cmds = map[string]Cmd{
		"clear":               handler.ClearCommandLine,
		"command_delete_rune": handler.CommandDeleteRune,
		"command_submit":      handler.CommandSubmit,
		"enter_command_mode":  handler.EnterCommandMode,
		"enter_edit_mode":     handler.EnterEditMode,
		"enter_replace_mode":  handler.EnterReplaceMode,
		"enter_view_mode":     handler.EnterViewMode,
		"execute":             handler.CmdExecute,
		"noop":                func(...string) {}, // bind-remover
		"open":                handler.OpenFile,
		"quit":                handler.Quit,
		"tab_delete_rune":     handler.TabDeleteRune,
		"tab_jump_line_end":   handler.TabJumpLineEnd,
		"tab_jump_line_start": handler.TabJumpLineStart,
		"tab_move_down":       handler.TabMoveDown,
		"tab_move_left":       handler.TabMoveLeft,
		"tab_move_right":      handler.TabMoveRight,
		"tab_move_up":         handler.TabMoveUp,
		"write":               handler.WriteFile,
	}
	handler.aliases = map[string]string{
		"o": "open",
		"q": "quit",
		"w": "write",
	}

	for k, v := range handler.aliases {
		if _, ok := handler.cmds[v]; !ok {
			panic("unknown alias: " + v + " for " + k)
		}
	}

	return handler
}

// CmdFromString returns cmd from a given (presumably user-inputted) string.
// Returns nil, false if a command wasn't found.
func (h *CommandHandler) CmdFromString(c string) (Cmd, bool) {
	c = strings.ToLower(c)
	if alias, exists := h.aliases[c]; exists {
		c = alias
	}
	if cmd, exists := h.cmds[c]; exists {
		return cmd, true
	}
	return nil, false
}

// CmdExecute is a general command used to call other commands.
// First argument must be the name of the command.
//
// If a command doesn't exist - it's a noop.
func (h *CommandHandler) CmdExecute(args ...string) {
	if len(args) == 0 {
		return
	}

	cmd, found := h.CmdFromString(args[0])
	if !found {
		badCmd := args[0]
		if len(args) > 1 {
			badCmd += "(" + strings.Join(args[1:], " ") + ")"
		}
		h.M.CommandLine.Errorf("unknown command: %s", badCmd)
		return
	}

	if len(args) > 1 {
		cmd(args[1:]...)
		return
	}
	cmd()
}

func (h *CommandHandler) ClearCommandLine(...string) {
	h.M.CommandLine.Draw()
}

func (h *CommandHandler) Info(s ...string) {
	h.M.CommandLine.Info(strings.Join(s, " "))
}

func (h *CommandHandler) Quit(...string) {
	os.Exit(0)
}
