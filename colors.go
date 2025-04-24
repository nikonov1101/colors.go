package colors

import (
	"os"
)

const (
	tblack   = "\033[0;30m"
	tred     = "\033[0;31m"
	tgreen   = "\033[0;32m"
	tyellow  = "\033[0;33m"
	tblue    = "\033[0;34m"
	tmagenta = "\033[0;35m"
	tcyan    = "\033[0;36m"
	twhite   = "\033[0;37m"
	tgray    = "\033[0;90m"

	btblack   = "\033[1;30m"
	btred     = "\033[1;31m"
	btgreen   = "\033[1;32m"
	btyellow  = "\033[1;33m"
	btblue    = "\033[1;34m"
	btmagenta = "\033[1;35m"
	btcyan    = "\033[1;36m"
	btwhite   = "\033[1;37m"
	bgray     = "\033[1;90m"

	noColor = "\033[0m" // no color
)

func colorize(color, s string) string {
	// follow the https://no-color.org/ spec.
	if len(os.Getenv("NO_COLOR")) > 0 {
		return s
	}
	return color + s + noColor
}

func Black(s string) string {
	return colorize(tblack, s)
}

func Red(s string) string {
	return colorize(tred, s)
}

func Green(s string) string {
	return colorize(tgreen, s)
}

func Yellow(s string) string {
	return colorize(tyellow, s)
}

func Blue(s string) string {
	return colorize(tblue, s)
}

func Magenta(s string) string {
	return colorize(tmagenta, s)
}

func Cyan(s string) string {
	return colorize(tcyan, s)
}

func White(s string) string {
	return colorize(twhite, s)
}

func Gray(s string) string {
	return colorize(tgray, s)
}

func BBlack(s string) string {
	return colorize(btblack, s)
}

func BRed(s string) string {
	return colorize(btred, s)
}

func BGreen(s string) string {
	return colorize(btgreen, s)
}

func BYellow(s string) string {
	return colorize(btyellow, s)
}

func BBlue(s string) string {
	return colorize(btblue, s)
}

func BMagenta(s string) string {
	return colorize(btmagenta, s)
}

func BCyan(s string) string {
	return colorize(btcyan, s)
}

func BWhite(s string) string {
	return colorize(btwhite, s)
}

func BGray(s string) string {
	return colorize(bgray, s)
}
