package main

import (
	"fmt"
	m "vamaracinggp/internal/motorcycles"
)

func main() {
	for _, v := range m.Menu {
		fmt.Println(v)
		fmt.Println("-------------------------")
	}
}
