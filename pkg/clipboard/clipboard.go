package clipboard

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

type Clipboard struct {
	Contents string `json:"Contents"`
}

func (c *Clipboard) GetCurrentClip() Clipboard {
	currentClip := exec.Command("xclip", "-o", "-selection", "clipboard")
	stdOut, err := currentClip.Output()

	if err != nil {
		log.Fatal("Error getting current clipboard", err)
	}

	return Clipboard{string(stdOut)}
}

func (c *Clipboard) GetHistory() []Clipboard {
	var history []Clipboard
	filePath := fmt.Sprintf("%s/.local/clipboard.json", os.Getenv("HOME"))

	bytes, err := ioutil.ReadFile(filePath)

	if err != nil {
		log.Fatal("There was an error reading the history file", err)
	}

	err = json.Unmarshal(bytes, &history)

	if err != nil {
		log.Fatal("There was an error unmarshaling the json", err)
	}

	return history
}

func (c *Clipboard) AddToHistory() {
	var cilpboards []Clipboard
	filePath := fmt.Sprintf("%s/.local/clipboard.json", os.Getenv("HOME"))
	history := c.GetHistory()
	currentClip := c.GetCurrentClip()

	for _, value := range history {
		if currentClip.Contents == value.Contents {
			return
		}
	}

	file, err := os.OpenFile(filePath, os.O_RDWR, 0644)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(bytes, &cilpboards)
	if err != nil {
		log.Fatal(err)
	}

	cilpboards = append(cilpboards, currentClip)

	newBytes, err := json.MarshalIndent(cilpboards, "", "    ")
	if err != nil {
		panic(err)
	}

	_, err = file.WriteAt(newBytes, 0)
	if err != nil {
		panic(err)
	}
}
