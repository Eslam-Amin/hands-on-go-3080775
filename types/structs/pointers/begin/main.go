// types/structs/pointers/begin/main.go
package main

import "fmt"

type author struct {
	first string
	last  string
}

// fullName returns the full name of the author
func (a author) fullName() string {
	return a.first + " " + a.last
}

// changeName changes the first and last name of the author
//
func (a *author) changeNames(firstName, lastName string) {
	a.first = firstName
	a.last = lastName
}
func main() {
	a := author{
		first: "Marc",
		last:  "Twain",
	}

	fmt.Println(a.fullName())

	// call changeName to update name of author
	//
	a.changeNames("Eslam", "Amin")
	fmt.Println(a.fullName())
}
