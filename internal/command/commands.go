package command

import (
	"os"
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/ravsii/minigun/internal/minigun"
	"github.com/ravsii/minigun/internal/mode"
	"github.com/ravsii/minigun/internal/screen"
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

	m *minigun.Minigun
}

func New(mg *minigun.Minigun) CommandHandler {
	handler := CommandHandler{m: mg}
	handler.cmds = map[string]Cmd{
		"enter_command_mode": handler.EnterCommandMode,
		"execute":            handler.CmdExecute,
		"move_down":          handler.MoveDown,
		"move_left":          handler.MoveLeft,
		"move_right":         handler.MoveRight,
		"move_up":            handler.MoveUp,
		"noop":               func(...string) {}, // bind-remover
		"open":               handler.OpenFile,
		"quit":               handler.Quit,
	}
	handler.aliases = map[string]string{
		"o": "open",
		"q": "quit",
	}

	for k, v := range handler.aliases {
		if _, ok := handler.cmds[v]; !ok {
			panic("unknown alias: " + v + " for " + k)
		}
	}

	return handler
}

// CmdFromString returns cmd from a given (presumably user-inputted) string.
// Returns nil, false if command is not found.
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
		h.m.CommandLine.Errorf("unknown command: %s", args[0])
		return
	}

	if len(args) > 1 {
		cmd(args[1:]...)
		return
	}
	cmd()
}

func (h *CommandHandler) EnterCommandMode(...string) {
	mode.Set(mode.Console)
	h.m.StatusBar.Draw()
	screen.Screen().SetCursorStyle(tcell.CursorStyleBlinkingBar)
	defer func() {
		mode.Set(mode.View)
		screen.Screen().SetCursorStyle(tcell.CursorStyleBlinkingBlock)
		screen.Screen().HideCursor()
		h.m.StatusBar.Draw()
	}()

	input := h.m.CommandLine.HandleUserInput()
	if input != "" {
		h.CmdExecute(strings.Split(input, " ")...)
	}
}

func (h *CommandHandler) ClearCommandLine(...string) {
	h.m.CommandLine.Draw()
}

func (h *CommandHandler) Info(s ...string) {
	h.m.CommandLine.Info(strings.Join(s, " "))
}

func (h *CommandHandler) Quit(...string) {
	os.Exit(0)
}
