package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/skykosiner/find-me-daddy/pkg/clip"
	"github.com/skykosiner/find-me-daddy/pkg/utils"
)

func main() {
	var c *clip.Clipboard

	args := os.Args[1:]

	switch args[0] {
	case "--add":
		c.AddToHistory()
	case "--get-list":
		for _, item := range c.GetItems() {
			fmt.Println(item)
		}
	case "--get-fuzzy":
		items := c.GetItems()
		clipCommand := utils.GetClipCommand()
		searcher := os.Getenv("FIND_ME_DADDY_FUZZY")
		fmt.Println(searcher)

		// If the user has not sellected anything just use dmenu (this only
		// works on Linux, as Mac L, and windows even bigger L)
		if searcher == "" {
			searcher = "dmenu"
		}

		// TODO: WHY WON"T THIS WORK???????
		cmd := fmt.Sprintf("echo '%s' | sed 's/\\[//' | sed 's/\\]//' | sed 's/^$//' | %s | %s", items, searcher, clipCommand)
		output := exec.Command("bash", "-c", cmd)

		_, err := output.Output()

		if err != nil {
			log.Fatal("Erro with finding the daddy", err)
		}
	}
}
