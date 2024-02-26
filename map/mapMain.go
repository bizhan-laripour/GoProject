package main

import "fmt"

func main() {
	myMap := generateFirstMap()
	fmt.Println(myMap["first"])
	secondMap := generateMapWithMakeFunction()
	fmt.Println(secondMap["second"])

}

func generateFirstMap() map[string]string {
	myMap := map[string]string{
		"first":  "one",
		"second": "two",
	}
	return myMap
}

func generateMapWithMakeFunction() map[string]string {
	myMap := make(map[string]string)
	myMap["first"] = "one"
	myMap["second"] = "two"
	return myMap
}
