package main

import "fmt"

func print(num int) string {
	var result string
	// looping the input so it can show number from 1 till the inputted integer
	for i := 1; i <= num; i++ {
		if i%3 == 0 {
			fmt.Println("Kulina") // if the number between 1 and the inputted number is divisible by 3, return string Kulina
		}
		if i%5 == 0 {
			fmt.Println("Food") // if the number between 1 and the inputted number is divisible by 5, return string Food
		}
		if i%5 == 0 && i%3 == 0 {
			fmt.Println("Kulina x Food") // if the number between 1 and the inputted number is divisible by 3 and 5, return string Kulina x Food
		} else {
			fmt.Println(i) // Otherwise, print the integer
		}
	}
	return result
}

func main() {
	fmt.Println(print(15))
}
