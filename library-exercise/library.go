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
    //register user
    registerUserToLibrary(user)
    
    // create book
    book := Book{
       name: "golang",
       ISBN: "123654789",
       isBooked: false,
    }
    
    // add book to library
    addBookToLibrary(book)
    bookABook(book , user);
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

func bookABook(book Book , user User)  {
    book.isBooked = true
    book.bookedBy = user.id
    user.book = append(user.book , book)
    for _ , value := range library.bookSlice{
       if(value.ISBN == book.ISBN){
          value.isBooked = true
          value.bookedBy = user.id
       }
    }
}
