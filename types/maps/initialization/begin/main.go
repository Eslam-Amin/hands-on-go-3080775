// types/maps/initialization/begin/main.go
package main

import "fmt"

type author struct {
	name string
}

func main() {
	// declare a map of string keys and author values
	// var authors map[string]author

	// initialize the map with make
	// authors = make(map[string]author)
	authors := map[string]author{
		"tm": {name: "Toni Morrison"},
		"ma": {name: "Marcus Aurelius"},
	}
	// add authors to the map
	// authors["tm"] = author{name: "Toni Morrison"}
	// authors["ma"] = author{name: "Marcus Aurelius"}

	// print the map with fmt.Printf
	fmt.Printf("%#v\n", authors)

	// read a value from the map with a known key
	tmAuthor := authors["tm"]
	jmAuthor, ok := authors["jm"]
	if !ok {
		fmt.Println("Jame Baldwin is not in the map")
	} else {
		fmt.Println(jmAuthor)
	}
	fmt.Println(tmAuthor)
}
