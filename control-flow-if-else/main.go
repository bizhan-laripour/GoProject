package main

import "fmt"

func main() {

	compareData(10, 11)
	or(11, 4)
	and("Bizhan", "Laripour")
	statementInitialization("hello")
	student1 := Student{name: "Bizhan", lastName: "Laripour", age: 23}
	student2 := Student{name: "Ahmad", lastName: "Jafari", age: 32}
	studentSlice1 := []Student{student1}
	students := append(studentSlice1, student2)
	class1 := Class{name: "A1", numberOfStudents: students}
	student3 := Student{name: "hossein", lastName: "hosseinZadeh", age: 65}
	student4 := Student{name: "amir", lastName: "alipour", age: 45}
	studentSlice2 := []Student{student3}
	students2 := append(studentSlice2, student4)
	class2 := Class{name: "B1", numberOfStudents: students2}
	compare(class1, class2)

}

type Class struct {
	name             string
	numberOfStudents []Student
}

type Student struct {
	name     string
	lastName string
	age      int
}

func checkClassAge(class Class) int {
	students := class.numberOfStudents
	average := 0
	for i := 0; i < len(students); i++ {
		average += students[i].age
	}
	return average
}

func compare(class1 Class, class2 Class) {
	var first int = checkClassAge(class1)
	var second int = checkClassAge(class2)
	if first > second {
		fmt.Println("first class is older than second class:", class1.name, "is older than", class2.name)
	} else if second > first {
		fmt.Println("second class is older than first class", class2.name, "is older than", class1.name)
	} else {
		fmt.Println("both of them are equal", class1.name, "is equals to ", class2.name)
	}
}
