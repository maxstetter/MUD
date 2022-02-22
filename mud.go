package main

import (
	"bufio"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

////GLOBAL////
var Commands = make(map[string]func(string))
var Zones = make(map[int]*Zone)
var Rooms = make(map[int]*Room)
var Directions = make(map[string]int)
var player = Player{}
var db *sql.DB

//END GLOBAL//

type Zone struct {
	ID    int
	Name  string
	Rooms []*Room
}

type Room struct {
	ID          int
	Zone        *Zone
	Name        string
	Description string
	Exits       [6]Exit
}

type Exit struct {
	To          *Room
	Description string
}

type Player struct {
	Room *Room
}

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
	//up, down, say, tell, shout, pretty call?
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

//This function opens the database, reads a single room and stores the ID, Name and Descriptions fields in a Room object, prints this object out
func readRooms(db *sql.DB) error {
	//select id, zone_id, name, description from rooms where id = 3001;
	rows, err := db.Query("SELECT id, name, description FROM rooms where ID = 3001")
	if err != nil {
		return fmt.Errorf("querying a room from the database: %v", err)
	}

	for rows.Next() {
		var id int
		var name, description string
		if err := rows.Scan(&id, &name, &description); err != nil {
			return fmt.Errorf("reading a room from the database: %v", err)
		}
		//var room = Room{id, Zones[zone_id], name, description, exits}
		room := Room{ID: id, Name: name, Description: description}
		fmt.Println(room)
	}

	return nil
}

//readZones() function reads all of the zones. Collects all of the zones into a map where the keys are zone IDs and the values are Zone pointers. Prints them all out.
func readZones(stmt *sql.Stmt) error {
	//rows, err := db.Query("SELECT * FROM zones")
	rows, err := stmt.Query()
	if err != nil {
		return fmt.Errorf("querying zones from database: %v", err)
	}

	for rows.Next() {
		var id int
		var name string
		var rooms []*Room
		if err := rows.Scan(&id, &name); err != nil {
			return fmt.Errorf("reading zones from database: %v", err)
		}
		zone := Zone{id, name, rooms}
		fmt.Println(zone)
		Zones[id] = &zone
	}
	fmt.Println(Zones)
	return nil
}

func main() {
	//	Directions["n"] = 0
	//	Directions["e"] = 1
	//	Directions["s"] = 2
	//	Directions["w"] = 3
	//	Directions["u"] = 4
	//	Directions["d"] = 5

	// the path to the database--this could be an absolute path
	path := "world.db"
	options := "?" + "_busy_timeout=10000" +
		"&" + "_foreign_keys=ON"
	db, err := sql.Open("sqlite3", path+options)
	if err != nil {
		log.Fatalf("opening database: %v", err)
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		log.Fatalf("begin room read transaction: %v", err)
	}
	stmt, err := tx.Prepare(`SELECT * FROM rooms`)
	if err != nil {
		log.Fatalf("prepare room read transaction: %v", err)
	}
	defer stmt.Close()

	if e := readRooms(db); e != nil {
		log.Fatalf("readRooms: %v", e)
		tx.Rollback()
	} else {
		tx.Commit()
	}

	tx, err = db.Begin()
	if err != nil {
		log.Fatalf("begin zones read transaction: %v", err)
	}
	stmt, err = tx.Prepare(`SELECT * FROM zones`)
	if err != nil {
		log.Fatalf("prepare room read transaction: %v", err)
	}
	defer stmt.Close()

	if e := readZones(stmt); e != nil {
		log.Fatalf("readZones: %v", e)
		tx.Rollback()
	} else {
		tx.Commit()
	}
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
