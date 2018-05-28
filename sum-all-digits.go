package main

import (
	"fmt"
	"strconv"
)

func sumDigitsInStr(str string) int {
	sum := 0
	for _, num := range []rune(str) {
		sum += int(num - 48)
	}
	return sum
}

func sumDigits(val interface{}) int {
	switch val.(type) {
	case int:
		return sumDigitsInStr(strconv.Itoa(val.(int)))
	case string:
		return sumDigitsInStr(val.(string))
	default:
		return 0
	}
}

func main() {
	fmt.Println(sumDigits(4))       // returns 4
	fmt.Println(sumDigits("4"))     // returns 4
	fmt.Println(sumDigits(22))      // returns 4
	fmt.Println(sumDigits("22"))    // returns 4
	fmt.Println(sumDigits(12345))   // returns 15
	fmt.Println(sumDigits("12345")) // returns 15
}
