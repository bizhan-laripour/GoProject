package main

import "fmt"

func main() {
	getNumber(1)
	conditionalSwitch(3)
}

func getNumber(number int) {
	switch number {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	default:
		fmt.Println("non of above numbers")

	}
}

func simpleSwitch(number int) string {

	switch number {
	case 1:
		return "one"
	case 2:
		return "two"
	case 3:
		return "three"
	default:
		return ""
	}
}

func conditionalSwitch(number int) {
	switch result := simpleSwitch(number); {
	case result == "one":
		fmt.Println("this is one")
	case result == "two":
		fmt.Println("this is two")
	case result == "three":
		fmt.Println("this is three")
	default:
		fmt.Println("this is none of them")

	}

}
