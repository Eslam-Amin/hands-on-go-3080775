// interfaces/type-assertions/begin/main.go
package main

import "fmt"

func main() {
	// perform a type assertion
	var i any = "hello"
	// str := i.(string)
	fmt.Println(i.(string))
	// perform a type assertion and handle the error
	str, ok := i.(string)
	if !ok {
		fmt.Printf("%v is not a string\n", i)
		fmt.Println(str)
	}
}
