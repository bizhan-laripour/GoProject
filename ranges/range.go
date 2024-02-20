package main

import "fmt"

/*
*
this is like for each in java
*/
func main() {
	makeRange()
	anotherRange()

}

func makeRange() {
	slice := []string{"hello", "world", "bizhan"}
	slice2 := append(slice, "bizhan", "laripour")
	for j, element := range slice {
		fmt.Println(j, element)
	}

	for i, el := range slice2 {
		fmt.Println(i, el)
	}
}

func anotherRange() {
	slice := []string{"hello", "bizhan", "laripour"}
	for _, element := range slice {
		fmt.Println(element)
	}
}
