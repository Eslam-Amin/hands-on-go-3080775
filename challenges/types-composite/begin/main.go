// challenges/types-composite/begin/main.go
package main

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"
)

// define an author type with a name
type author struct {
	name string
}

// define a book type with a title and author
type book struct {
	title  string
	author author
}

// define a library type to track a list of books
type library map[string][]book

// define addBook to add a book to the library
func (l library) addBook(b book) {
	l[b.author.name] = append(l[b.author.name], b)
}

// define a lookupByAuthor function to find books by an author's name
func (l library) lookupByAuthorName(name string) ([]book, bool) {
	books, ok := l[name]
	return books, ok
}
func main() {
	// create a new library
	lib := library{}

	// add 2 books to the library by the same author
	jb := author{name: "James Baldwin"}
	tm := author{name: "Toni Morrison"}
	// add 1 book to the library by a different author
	lib.addBook(book{
		author: jb,
		title:  "The Fire Next Time",
	})

	lib.addBook(book{
		author: jb,
		title:  "Go Tell It on the Mountain",
	})

	lib.addBook(book{
		author: author{name: "Marcus Aurelius"},
		title:  "Meditations",
	})

	// dump the library with spew
	spew.Dump(lib)

	// lookup books by known author in the library
	books, ok := lib.lookupByAuthorName(jb.name)

	// print out the first book's title and its author's name
	if !ok {
		fmt.Printf("No books by %v\n", jb.name)
	} else {
		spew.Dump(books)
	}
	books, ok = lib.lookupByAuthorName(tm.name)
	if !ok {
		fmt.Printf("No books by %v\n", tm.name)
	} else {
		spew.Dump(books)
	}
}
