package main

import "fmt"

type Passenger struct {
	name     string
	lastName string
	age      int
}

type Buss struct {
	passenger Passenger
	bussType  string
}

func main() {
	bizhan := Passenger{
		name:     "bizhan",
		lastName: "Laripour",
		age:      32,
	}
	bus := Buss{
		passenger: bizhan,
		bussType:  "volvo",
	}

	fmt.Println(bus)
}
