package cmdline

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
)

var (
	errStyle  = tcell.StyleDefault.Foreground(tcell.NewHexColor(0xFF0000).TrueColor()).Background(tcell.ColorBlack.TrueColor())
	infoStyle = tcell.StyleDefault.Foreground(tcell.NewHexColor(0x00FF00).TrueColor()).Background(tcell.ColorBlack.TrueColor())
)

func (c *CommandLine) Error(msg string) {
	c.printStyled("Error: "+msg, errStyle, -1)
}

func (c *CommandLine) Errorf(format string, args ...interface{}) {
	c.Error(fmt.Sprintf(format, args...))
}

func (c *CommandLine) Info(msg string) {
	c.printStyled("Info: "+msg, infoStyle, -1)
}

func (c *CommandLine) Infof(format string, args ...interface{}) {
	c.Info(fmt.Sprintf(format, args...))
}
