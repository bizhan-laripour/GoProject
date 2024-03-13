//this library developed by maps
package main

import "fmt"

func main() {
    user1 := User2{
       name:     "bizhan",
       lastName: "Laripour",
       id:       1,
    }

    book1 := Book2{
       name:     "java language",
       ISBN:     "123456789",
       isBooked: false,
    }

    registerUser(user1)
    addBook(book1)
    isBookedAdd(user1, book1)
    fmt.Println(library2)
}

type Library2 struct {
    name       string
    bookingMap map[int][]Book2
    users      []User2
    books      []Book2
}

type User2 struct {
    id       int
    name     string
    lastName string
}

type Book2 struct {
    name     string
    ISBN     string
    isBooked bool
}

var library2 Library2

func registerUser(user User2) {
    library2.users = append(library2.users, user)
}

func addBook(book Book2) {
    library2.books = append(library2.books, book)
}

func isBookedAdd(user User2, book Book2) {
    bookingMap := library2.bookingMap
    if bookingMap != nil {
       for _, value := range bookingMap {
          for _, BookValue := range value {
             if BookValue.ISBN == book.ISBN {
                fmt.Println("this book is booked")
                return
             }
          }
       }
    } else {
       bookingMap = make(map[int][]Book2)
       library2.bookingMap = bookingMap
    }
    book.isBooked = true
    value, isExist := library2.bookingMap[user.id]
    if !isExist {
       bookSlice := []Book2{}

       bookSlice = append(bookSlice, book)
       library2.bookingMap[user.id] = bookSlice
    } else {
       value = append(value, book)
       library2.bookingMap[user.id] = value
    }
}
