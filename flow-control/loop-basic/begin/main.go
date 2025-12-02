// flow-control/loop-basic/begin/main.go
package main

import "fmt"

func main() {
	// declare a string to iterate over
	s := "Hello"
	nums := []int{1, 2, 3, 4, 5, 6, 10}
	numsMap := map[string]int{"one": 1, "two": 2, "three": 3}
	for i := 0; i < len(s); i++ {
		fmt.Print(s[i])
	}
	fmt.Println("")
	// iterate over the string with basic for loop
	for i := range s {
		fmt.Print(string(s[i]))
	}
	fmt.Println("")

	for i, val := range nums {
		fmt.Printf("%d: %v\n", i, val)
	}
	fmt.Println("")

	for key, val := range numsMap {
		fmt.Printf("%s: %v\n", key, val)
	}
}
