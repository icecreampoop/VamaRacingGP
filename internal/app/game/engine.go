package game

import (
	"fmt"
	"strings"
)

var (
	userInput string = ""
)

func GameStart() {
	fmt.Println(WELCOME)

	for userInput != "quit" {
		fmt.Println(MAIN_MENU)
		fmt.Print("> ")
		fmt.Scanln(&userInput)
		userInput = strings.TrimSpace(strings.ToLower(userInput))

		if userInput == "3" {
			instructions()
		}
	}
}
