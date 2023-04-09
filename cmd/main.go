package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/atotto/clipboard"
	"github.com/ktr0731/go-fuzzyfinder"
	"github.com/skykosiner/find-me-daddy/pkg/clip"
)

func main() {
	var c *clip.Clipboard

	args := os.Args[1:]

	switch args[0] {
	case "-add":
		c.AddToHistory()
	case "-get-list":
		for _, item := range c.GetItems() {
			fmt.Println(item)
		}
	case "-get-fuzzy":
		items := c.GetItems()
		searcher := os.Getenv("FIND_ME_DADDY_FUZZY")

		if searcher == "" {
			searcher = "dmenu"
		}

		if len(args) > 1 {
			switch args[1] {
			case "-f":
				searcher = "fzf"
			case "-d":
				searcher = "dmenu"
			case "-r":
				searcher = "rofi -dmenu"
			default:
				searcher := os.Getenv("FIND_ME_DADDY_FUZZY")

				if searcher == "" {
					searcher = "dmenu"
				}
			}
		}

		if searcher == "fzf" {
			idx, err := fuzzyfinder.FindMulti(
				items,
				func(i int) string {
					return items[i]
				},
			)

			if err != nil {
				log.Fatal("Error fuzzy finding", err)
			}

			err = clipboard.WriteAll(items[idx[0]])
			if err != nil {
				log.Fatal("Error wirting to clipboard", err)
			}
		} else {
			cmd := fmt.Sprintf("echo '%s' | sed 's/\\[//' | sed 's/\\]//' | sed 's/^$//' | %s ", items, searcher)
			output := exec.Command("bash", "-c", cmd)

			stdOut, err := output.Output()

			if err != nil {
				log.Fatal("Error with finding the daddy ", err)
			}

			err = clipboard.WriteAll(string(stdOut))
			if err != nil {
				log.Fatal("Error wirting to clipboard", err)
			}
		}

	}
}
