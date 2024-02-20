package main

import "fmt"

func main() {
	createMySlice()
	anotherSlice()
	dynamicArray()
	dynamicArray2()
	preAllocateDynamicArray()

	mySlice := []int{1, 2, 3, 4}
	secondSlice := []string{"hello", "world"}
	iterateOverSlices(mySlice)
	iterateOverStringSlices(secondSlice)

}

func createMySlice() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}
}

func anotherSlice() {
	array := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	slice1 := array[:]
	slice2 := array[1:5]
	for i := 0; i < len(slice1); i++ {
		fmt.Println(slice1[i])
	}
	for i := 0; i < len(slice2); i++ {
		fmt.Println(slice2[i])
	}
}

func dynamicArray() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	numbers = append(numbers, 10, 11, 12, 13, 14, 15)
	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	}
}

func dynamicArray2() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	numbers2 := []int{10, 11, 12, 13, 14, 15}
	combined := append(numbers, numbers2...)

	for i := 0; i < len(combined); i++ {
		fmt.Println(combined[i])
	}

}

func preAllocateDynamicArray() {
	slice := make([]int, 10)
	fmt.Println("the pre allocate slice has size of", len(slice))
	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}
}

func iterateOverSlices(slice []int) {
	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}
}

func iterateOverStringSlices(slice []string) {
	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}
}
