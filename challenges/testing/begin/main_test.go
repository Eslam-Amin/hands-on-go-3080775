// challenges/testing/begin/main_test.go
package main

import "testing"

type test struct {
	input string
	want  int
}

// write a test for letterCounter.count
func TestLetterCount(t *testing.T) {
	tests := map[string]test{
		"empty": {input: "", want: 0},
		"one":   {input: "a2 32 ^ &/)", want: 1},
		"two":   {input: "812 %6ab//", want: 2},
	}
	lc := letterCounter{}
	for _, val := range tests {
		t.Run(val.input, func(t *testing.T) {
			if got := lc.count(val.input); got != val.want {
				t.Errorf("got %v want %v", got, val.want)
			} else {
				t.Logf("got %v want %v", got, val.want)
			}
		})
	}
}

// write a test for numberCounter.count
func TestNumberCount(t *testing.T) {
	tests := map[string]test{
		"empty": {input: "", want: 0},
		"one":   {input: "a2 32 ^ &/)", want: 3},
		"two":   {input: "812 %6ab//", want: 4},
	}

	nc := numberCounter{}
	for _, val := range tests {
		t.Run(val.input, func(t *testing.T) {
			if got := nc.count(val.input); got != val.want {
				t.Errorf("got %v want %v", got, val.want)
			} else {
				t.Logf("got %v want %v", got, val.want)
			}
		})

	}

}

// write a test for symbolCounter.count
func TestSymbolCount(t *testing.T) {
	tests := map[string]test{
		"empty": {input: "", want: 0},
		"one":   {input: "a2 32 ^ &/)", want: 7},
		"two":   {input: "812 %6ab//", want: 4},
	}
	sc := symbolCounter{}
	for _, val := range tests {
		t.Run(val.input, func(t *testing.T) {
			if got := sc.count(val.input); got != val.want {
				t.Errorf("got %v want %v", got, val.want)
			} else {
				t.Logf("got %v want %v", got, val.want)
			}
		})
	}

}
