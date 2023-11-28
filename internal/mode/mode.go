package mode

type Mode int

const (
	View Mode = iota
	Command
	Replace
	Edit
)

var modeString = map[Mode]string{
	View:    "View",
	Command: "Console",
	Replace: "Replace",
	Edit:    "Edit",
}

var current = View

func Set(m Mode) {
	current = m
}

func Current() Mode {
	return current
}

func (m Mode) String() string {
	return modeString[m]
}
