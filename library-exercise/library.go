//this is a library that a user can book a book
package main

import "fmt"

type Book struct {
    name     string
    ISBN     string
    isBooked bool
    bookedBy int
}

type User struct {
    id       int
    name     string
    lastName string
    book     []Book
}

type Library struct {
    name      string
    bookSlice []Book
    users     []User
}

var library Library
var books []Book

var users []User

func main() {

    // create user
    user := User{
       name:     "bizhan",
       lastName: "laripour",
       id:       1,
    }
    user2 := User{
       name:     "ali",
       lastName: "ahmadi",
       id:       2,
    }
    //register user
    registerUserToLibrary(user)
    registerUserToLibrary(user2)

    // create book
    book := Book{
       name:     "golang",
       ISBN:     "123654789",
       isBooked: false,
    }
    book2 := Book{
       name:     "java",
       ISBN:     "987654321",
       isBooked: false,
    }

    // add book to library
    addBookToLibrary(book)
    addBookToLibrary(book2)

    // booking a book from library
    bookABook(book, user)
    bookABook(book, user2)

    getUsersWithBook()
    getBookedBooksOfLibrary()
    fmt.Println(library)
}

func addUserToUserSlice(user User) []User {
    users = append(users, user)
    return users
}

func addBookToBookSlice(book Book) []Book {
    books = append(books, book)
    return books
}

func addBookToLibrary(book Book) {
    library.bookSlice = append(library.bookSlice, book)
}

func registerUserToLibrary(user User) {
    library.users = append(library.users, user)
}

func bookABook(book Book, user User) {
    isBooked := checkIsBooked(book)
    if !isBooked {
       book.isBooked = true
       book.bookedBy = user.id
       user.book = append(user.book, book)
       for id, value := range library.bookSlice {
          if value.ISBN == book.ISBN {
             value.isBooked = true
             value.bookedBy = user.id
             library.bookSlice[id] = value
          }
       }
       for id, value := range library.users {
          if value.id == user.id {
             value.book = append(value.book, book)
             library.users[id] = value
          }

       }
       fmt.Println("book with name", book.name, "and ISBN", book.ISBN, "booked for user", user.name, user.lastName)
    } else {
       fmt.Println("book with name", book.name, "and ISBN", book.ISBN, "has been booked already")
    }
}

func checkIsBooked(book Book) bool {
    for _, value := range library.bookSlice {
       if value.ISBN == book.ISBN && value.isBooked {
          return true
       }
    }
    return false
}

func getAllBooksOfLibrary() []Book {
    return library.bookSlice
}

func getBookedBooksOfLibrary() []Book {
    books := library.bookSlice
    booked := []Book{}
    for _, value := range books {
       if value.isBooked {
          booked = append(booked, value)
       }
    }
    fmt.Println(booked)
    return booked
}

func getUsersWithBook() []User {
    users := library.users
    owner := []User{}
    for _, user := range users {
       if len(user.book) > 0 {
          owner = append(owner, user)
       }
    }
    fmt.Println(owner)
    return owner
}
