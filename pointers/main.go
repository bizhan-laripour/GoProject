package main

import "fmt"

/*
*
by default using inputs are by value that means golang uses a copy of an input no matter theirs sizes
it costs a lot but you can use call by reference by pointers.
pointers help us to access the memory address of a value and fetch it without coping the input value
*/
func main() {
	i := 6
	generate(&i)
	fmt.Println(i)
}

func generate(pointer *int) {
	*pointer += 1
}
