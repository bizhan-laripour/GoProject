package main

import "fmt"
/**
with pointers latency will decrease and i dont know why!!!!
*/
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
    bookABook(&book, &user)
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
    user     User
    name     string
    ISBN     string
    isBooked bool
    bookedBy int
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

func bookABook(book *Book, user *User) {
    pointer := &library
    for _, value := range pointer.books {
       if value.ISBN == book.ISBN && value.isBooked {
          fmt.Println("this book is booked")
          return
       }
    }
    book.bookedBy = user.id
    book.user = *user
    book.isBooked = true
    pointer.books = append(pointer.books, *book)
    fmt.Println("book booked")
}
