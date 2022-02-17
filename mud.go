package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)

////GLOBAL////
var Commands = make(map[string]func(string))

//END GLOBAL//

func addCommand(command string, action func(string)) {
	for i := range command {
		if i == 0 {
			continue
		}
		prefix := command[:i]
		Commands[prefix] = action
	}
	Commands[command] = action
}

func doLook(direction string) {
	if direction == "" {
		fmt.Fprintf(os.Stdout, "What are you even looking at??\n")
	} else {
		fmt.Fprintf(os.Stdout, "You looked %s\n", direction)
	}
}

func doLaugh(how string) {
	if how == "" {
		fmt.Fprintf(os.Stdout, "hahaha\n")
	} else if how == "maniacally" {
		fmt.Fprintf(os.Stdout, "HAHAHAHA\n")
	}
}

func doSmile(s string) {
	fmt.Printf("You smile happily.\n")
}

func doSouth(s string) {
	fmt.Printf("You move south.\n")
}

func doNorth(s string) {
	fmt.Printf("You move north.\n")
}

func doEast(s string) {
	fmt.Printf("You move east.\n")
}

func doWest(s string) {
	fmt.Printf("You move west.\n")
}

//not implemented yet
func DetermineCommand(input string) string {
	var command string
	switch {
	case input == "":
		command = ""
		return command
	case input == "":
		command = ""
		return command
	case input == "":
		command = ""
		return command
	case input == "":
		command = ""
		return command
	case input == "":
		command = ""
		return command
	case input == "":
		command = ""
		return command
	}
	return "asdf"
}

//initialize the commands
func initialize() {
	addCommand("look", doLook)
	addCommand("laugh", doLaugh)
	addCommand("smile", doSmile)
	addCommand("south", doSouth)
	addCommand("north", doNorth)
	addCommand("east", doEast)
	addCommand("west", doWest)
}

func doCommand(command string) error {
	input := strings.Fields(command)
	target := ""
	if len(input) == 0 {
		return errors.New("empty input, try again")
	} else if len(input) >= 2 {
		command = input[0]
		for i := 1; i < len(input); i++ {
			if i == len(input)-1 {
				target += input[i]
			} else {
				target += input[i] + " "
			}
		}
	}

	if function, exists := Commands[strings.ToLower(command)]; exists {
		function(target)
	} else {
		fmt.Printf("You said wut?\n")
	}
	return nil
}

func commandLoop() error {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Print("--> ")
		line := scanner.Text()
		err := doCommand(line)
		if err != nil {
			fmt.Printf("ERROR: %v \n", err)
			err = nil
		}
	}
	if err := scanner.Err(); err != nil {
		//fmt.Fprintln(os.Stderr, "Reading standard input:", err)
		return fmt.Errorf("in main command loop: %v", err)
	}
	return nil
}

func main() {
	fmt.Println("WELCOME TO THE DUNGEON")
	fmt.Println("Enter: ") //Ask for input here?
	// use time and origin file for log prefixes
	log.SetFlags(log.Ltime | log.Lshortfile)
	initialize()
	if err := commandLoop(); err != nil {
		log.Fatalf("%v", err)
	}

	//	var command string
	//	var target string
	//	//split input into individual words by white space.
	//	input := strings.Fields(scanner.Text())
	//	if len(input) == 0 {
	//		fmt.Fprint(os.Stdout, "Empty input. Try again.\n")
	//		fmt.Println("Enter: ")
	//	} else if len(input) >= 2 {
	//		command = input[0]
	//		target = input[1]
	//		fmt.Fprint(os.Stdout, "Command = ", command, "\n")
	//		fmt.Fprint(os.Stdout, "Target = ", target, "\n")
	//	} else {
	//		command = input[0]
	//		fmt.Fprint(os.Stdout, "Command = ", command, "\n")
	//	}
	//	//TODO: handle say command and implement command shortcuts.
	//	if function, ok := Commands[command]; ok {
	//		function(target)
	//	}
	//	//prints input.
	//	fmt.Fprint(os.Stdout, "--> ", scanner.Text(), "\n") //remove when finished.
	//
	// 	}

}
