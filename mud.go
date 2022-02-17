package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

////GLOBAL////
var Commands = make(map[string]func(string))

//END GLOBAL//

func AddCommand(command string, action func(string)) {
	Commands[command] = action
}

func doLook(direction string) {
	if direction == "" {
		fmt.Fprintf(os.Stdout, "what are you even looking at??\n")
	} else {
		fmt.Fprintf(os.Stdout, "looked %s\n", direction)
	}
}

func doLaugh(how string) {
	if how == "" {
		fmt.Fprintf(os.Stdout, "hahaha\n")
	} else if how == "maniacally" {
		fmt.Fprintf(os.Stdout, "HAHAHAHA\n")
	}
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
	AddCommand("look", doLook)
	AddCommand("laugh", doLaugh)
}

func doCommand(command string) error {
	input := strings.Fields(command)
	target := ""
	//var target []string
	if len(input) == 0 {
		return errors.New("empty input, try again")
	} else if len(input) >= 2 {
		command = input[0]
		//target = input[1:]
		for i := 1; i < len(input); i++ {
			if i == len(input)-1 {
				target += input[i]
			} else {
				target += input[i] + " "
			}
		}
	}

	if function, ok := Commands[command]; ok {
		function(target)
	} else {
		return errors.New("that command was not found, try another")
	}
	return nil
}

func main() {
	initialize()

	fmt.Println("WELCOME TO THE DUNGEON")
	fmt.Println("Enter: ") //Ask for input here?

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		err := doCommand(line)
		if err != nil {
			fmt.Printf("ERROR: %v \n", err)
			err = nil
		}
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
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Reading standard input:", err)
	}
}
