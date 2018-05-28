package main

import (
	"fmt"
)

func fizzBuzz(val int) {
	for num := 1; num <= val; num++ {
		switch {
		case num%3 == 0 && num%5 == 0:
			fmt.Println("Fizz Buzz")

		case num%3 == 0:
			fmt.Println("Fizz")

		case num%5 == 0:
			fmt.Println("Buzz")

		default:
			fmt.Println(num)
		}
	}
}

func main() {
	fizzBuzz(5)  // returns 1, 2, Fizz, 4, Buzz
	fizzBuzz(15) // returns 1, 2, Fizz, 4, Buzz, Fizz, 7, 8, Fizz, Buzz, 11, Fizz, 13, 14, FizzBuzz
}
