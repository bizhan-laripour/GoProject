package main

import "fmt"

func main() {
	mobile := Product{
		name:  "mobile",
		price: 3000000,
	}
	book := Product{
		name:  "book",
		price: 250000,
	}

	addProductToTheList(mobile, 0)
	addProductToTheList(book, 1)
	fmt.Println(getProduct(0))
	fmt.Println(getTotalPrice())
	fmt.Println(getTotalItems())
	totalPrice, totalItems := getTotalPriceAndItemsInSingleFunction()
	fmt.Println("totalItems is", totalItems)
	fmt.Println("getTotalPrice is", totalPrice)
}

var productList [4]Product

type Product struct {
	name  string
	price int
}

func addProductToTheList(product Product, index int) {
	productList[index] = product
}

func getProduct(index int) Product {
	return productList[index]
}

func getTotalPrice() int {
	price := 0
	for i := 0; i < len(productList); i++ {
		price += productList[i].price
	}
	return price
}

func getTotalItems() int {
	totalItems := 0
	for i := 0; i < len(productList); i++ {
		item := productList[i]
		if item.name != "" {
			totalItems++
		}
	}
	return totalItems
}

// this kind of using functions is fantastic
func getTotalPriceAndItemsInSingleFunction() (int, int) {
	return getTotalPrice(), getTotalItems()
}
