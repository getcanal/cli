package util

import (
	"fmt"
	"github.com/fatih/color"
)

func PrintError(err error) {
	fmt.Printf("%s %v\n", color.RedString("error:"), err)
}
