package main

import (
	"GoProject/functions"
	"fmt"
)

func main() {
	// creating basic variable
	var example = "hello"
	fmt.Println("hello ", example)
	//creating variables with Type
	var name string = "bizhan"
	fmt.Println(name)

	//creating variable and assign value to it
	username := "bizhan.lp"
	fmt.Println(username)

	//get default value of a numeric variable
	var sum int
	fmt.Println(sum)
	//get default value of a String value
	var defaultOfString string
	fmt.Println(defaultOfString)

	//compound value assignment
	part1, part2 := 1, 2
	fmt.Println(part1, part2)
	sum = part1 + part2
	fmt.Println(sum)

	//multiple variables
	var (
		lastName  = "laripour"
		className = "Golang"
	)
	fmt.Println(lastName, className)
	fmt.Println(functions.SingleReturnFunction("hello", "bizhan"))

	fmt.Println(functions.SumNumbers())
}
