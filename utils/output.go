package utils

import "github.com/fatih/color"

func ErrorOutput(str string) {
	white := color.New(color.FgWhite)
	errorStr := white.Add(color.BgRed)
	_, _ = errorStr.Println(str)
}

func SuccessOutput(str string) {
	white := color.New(color.FgWhite)
	successStr := white.Add(color.BgGreen)
	_, _ = successStr.Println(str)
}

func InfoOutput(str string) {
	white := color.New(color.FgWhite)
	infoStr := white.Add(color.BgBlue)
	_, _ = infoStr.Println(str)
}
