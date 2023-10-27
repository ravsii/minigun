package cmd

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
)

var (
	errStyle  = tcell.StyleDefault.Foreground(tcell.ColorRed)
	infoStyle = tcell.StyleDefault.Foreground(tcell.ColorBlue)
)

func Error(msg string) {
	command.printStyled("Error: "+msg, errStyle, -1)
}

func Errorf(format string, args ...interface{}) {
	command.printStyled("Error: "+fmt.Sprintf(format, args...), errStyle, -1)
}

func Info(msg string) {
	command.printStyled("Info: "+msg, infoStyle, -1)

}
