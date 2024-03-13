
//this is some struct samples
package main

import "fmt"

/*
*
This is a struct like a class in java
*/
type Sample struct {
	name     string
	lastName string
	age      int
}

/*
*
this is anonymous type that used in a function
*/
func anonymous() {

	var anonymousType struct {
		name     string
		lastName string
		age      int
	}
	anonymousType.age = 10
	anonymousType.name = "anonymous"
	anonymousType.lastName = "lastName"
	fmt.Println(anonymousType.lastName)
	fmt.Println(anonymousType.name)
	fmt.Println(anonymousType.age)
}

/*
*
this is inline anonymous type that used in a function
*/
func inlineAnonymous() {
	inline := struct {
		name     string
		lastName string
		age      int
	}{
		"Inline anonymous name",
		"Inline anonymous lastName",
		5,
	}
	fmt.Println(inline.age)
	fmt.Println(inline.name)
	fmt.Println(inline.lastName)
}
