package main

import "fmt"

func main() {
    book := Book{
       name: "golang",
       ISBN: "123456789",
    }

    user := User{
       name:     "Bizhan",
       lastName: "Laripour",
       id:       1,
    }
    addUserToUser(&user)
    addBookToLibrary(&book)
    fmt.Println(library)
}

type Library struct {
    name  string
    users []User
    books []Book
}

type User struct {
    id       int
    name     string
    lastName string
}

type Book struct {
    user User
    name string
    ISBN string
}

var books []Book
var library Library

var users []User

func addBookToLibrary(book *Book) Book {
    books = append(books, *book)
    lib := &library
    lib.books = books
    return *book
}

func addUserToUser(user *User) User {
    users = append(users, *user)
    lib := &library
    lib.users = users
    return *user
}


func bookABook(book *Book , user *User){
    
}
