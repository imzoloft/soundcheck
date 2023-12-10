package inout

import (
	"fmt"
	"os"

	"github.com/imzoloft/soundcheck/config"
)

func PrintError(errorMessage string, errorProp string) {
	fmt.Printf("[%s!%s] %s%s%s  | %s\n", config.TextRed, config.TextReset, config.TextRed, errorMessage, config.TextReset, errorProp)
}

func FatalError(errorMessage string, errorProp string) {
	fmt.Printf("[%s!%s] %s%s%s  | %s\n", config.TextRed, config.TextReset, config.TextRed, errorMessage, config.TextReset, errorProp)
	os.Exit(1)
}
