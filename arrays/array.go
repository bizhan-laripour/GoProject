package main

import "fmt"

func main() {
	fmt.Println(array())
	forOnArray()
	runtimeError()
	compileTimeError1()
	compileTimeError2()
}

func array() [3]int {
	var myArray [3]int
	myArray[0] = 7
	myArray[1] = 8
	myArray[2] = 10
	return myArray
}

func forOnArray() {
	myArray := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i := 0; i < len(myArray); i++ {
		fmt.Println(myArray[i])
	}
}

func runtimeError() {
	var myArray [3]int
	for i := 0; i < 10; i++ {
		fmt.Println(myArray[i])
	}
}

func compileTimeError1() {
	var myArray [3]int
	myArray[0] = 1
	myArray[1] = 2
	myArray[2] = 3
	myArray[3] = 4
}

func compileTimeError2() {
	var myArray [3]int
	item := myArray[4]
	fmt.Println(item)
}
