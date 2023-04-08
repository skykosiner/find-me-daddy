package utils

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func GetClipCommand() string {
	var cmd string
	osSystem := exec.Command("uname")
	stdOout, err := osSystem.Output()

	if err != nil {
		log.Fatal("Error getting os of computer. You're most likley on Windows. L (this doesn't work on windows, unless you want to make a pull request to make it work)", err)
	}

	if strings.TrimSuffix(string(stdOout), "\n") == "Linux" {
		cmd =  "xclip -in -selection clipboard"
		fmt.Println("aeou")
	} else {
		cmd = "pbcopy"
	}

	return cmd
}
