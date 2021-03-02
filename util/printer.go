package util

import (
	"fmt"
	"github.com/fatih/color"
)

func PrintlnError(err error) {
	fmt.Printf("%s %v\n", color.RedString("Error:"), err)
}

func PrintlnInfo(info string) {
	fmt.Printf("%v: %v\n", color.CyanString("Canal"), info)
}
