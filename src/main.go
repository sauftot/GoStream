package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var (
	ingestServer string
	streamKey    string
	directory    string
	currentFile  string
	streaming    bool
)

func streamController(consoleCommands chan string, response chan string) {

}

func consoleReader(consoleCommands chan string) {
	var buf string
	for {
		fmt.Scanln(&buf)
		consoleCommands <- buf
	}
}

// wait for user input from console, then send input to main and streamController
func consoleController(Input chan string, Output chan string, StrIn chan string) {
	consoleInput := make(chan string)
	go consoleReader(consoleInput)
	for {
		select {
		case input := <-consoleInput:
			iParts := strings.Split(input, " ")
			if len(iParts) == 1 {
				if strings.Contains(strings.ToLower(input), "exit") {
					fmt.Println("Exiting...")
					os.Exit(0)
				} else if strings.Contains(strings.ToLower(input), "help") {
					fmt.Println("Commands: ")
				} else {
					fmt.Println("Unknown command or wrong number of parameters. Type 'help' for a list of commands.")
				}
			} else if strings.Contains(strings.ToLower(input), "key") {
				streamKey = strings.Split(input, " ")[1]
			} else if strings.Contains(strings.ToLower(input), "server") {

			} else if strings.Contains(strings.ToLower(input), "file") {

			} else if strings.Contains(strings.ToLower(input), "cd") {

			} else if strings.Contains(strings.ToLower(input), "start") {

			} else if strings.Contains(strings.ToLower(input), "stop") {

			} else {
				fmt.Println("Unknown command. Type 'help' for a list of commands.")
			}
		}
	}
}

func main() {
	exePath, err := os.Executable()
	if err != nil {
		fmt.Println("Error getting executable path:", err)
		return
	}

	// Get the directory part from the absolute path
	directory = filepath.Dir(exePath)

	cIn, cOut, cToStr := make(chan string), make(chan string), make(chan string)
	go consoleController(cIn, cOut, cToStr)
	go streamController(cIn, cToStr)
	for {
		select {
		case <-cIn:

		default:
		}
	}
}
