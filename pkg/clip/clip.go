package clip

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/atotto/clipboard"
)

type Clipboard struct {
	Contents string `json:"contents"`
}

func (c *Clipboard) GetCurrentClip() Clipboard {
	clip, err := clipboard.ReadAll()

	if err != nil {
		log.Fatal("There was an error getting the clipboard ", err)
	}

	replacer := strings.NewReplacer("\t", "", "\n", "")

	return Clipboard{replacer.Replace(string(clip))}
}

func (c *Clipboard) GetHistory() []Clipboard {
	var History []Clipboard
	path := fmt.Sprintf("%s/.local/share/clipboard.json", os.Getenv("HOME"))
	bytes, err := ioutil.ReadFile(path)

	if err != nil {
		log.Fatal("Error getting clipboard history", err)
	}

	err = json.Unmarshal(bytes, &History)

	if err != nil {
		log.Fatal("Error unmarshling json history ", err)
	}

	return History
}

func (c *Clipboard) AddToHistory() {
	var clipboards []Clipboard
	path := fmt.Sprintf("%s/.local/share/clipboard.json", os.Getenv("HOME"))
	history := c.GetHistory()
	currentClip := c.GetCurrentClip()

	// Make sure the current item isn't already stored
	for _, item := range history {
		if item == currentClip {
			return
		}
	}

	file, err := os.OpenFile(path, os.O_RDWR, 0644)

	if err != nil {
		log.Fatal("Error reading history file", err)
	}

	defer file.Close()

	bytes, err := ioutil.ReadAll(file)

	if err != nil {
		log.Fatal("Erorr reading history file", err)
	}

	err = json.Unmarshal(bytes, &clipboards)
	if err != nil {
		log.Fatal(err)
	}

	clipboards = append(clipboards, currentClip)

	newBytes, err := json.MarshalIndent(clipboards, "", "	")
	if err != nil {
		log.Fatal(err)
	}

	_, err = file.WriteAt(newBytes, 0)
	if err != nil {
		log.Fatal(err)
	}
}

func (c *Clipboard) GetItems() []string {
	var items []string
	for _, item := range c.GetHistory() {
		items = append(items, fmt.Sprintf("%s\n", item.Contents))
	}

	return items
}
