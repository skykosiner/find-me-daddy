package main

import (
	"fmt"
	"os"

	"github.com/yonikosiner/find-me-daddy/pkg/clipboard"
)

func main() {
	var ui *clipboard.Ui = &clipboard.Ui{}
	var clip *clipboard.Clipboard = &clipboard.Clipboard{}

	if len(os.Args[1:]) <= 0 {
		fmt.Println("No command line arguments passed in")
		return
	}

	args := os.Args[1:]

	switch args[0] {
	case "--add":
		clip.AddToHistory()
	case "--search":
		ui.InitUi()
	}
}
