// types/structs/fields/begin/main.go
package main

import "fmt"

// define a struct type for author
//
type author struct {
	firstName, lastName string
}

func main() {
	// intialize author
	//
	a := author{
		"Eslam",
		"Amin",
	}

	// print the author, type, and value
	fmt.Printf("%#v\n", a)
}
