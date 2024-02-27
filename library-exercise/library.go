
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
