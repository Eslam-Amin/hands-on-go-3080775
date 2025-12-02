// flow-control/defer-panic-recover/begin/main.go
package main

import "fmt"

func cleanup(msg string) {
	fmt.Println(msg)
}

func main() {
	// defer function calls
	defer cleanup("first")
	defer cleanup("second")
	// defer recovery
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recovered from panic:", err)
		}
	}()
	// panic
	cleanup("running in main...")

	panic("something bad happened")
}
