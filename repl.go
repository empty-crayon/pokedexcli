package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Printf("Pokidex>")
		scanner.Scan()
		if err := scanner.Err(); err != nil {
			fmt.Printf("An unexpected error happened in reading from buffio: %err", err)
			return
		}

		text := scanner.Text()
		cleaned := cleanInput(text)
		if len(cleaned) == 0 {
			continue
		}
		commandName := cleaned[0]

		availableCommands := getCommands()

		command, ok := availableCommands[commandName]

		if !ok {
			fmt.Println("Woops invalid command!")
			continue
		}

		command.callback()

	}
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Prints the help menu of the CLI",
			callback:    callbackHelp,
		},
		"exit": {
			name:        "exit",
			description: "Terminate the CLI",
			callback:    callbackExit,
		},
		"map": {
			name:        "map",
			description: "Displays the names of 20 location areas in the Pokemon world",
			callback:    callbackMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Instead of displaying the next 20 locations, it displays the previous 20 locations. It's a way to go back.",
			callback:    callbackMapb,
		},
	}
}

func cleanInput(str string) []string {
	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)

	return words

}
