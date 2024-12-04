package main

import "fmt"



func commandHelp() error {
	fmt.Println("All help is provided from the helper")
	return nil
}

func commandExit() error {
	fmt.Println("The program should exit when this is called")

	return nil
}

func main() {

	startRepl()
	// commandMap := map[string]cliCommand{
	// 	"help": {
	// 		name:        "help",
	// 		description: "Displays a help message",
	// 		callback:    commandHelp,
	// 	},
	// 	"exit": {
	// 		name:        "exit",
	// 		description: "Exit the Pokedex",
	// 		callback:    commandExit,
	// 	},
	// }

	// for {
	// 	fmt.Printf("Pokdex > ")

	// 	Scanner := bufio.NewScanner(os.Stdin)
	// 	Scanner.Scan()
	// 	err := Scanner.Err()

	// 	if err != nil {
	// 		fmt.Println("An unexpected error in buffio scanner")
	// 		return
	// 	}

	// 	command := Scanner.Text()
	// 	if command == "exit" {
	// 		return
	// 	}

	// 	if val, exists := commandMap[command]; exists {
	// 		val.callback()
	// 	} else {
	// 		fmt.Println("An unexpected error reading the command!")
	// 		return
	// 	}
	// }
}
