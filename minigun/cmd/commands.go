package cmd

import (
	"os"
	"strings"
)

// Cmd a function type for ANY command possible inside the minigun.
type Cmd func(args ...string)

var commands = map[string]Cmd{
	"q":    Quit,
	"quit": Quit,
}

// Execute is a general command used to call other commands.
// First argument must be a command name.
//
// If a command doesn't exist or no args passed - it's a noop.
func Execute(args ...string) {
	if len(args) == 0 {
		return
	}

	cmdName := strings.ToLower(args[0])
	cmd, ok := commands[cmdName]
	if !ok {
		Errorf("unknown command: %s", cmdName)
		return
	}

	if len(args) > 1 {
		cmd(args[1:]...)
		return
	}
	cmd()
}

func Quit(...string) {
	os.Exit(0)
}
