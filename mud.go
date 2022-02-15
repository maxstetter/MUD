package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
)

////GLOBAL////
var Commands = make(map[string]func(string))
//END GLOBAL//

func AddCommand(command string, action func(string)){
	Commands[command] = action
}

func doLook(direction string) {
	fmt.Fprintf(os.Stdout, "looked %s\n", direction)
}

func doLaugh(laughylaugh string) {
	fmt.Fprintf(os.Stdout, "HAHAHAHA\n")
}

//not implemented yet
func DetermineCommand(input string) string{
	var command string
	switch{
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
func initialize(){
	AddCommand("look", doLook)
	AddCommand("laugh", doLaugh)
}

func main(){
	initialize()

	fmt.Println("WELCOME TO THE DUNGEON\n")
	fmt.Println("Enter: ") //Ask for input here?

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		var command string
		var target string
		//split input into individual words by white space.
		input := strings.Fields(scanner.Text())
		if len(input) == 0 {
			fmt.Fprint(os.Stdout, "Empty input. Try again.\n")
			fmt.Println("Enter: ")
		}else if len(input) >= 2{
			command = input[0]
			target = input[1]
			fmt.Fprint(os.Stdout,"Command = " ,command, "\n")
			fmt.Fprint(os.Stdout,"Target = " ,target, "\n")
		}else{
			command = input[0]
			fmt.Fprint(os.Stdout,"Command = " ,command, "\n")
		}
		//TODO: handle say command and implement command shortcuts.
		if function, ok := Commands[command]; ok{
			function(target)
		}
		//prints input.
		fmt.Fprint(os.Stdout, "--> ", scanner.Text(), "\n") //remove when finished.

	}
	if err := scanner.Err(); err != nil{
		fmt.Fprintln(os.Stderr, "Reading standard input:", err)
	}
}
