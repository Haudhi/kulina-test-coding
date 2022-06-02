package main

import (
	"fmt"
)

// mapping all the list of roman numerals code
var decoder = map[rune]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

// making function to convert roman to integer

func Roman(roman string) int {
	if len(roman) == 0 {
		return 0
	}
	first := decoder[rune(roman[0])]
	if len(roman) == 1 {
		return first
	}
	next := decoder[rune(roman[1])]
	if next > first {
		return (next - first) + Roman(roman[2:])
	}
	return first + Roman(roman[1:])
}

func main() {
	fmt.Println(Roman("XII"))
}
