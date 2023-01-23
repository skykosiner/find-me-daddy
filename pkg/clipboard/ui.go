package clipboard

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/manifoldco/promptui"
)

type Ui struct {
	items []string
}

func (u *Ui) InitUi() {
	var c *Clipboard = &Clipboard{}
	for _, value := range c.GetHistory() {
		u.items = append(u.items, value.Contents)
	}

	prompt := promptui.Select{
		Label: "Select Item",
		Items: u.items,
	}

	_, result, err := prompt.Run()

	if err != nil {
		log.Fatal("Prompt failed", err)
	}

	u.CopyItem(result)
}

func (u *Ui) CopyItem(result string) {
	//TODO: Figure out why it's only being coppied when stil in golang app????
	cmd := fmt.Sprintf("echo '%s' | xclip -sel c", result)
	setClip := exec.Command("bash", "-c", cmd)
	_, err := setClip.Output()

	if err != nil {
		log.Fatal("Error setting clipboard to item", err)
	}
}
