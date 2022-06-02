package main

import (
	"fmt"
)

func input(str, str2 string) string {
	var result string
	// looping both inputted string
	for _, i := range str {
		for _, j := range str2 {
			if i != j {
				result = string(j) // if letter in one string doesnt contains letter in the other strings, print the different
			}
			if i != j {
				result = string(i) // if letter in one string doesnt contains letter in the other strings, print the different
			}
		}
	}

	return result
}

func main() {
	fmt.Println(input("bxcn", "abncx"))
}
