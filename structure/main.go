package main

import "fmt"

func main() {
	// use of a simple struct
	data := Sample{
		name:     "bizhan",
		lastName: "laripour",
		age:      32,
	}
	data.age = 25
	fmt.Println(data.age)
	fmt.Println(data.name)

	//use of anonymous and inline anonymous struct
	anonymous()
	inlineAnonymous()

}
