package main

import "fmt"

func main() {
	//you can use this kind of populating the array
	room1 := Room{name: "room1", cleaned: false}
	room2 := Room{name: "room2", cleaned: true}
	room3 := Room{name: "room3", cleaned: true}
	room4 := Room{name: "room4", cleaned: false}
	rooms := [...]Room{room1, room2, room3, room4}
	checkCleanliness(rooms)
	// or you can use this kind of populating the array
	secondRoomsArray := [...]Room{
		{name: "room1", cleaned: false},
		{name: "room2", cleaned: true},
		{name: "room3", cleaned: true},
		{name: "room4", cleaned: false},
	}

	checkCleanliness(secondRoomsArray)

}

type Room struct {
	name    string
	cleaned bool
}

func checkCleanliness(rooms [4]Room) {
	for i := 0; i < len(rooms); i++ {
		room := rooms[i]
		if room.cleaned == true {
			fmt.Println(room.name, "is cleaned")
		}
	}
}
