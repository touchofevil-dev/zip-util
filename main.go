package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 4 {
		fmt.Println("Usage: main <ZIP | Unzip> <source> <destination>")
		os.Exit(1)
	}
	command := os.Args[1]
	source := os.Args[2]
	destination := os.Args[3]

	if strings.ToLower(command) == "zip" {
		zip(source, destination)
	} else if strings.ToLower(command) == "unzip" {
		unzip(source, destination)
	} else {
		fmt.Println("Invalid command")
		os.Exit(1)
	}
}
