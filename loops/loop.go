package main

import "fmt"

func main() {
	cStyleLoop()
	whileLoop()
	breakKeyWord()
	continueKeyWord()
	multipleTypeLoop()
}

func cStyleLoop() {
	for i := 0; i < 10; i++ {
		fmt.Println("this is ", i)
	}
}

func whileLoop() {
	i := 0
	for i < 10 {
		fmt.Println("this is ", i)
		i++
	}
}

func breakKeyWord() {
	for i := 0; i < 10; i++ {
		fmt.Println("this is ", i)
		if i == 5 {
			fmt.Println("break on", i)
			break
		}
	}
}

func continueKeyWord() {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println("the number is odd", i)

	}
}

func multipleTypeLoop() {
	sum := 0
	// this is basic loop section
	for i := 0; i <= 10; i++ {
		sum += i
		fmt.Println("sum is", sum)
	}

	// this is while loop section
	for sum > 10 {
		sum -= 5
		fmt.Println("decrement sum is", sum)
	}
}
