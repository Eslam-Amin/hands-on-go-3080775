// challenges/generics/begin/main.go
package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type numeric interface {
	constraints.Integer | constraints.Float
}

// Part 1: print function refactoring
func genericPrint[T any](val T) {
	fmt.Println(val)
}

// non-generic print functions

func printString(s string) { fmt.Println(s) }

func printInt(i int) { fmt.Println(i) }

func printBool(b bool) { fmt.Println(b) }

// Part 2 sum function refactoring
func genericSum[T numeric](numbers ...T) T {
	var result T = 0
	for _, val := range numbers {
		result += val
	}
	return result
}

// sum sums a slice of any type
func sum(numbers []interface{}) interface{} {
	var result float64
	for _, n := range numbers {
		switch n.(type) {
		case int:
			result += float64(n.(int))
		case float32, float64:
			result += n.(float64)
		default:
			continue
		}
	}
	return result
}

func main() {
	// call non-generic print functions
	// printString("Hello")
	// printInt(42)
	// printBool(true)

	// call generic printAny function for each value above
	genericPrint("Hello")
	genericPrint(42)
	genericPrint(true)

	// call sum function
	// fmt.Println("result", sum([]interface{}{1, 2, 3}))

	// call generics sumAny function
	// fmt.Println("result", genericSum([]int{1, 2, 3}))
	fmt.Println("result", genericSum(1.1, 2.2, 3.3))
}
