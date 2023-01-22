package main

import (
	"github.com/yonikosiner/find-me-daddy/pkg/clipboard"
)

func main() {
	var c *clipboard.Clipboard = &clipboard.Clipboard{}
	c.AddToHistory()
}
