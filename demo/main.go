package main

import (
	"fmt"

	"github.com/nikonov1101/colors.go"
)

func main() {
	fmt.Println()
	fmt.Println(colors.Black("normal_black"))
	fmt.Println(colors.Red("normal_red"))
	fmt.Println(colors.Green("normal_green"))
	fmt.Println(colors.Yellow("normal_yellow"))
	fmt.Println(colors.Blue("normal_blue"))
	fmt.Println(colors.Magenta("normal_magenta"))
	fmt.Println(colors.Cyan("normal_cyan"))
	fmt.Println(colors.White("normal_white"))

	fmt.Println()
	fmt.Println(colors.BBlack("bright_black"))
	fmt.Println(colors.BRed("bright_red"))
	fmt.Println(colors.BGreen("bright_green"))
	fmt.Println(colors.BYellow("bright_yellow"))
	fmt.Println(colors.BBlue("bright_blue"))
	fmt.Println(colors.BMagenta("bright_magenta"))
	fmt.Println(colors.BCyan("bright_cyan"))
	fmt.Println(colors.BWhite("bright_white"))
}
