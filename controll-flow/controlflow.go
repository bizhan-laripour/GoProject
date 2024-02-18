package main

import "fmt"

func compareData(number1, number2 int) {

	if number1 > number2 {
		fmt.Println(number1, "is grater than ", number2)
	} else if number1 == number2 {
		fmt.Println("both are equals")
	} else {
		fmt.Println(number2, "is grater than ", number1)
	}
}

func or(firstNumber, secondNumber int) {

	if firstNumber > 10 || secondNumber < 5 {
		fmt.Println(firstNumber, "is grater than 10 and ", secondNumber, " is less than 5")
	} else if firstNumber < 5 || secondNumber > 10 {
		fmt.Println(firstNumber, "is less than 5 ", secondNumber, " is grater than 10")

	} else {
		fmt.Println("nothing is correct")
	}
}

func and(firstParam, secondParam string) {
	if firstParam == "Bizhan" && secondParam == "Laripour" {
		fmt.Println(firstParam, secondParam, "you are a valid user")
	} else {
		fmt.Println(firstParam, secondParam, "you are an invalid user")
	}
}

func statementInitialization(param string) {
	if i := param; i == "hello" {
		fmt.Println(i, " is equals hello")
	} else {
		fmt.Println(i, "is not equals hello")
	}

}
