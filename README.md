## Running the Project
// Book represents a book in the library with ID, Title, Author, and Year.
type Book struct {
	ID     int
	Title  string
	Author string
	Year   int
}

// addBook adds a new book to the library collection.
func addBook(scanner *bufio.Scanner) { ... }
``` bash 
go run library_manager.go
