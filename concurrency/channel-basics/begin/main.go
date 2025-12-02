// concurrency/channels/begin/main.go
package main

import "fmt"

// sum calculates and prints the sum of numbers
func sum(nums []int, ch chan<- int) {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	// fmt.Println("Result:", sum)
	ch <- sum
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	ch := make(chan int)
	fmt.Println(nums)
	// invoke the sum function as a goroutine
	go sum(nums, ch)

	result := <-ch
	fmt.Println("Result:", result)

	ch2 := make(chan string, 2)
	ch2 <- "James"
	fmt.Println(<-ch2)
	ch2 <- "Eca"
	fmt.Println(<-ch2)
	demo()
}
func demo() {
	x := 10
	defer fmt.Println(x) // x's value captured now
	x = 20
}
