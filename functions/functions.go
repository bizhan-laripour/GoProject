package main

/*
*
Single return function
*/
func singleReturnFunction(param1 string, param2 string) string {
	return param1 + param2
}

/*
*
multiReturn function
*/
func multiReturnFunction() (int, int, int) {
	return 1, 2, 3
}

func sumNumbers() int {
	firstNumber, secondNumber, _ := multiReturnFunction()
	return firstNumber + secondNumber
}
