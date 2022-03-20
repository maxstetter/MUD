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
var Commands = make(map[string]func(string, *Player))
var Zones = make(map[int]*Zone)
var Rooms = make(map[int]*Room)
var Directions = make(map[string]int)

//var mPlayer = Player{}
var DirectionLabels = map[int]string{
	0: "n",
	1: "e",
	2: "s",
	3: "w",
	4: "u",
	5: "d",
}

//END GLOBAL//

type Zone struct {
	ID   int
	Name string
	//Rooms []*Room
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

func addCommand(command string, action func(string, *Player)) {
	for i := range command {
		if i == 0 {
			continue
		}
		prefix := command[:i]
		Commands[prefix] = action
	}
	Commands[command] = action
}

func doLook(direction string, p *Player) {
	if direction == "" {
		//fmt.Fprintf(os.Stdout, "What are you even looking at??\n")
		for dir, val := range p.Room.Exits {
			if val.Description != "" {
				fmt.Println(DirectionLabels[dir], val.Description)
			}
		}
	} else {
		fmt.Fprintf(os.Stdout, "You looked %s\n", direction)
		if p.Room.Exits[Directions[direction]].Description == "" {
			fmt.Println("There is nothing to look at in this direction")
		} else {
			fmt.Println(p.Room.Exits[Directions[direction]].Description)
		}
	}
}

func doLaugh(how string, p *Player) {
	if how == "" {
		fmt.Fprintf(os.Stdout, "hahaha\n")
	} else if how == "maniacally" {
		fmt.Fprintf(os.Stdout, "HAHAHAHA\n")
	}
}

func doSmile(s string, p *Player) {
	fmt.Printf("You smile happily.\n")
}

func doSouth(s string, p *Player) {
	if len(p.Room.Exits[2].Description) == 0 {
		fmt.Println("Illegal move")
	} else {
		p.Room = p.Room.Exits[2].To
		fmt.Printf("You move south.\n")
		fmt.Println(p.Room.Description)
	}
}

func doNorth(s string, p *Player) {
	if len(p.Room.Exits[0].Description) == 0 {
		fmt.Println("Illegal move")
	} else {
		p.Room = p.Room.Exits[0].To
		fmt.Printf("You move north.\n")
		fmt.Println(p.Room.Description)
	}
}

func doEast(s string, p *Player) {
	if len(p.Room.Exits[1].Description) == 0 {
		fmt.Println("Illegal move")
	} else {
		p.Room = p.Room.Exits[1].To
		fmt.Printf("You move east.\n")
		fmt.Println(p.Room.Description)
	}
}

func doWest(s string, p *Player) {
	if len(p.Room.Exits[3].Description) == 0 {
		fmt.Println("Illegal move")
	} else {
		p.Room = p.Room.Exits[3].To
		fmt.Printf("You move west.\n")
		fmt.Println(p.Room.Description)
	}
}
func doUp(s string, p *Player) {
	if len(p.Room.Exits[4].Description) == 0 {
		fmt.Println("Illegal move")
	} else {
		p.Room = p.Room.Exits[4].To
		fmt.Printf("You move up.\n")
		fmt.Println(p.Room.Description)
	}
}
func doDown(s string, p *Player) {
	if len(p.Room.Exits[5].Description) == 0 {
		fmt.Println("Illegal move")
	} else {
		p.Room = p.Room.Exits[5].To
		fmt.Printf("You move down.\n")
		fmt.Println(p.Room.Description)
	}
}

func doRecall(s string, p *Player) {
	p.Room = Rooms[3001]
	fmt.Printf("You return to the beginning\n")
	fmt.Printf(p.Room.Description)
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
	addCommand("recall", doRecall)
	addCommand("up", doUp)
	addCommand("down", doDown)
	//up, down, say, tell, shout, pretty call?
}

func doCommand(command string, player *Player) error {
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
		function(target, player)
	} else {
		fmt.Printf("You said wut?\n")
	}
	return nil
}

func commandLoop() error {
	mPlayer := Player{Room: Rooms[3001]}
	//mPlayer.Room = Rooms[3001]
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Print("--> ")
		line := scanner.Text()
		err := doCommand(line, &mPlayer)
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
func readSingleRoom(db *sql.DB) error {
	//select id, zone_id, name, description from rooms where id = 3001;
	rows, err := db.Query("SELECT id, name, description FROM rooms where ID = 3001")
	if err != nil {
		return fmt.Errorf("querying a room from the database: %v", err)
	}

	var room = Room{}
	for rows.Next() {
		var id int
		var name, description string
		if err := rows.Scan(&id, &name, &description); err != nil {
			return fmt.Errorf("reading a room from the database: %v", err)
		}
		//var room = Room{id, Zones[zone_id], name, description, exits}
		room = Room{ID: id, Name: name, Description: description}
	}

	//TODO: get rid of redundant room, implement recall
	fmt.Println("ID: ", room.ID)
	fmt.Println("Name: ", Rooms[3001].Name)
	fmt.Println("Description: ", Rooms[3001].Description)
	//fmt.Println("Zone is: ", Rooms[3001].Zone)
	//fmt.Println("Exits: ", Rooms[3001].Exits)
	for dir, val := range Rooms[3001].Exits {
		if val.Description != "" {
			fmt.Println(DirectionLabels[dir], val.Description)
		}
	}

	return nil
}

//readZones() function reads all of the zones. Collects all of the zones into a map where the keys are zone IDs and the values are Zone pointers. Prints them all out.
func readZones(stmt *sql.Stmt) (map[int]*Zone, error) {
	//rows, err := db.Query("SELECT * FROM zones")
	rows, err := stmt.Query()
	if err != nil {
		return nil, fmt.Errorf("querying zones from database: %v", err)
	}

	for rows.Next() {
		var id int
		var name string
		//var rooms []*Room
		if err := rows.Scan(&id, &name); err != nil {
			return nil, fmt.Errorf("reading zones from database: %v", err)
		}
		zone := Zone{id, name} //, rooms}
		//fmt.Println(zone)
		Zones[id] = &zone
	}
	//for key, value := range Zones {
	//	fmt.Println("zoneID:", key, " ", *value)
	//}
	return Zones, nil
}

//The readRooms function reads in all of the rooms. It accepts an open transaction as a paramter and returns a map from IDs to Room pointers. In addition, have it accept the map of zones as a parameter. When you get a zone ID from the database, use it to find the corresponding Zone pointer and store it in the Room object.
func readRooms(stmt *sql.Stmt, ZoneMap map[int]*Zone) (map[int]*Room, error) {
	rows, err := stmt.Query()
	if err != nil {
		return nil, fmt.Errorf("querying rooms from database: %v", err)
	}
	for rows.Next() {
		var room_id, zone_id int
		var name, description string
		var exits [6]Exit
		if err := rows.Scan(&room_id, &zone_id, &name, &description); err != nil {
			return nil, fmt.Errorf("reading rooms from database: %v", err)
		}
		zonePointer := ZoneMap[zone_id]
		room := Room{room_id, zonePointer, name, description, exits}
		//fmt.Println(room)
		Rooms[room_id] = &room
		//ZoneMap[zone_id].Rooms = []*Room{&room}
	}
	return Rooms, nil
}

//The readExits() function reads in all of the exits, finds the room it leaves from and fills in the corresponding exit field of the room.``
func readExits(stmt *sql.Stmt) (map[int]*Room, error) {
	rows, err := stmt.Query()
	if err != nil {
		return nil, fmt.Errorf("querying exits from database: %v", err)
	}
	for rows.Next() {
		var fromRoomId, toRoomId int
		var direction, description string
		if err := rows.Scan(&fromRoomId, &toRoomId, &direction, &description); err != nil {
			return nil, fmt.Errorf("reading exits from database: %v", err)
		}
		exit := Exit{Rooms[toRoomId], description}
		//fmt.Println("the exit is: ", exit)
		//if exit.To != nil {
		//	//why are there doors with no description? Do we still tell the user about them?
		//	fmt.Println("there is a door leading to: ", exit.To.Name)
		//}
		Rooms[fromRoomId].Exits[Directions[direction]] = exit
	}
	return Rooms, nil
}

func printRooms() {
	for key, _ := range Rooms {
		fmt.Println("the key is: ", key)
		fmt.Println(Rooms[key].Name)
		fmt.Println(Rooms[key].Description)
		fmt.Println("Zone is: ", Rooms[key].Zone)
		fmt.Println("Exits: ", Rooms[key].Exits)
		for dir, val := range Rooms[key].Exits {
			fmt.Println("dir: ", dir, " val: ", val)
		}
	}
}

func main() {
	Directions["n"] = 0
	Directions["e"] = 1
	Directions["s"] = 2
	Directions["w"] = 3
	Directions["u"] = 4
	Directions["d"] = 5

	// use time and origin file for log prefixes
	log.SetFlags(log.Ltime | log.Lshortfile)
	initialize()
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
		log.Fatalf("begin zones read transaction: %v", err)
	}
	stmt, err := tx.Prepare(`SELECT * FROM zones`)
	if err != nil {
		log.Fatalf("prepare room read transaction: %v", err)
	}
	defer stmt.Close()

	zoneMap, e := readZones(stmt)
	if e != nil {
		log.Fatalf("readZones: %v", e)
		tx.Rollback()
	} else {
		tx.Commit()
	}

	tx, err = db.Begin()
	if err != nil {
		log.Fatalf("begin room read transaction: %v", err)
	}
	stmt, err = tx.Prepare(`SELECT * FROM rooms`)
	if err != nil {
		log.Fatalf("prepare room read transaction: %v", err)
	}
	defer stmt.Close()

	if _, e := readRooms(stmt, zoneMap); e != nil {
		log.Fatalf("readRooms: %v", e)
		tx.Rollback()
	} else {
		tx.Commit()
	}

	tx, err = db.Begin()
	if err != nil {
		log.Fatalf("begin exit read transaction: %v", err)
	}
	stmt, err = tx.Prepare(`SELECT * FROM exits`)
	if err != nil {
		log.Fatalf("prepare exit read transaction: %v", err)
	}
	defer stmt.Close()

	if _, err := readExits(stmt); err != nil {
		log.Fatalf("readExits: %v", err)
		tx.Rollback()
	} else {
		tx.Commit()
	}

	//printRooms()
	if e := readSingleRoom(db); e != nil {
		log.Fatalf("readSingleRoom: %v", e)
	}

	fmt.Println("WELCOME TO THE DUNGEON")
	fmt.Println("Enter: ")
	if err := commandLoop(); err != nil {
		log.Fatalf("%v", err)
	}
}
