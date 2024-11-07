package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Book struct to store information about each book
type Book struct {
	ID     int
	Title  string
	Author string
	Year   int
}

// Books slice to hold list of all books
var books []Book
var bookID int = 1

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\nLibrary Management System")
		fmt.Println("1. Add a Book")
		fmt.Println("2. List All Books")
		fmt.Println("3. Search Book by Title")
		fmt.Println("4. Delete a Book by ID")
		fmt.Println("5. Exit")
		fmt.Print("Choose an option: ")

		scanner.Scan()
		option := scanner.Text()

		switch option {
		case "1":
			addBook(scanner)
		case "2":
			listBooks()
		case "3":
			searchBook(scanner)
		case "4":
			deleteBook(scanner)
		case "5":
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid option. Please choose again.")
		}
	}
}

// Function to add a new book
func addBook(scanner *bufio.Scanner) {
	fmt.Print("Enter book title: ")
	scanner.Scan()
	title := scanner.Text()

	fmt.Print("Enter author name: ")
	scanner.Scan()
	author := scanner.Text()

	fmt.Print("Enter publication year: ")
	scanner.Scan()
	yearStr := scanner.Text()
	year, err := strconv.Atoi(yearStr)
	if err != nil {
		fmt.Println("Invalid year, please enter a valid number.")
		return
	}

	book := Book{
		ID:     bookID,
		Title:  title,
		Author: author,
		Year:   year,
	}

	books = append(books, book)
	bookID++
	fmt.Printf("Book '%s' by %s added successfully!\n", title, author)
}

// Function to list all books
func listBooks() {
	if len(books) == 0 {
		fmt.Println("No books available.")
		return
	}

	fmt.Println("\nList of Books:")
	for _, book := range books {
		fmt.Printf("ID: %d | Title: %s | Author: %s | Year: %d\n", book.ID, book.Title, book.Author, book.Year)
	}
}

// Function to search for a book by title
func searchBook(scanner *bufio.Scanner) {
	fmt.Print("Enter book title to search: ")
	scanner.Scan()
	title := scanner.Text()

	found := false
	for _, book := range books {
		if strings.Contains(strings.ToLower(book.Title), strings.ToLower(title)) {
			fmt.Printf("ID: %d | Title: %s | Author: %s | Year: %d\n", book.ID, book.Title, book.Author, book.Year)
			found = true
		}
	}

	if !found {
		fmt.Println("No book found with that title.")
	}
}

// Function to delete a book by ID
func deleteBook(scanner *bufio.Scanner) {
	fmt.Print("Enter book ID to delete: ")
	scanner.Scan()
	idStr := scanner.Text()
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("Invalid ID.")
		return
	}

	for i, book := range books {
		if book.ID == id {
			books = append(books[:i], books[i+1:]...)
			fmt.Printf("Book '%s' by %s deleted successfully!\n", book.Title, book.Author)
			return
		}
	}
	fmt.Println("Book not found.")
}
