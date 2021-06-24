package colr

import "fmt"

type Color int

const (
	None     Color = 0
	RedC     Color = 31
	GreenC   Color = 32
	YellowC  Color = 33
	BlueC    Color = 34
	MagentaC Color = 35
	CyanC    Color = 36
	WhiteC   Color = 37
)

func Colorized(c Color, content interface{}) string {
	return fmt.Sprintf("\x1b[%dm%v\x1b[%dm", c, content, None)
}

func Red(content interface{}) string {
	return Colorized(RedC, content)
}

func Green(content interface{}) string {
	return Colorized(GreenC, content)
}

func Yellow(content interface{}) string {
	return Colorized(YellowC, content)
}

func Blue(content interface{}) string {
	return Colorized(BlueC, content)
}

func Magenta(content interface{}) string {
	return Colorized(MagentaC, content)
}

func Cyan(content interface{}) string {
	return Colorized(CyanC, content)
}

func White(content interface{}) string {
	return Colorized(WhiteC, content)
}
