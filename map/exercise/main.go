package main

import "fmt"

var shopingList = make(map[string]int)

func main() {
	generateShoppingList()
	deleteFromShoppingList("eggs")
	fmt.Println(shopingList)
	addToShoppingList("meat", 2)
	fmt.Println(shopingList)
	deleteAnItemFromShoppingList("chicken")
	fmt.Println(shopingList)
	iterateOnMap(shopingList)
}

func generateShoppingList() {
	shopingList["eggs"] = 200
	shopingList["meat"] = 12
	shopingList["chicken"] = 56
	shopingList["milk"] = 6
	shopingList["bread"] = 1
	shopingList["bread"] += 1
}

func deleteFromShoppingList(key string) {
	currentNumber := shopingList[key]
	if currentNumber != 0 {
		shopingList[key] = currentNumber - 1
	}
}

func addToShoppingList(key string, numbers int) {
	currentNumbers := shopingList[key]
	shopingList[key] = currentNumbers + numbers
}

func deleteAnItemFromShoppingList(key string) {
	delete(shopingList, key)
	value, isExist := shopingList[key]
	if isExist {
		fmt.Println(value, "not deleted from shoppingList")
	}
}

func iterateOnMap(myMap map[string]int) {
	for key, value := range myMap {
		fmt.Println(key, value)
	}
}
