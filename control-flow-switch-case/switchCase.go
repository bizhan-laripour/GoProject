package main

import "fmt"

const (
	Economy    = 0
	Business   = 1
	FirstClass = 2
)

func main() {
	getNumber(1)
	conditionalSwitch(3)
	caseList(5)
	fallthroughCase(1)
	yourCategory(Economy)
	categoryByConst()
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
func caseList(number int) {
	switch number {
	case 1, 2, 3:
		fmt.Println("the number is one of 1 , 2 , 3")
	case 4, 5, 6:
		fmt.Println("the number is one of 4,5,6")
	default:
		fmt.Println("none of them")
	}
}

func fallthroughCase(number int) {
	switch number {
	case 1:
		fmt.Println("this is one")
		fallthrough
	case 2:
		fmt.Println("this will execute by fallthrough")
		fallthrough
	case 3:
		fmt.Println("this will execute too")
		fallthrough
	default:
		fmt.Println("this is default block that will execute too")

	}
}

func price(category int) int {
	return category
}

func yourCategory(category int) {
	switch price := price(category); {
	case price == 0:
		fmt.Println("this is economy class and is cheap")
	case price == 1:
		fmt.Println("this is business class and is ok")
	case price == 2:
		fmt.Println("this is first class and is expensive")
	default:
		fmt.Println("choose a category")
	}
}

func categoryByConst() {
	ticket := Economy
	switch ticket {
	case Economy:
		fmt.Println("this is ec")
	case Business:
		fmt.Println("business")
	case FirstClass:
		fmt.Println("first class")
	default:
		fmt.Println("other")
	}

}
