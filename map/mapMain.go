package main

import "fmt"

func main() {
	myMap := generateFirstMap()
	fmt.Println(myMap["first"])
	secondMap := generateMapWithMakeFunction()
	fmt.Println(secondMap["second"])
	deleteFromMap()

	//	iterate on map
	myMap1 := map[string]string{
		"hello": "hello1",
		"world": "world1",
	}
	iterateOnMap(myMap1)
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

func deleteFromMap() {
	myMap := map[string]string{
		"hello": "hello1",
		"world": "world1",
	}
	delete(myMap, "hello")
	value, isExist := myMap["hello"]
	if !isExist {
		fmt.Println("not found")

	} else {
		fmt.Println(value)
	}
}

func iterateOnMap(myMap map[string]string) {
	for key, value := range myMap {
		fmt.Println(key, value)
	}
}
