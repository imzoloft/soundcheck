package inout

import (
	"fmt"

	"github.com/imzoloft/soundcheck/config"
	"github.com/imzoloft/soundcheck/model"
)

func ShowBanner() {
	fmt.Printf(`
	%s%s
	%s
	%s
	%s%s%s%s%s
	%s`,
		config.TextBlue,
		"╔═╗╔═╗╦ ╦╔╗╔╔╦╗╔═╗╦ ╦╔═╗╔═╗╦╔═╔═╗╦═╗",
		"╚═╗║ ║║ ║║║║ ║║║  ╠═╣║╣ ║  ╠╩╗║╣ ╠╦╝",
		"╚═╝╚═╝╚═╝╝╚╝═╩╝╚═╝╩ ╩╚═╝╚═╝╩ ╩╚═╝╩╚═",
		config.TextReset,
		"     	     Mode: ", config.TextBlue, config.ProxyType, config.TextReset,
		"         gitlab.com/zoloft\n\n",
	)
}

func ClearScreen() {
	fmt.Print("\033[H\033[2J")
}

func SelectProxyType() {
	fmt.Printf("[%s?%s] %s | %s%s%s%s | %s%s%s%s | %s%s%s%s | %s%s%s%s\n",
		config.TextBlue, config.TextReset,
		"Select proxy type",
		config.TextBlue, "1", config.TextReset, " HTTPS",
		config.TextBlue, "2", config.TextReset, " SOCKS4",
		config.TextBlue, "3", config.TextReset, " SOCKS5",
		config.TextBlue, "4", config.TextReset, " NONE")

	var proxyType int
	fmt.Scanln(&proxyType)

	switch proxyType {
	case 1:
		config.ProxyType = "https"
	case 2:
		config.ProxyType = "socks4"
	case 3:
		config.ProxyType = "socks5"
	case 4:
		config.ProxyType = "none"
	default:
		FatalError("INVALID PROXY TYPE", "Please select a valid proxy type")
	}
}

func SelectRoutineNumber() int {
	fmt.Printf("[%s?%s] %s: ",
		config.TextBlue, config.TextReset,
		"Select number of threads")

	var routineNumber int
	fmt.Scanln(&routineNumber)

	return routineNumber
}

func DisplayResult(response model.Response) string {
	var (
		statusIcon string
		statusText string
		colorOpen  string
		fileName   string
	)

	if response.Status == "available" {
		statusIcon = "!"
		statusText = "BAD"
		colorOpen = config.TextRed
		fileName = "list/bad.txt"
	} else {
		statusIcon = "+"
		statusText = "HIT"
		colorOpen = config.TextGreen
		fileName = "list/hit.txt"
	}

	fmt.Printf("[%s%s%s] %s%s%s | %s\n", colorOpen, statusIcon, config.TextReset, colorOpen, statusText, config.TextReset, response.Identifier)

	return fileName

}
