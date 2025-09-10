package game

import (
	"fmt"
	"strings"
)

func instructions() {
	fmt.Println(INSTRUCTIONS)

	for userInput != "quit" && userInput != "1" {
		fmt.Print("> ")
		fmt.Scanln(&userInput)
		userInput = strings.TrimSpace(strings.ToLower(userInput))
	}
}
