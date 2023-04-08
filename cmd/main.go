package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/skykosiner/find-me-daddy/pkg/clip"
	"golang.design/x/clipboard"
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
		searcher := os.Getenv("FIND_ME_DADDY_FUZZY")

		// If the user has not sellected anything just use dmenu (this only
		// works on Linux, as Mac L, and windows even bigger L)
		if searcher == "" {
			searcher = "dmenu"
		}

		fmt.Println("seacher", searcher)

		// TODO: WHY WON"T THIS WORK???????
		cmd := fmt.Sprintf("echo '%s' | sed 's/\\[//' | sed 's/\\]//' | sed 's/^$//' | %s ", items, searcher)
		output := exec.Command("bash", "-c", cmd)

		stdOut, err := output.Output()

		if err != nil {
			log.Fatal("Error with finding the daddy ", err)
		}

		err = clipboard.Init()
		if err != nil {
			log.Fatal("Error with clipboard ", err)
		}

		fmt.Println(string(stdOut))
		clipboard.Write(clipboard.FmtText, []byte("test"))
	}
}
