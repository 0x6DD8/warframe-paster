package main

import (
	"fmt"
	"strings"

	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
	"golang.design/x/clipboard"
)

var (
	PASTE_SHORTCUT      = []string{"alt", "v"}
	PASTE_CONTENT_LIMIT = 300 // characters
	PASTE_TIMEOUT       = 50  // milliseconds
)

func main() {
	err := clipboard.Init()
	if err != nil {
		panic(err)
	}

	hook.Register(hook.KeyUp, PASTE_SHORTCUT, func(e hook.Event) {
		var content = getClipboardContent()

		if content == "" {
			fmt.Println("Clipboard is empty")
			return
		}

		if len(content) == PASTE_CONTENT_LIMIT {
			fmt.Println("Clipboard content is too long, not pasting.")
			return
		}
		fmt.Println("Pasting content from clipboard: ", content)
		robotgo.Sleep(1)
		robotgo.TypeStr(content, 0, PASTE_TIMEOUT)

	})

	var s = hook.Start()

	println("Listening for paste shortcut:", strings.Join(PASTE_SHORTCUT, " + "))

	<-hook.Process(s)
}

func getClipboardContent() string {
	var content = clipboard.Read(clipboard.FmtText)
	return string(content)
}
