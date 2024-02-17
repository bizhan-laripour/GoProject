package main

import "fmt"

func getData(number1, number2 int) {

	if number1 > number2 {
		fmt.Println(number1, "is grater than ", number2)
	} else if number1 == number2 {
		fmt.Println("both are equals")
	} else {
		fmt.Println(number2, "is grater than ", number1)
	}
}
