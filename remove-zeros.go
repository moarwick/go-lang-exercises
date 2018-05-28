package main

import (
	"fmt"
)

func argsToSlice(args ...interface{}) []interface{} {
	var nonZeros []interface{}
	var zeros []interface{}
	for _, val := range args {
		if val == 0 {
			zeros = append(zeros, val)
		} else {
			nonZeros = append(nonZeros, val)
		}
	}
	return append(nonZeros, zeros...)
}

func main() {
	fmt.Println(argsToSlice(1))                       // returns [1]
	fmt.Println(argsToSlice(2, 4, 1, 8, 2))           // returns [2, 4, 1, 8, 2]
	fmt.Println(argsToSlice(2, 0, 4, 1, 0, 8, 2))     // returns [2, 4, 1, 8, 2, 0, 0]
	fmt.Println(argsToSlice(1))                       // returns [1]
	fmt.Println(argsToSlice(2, "a", 1, "b", 2))       // returns [2, "a", 1, "b", 2]
	fmt.Println(argsToSlice(2, 0, "a", 1, 0, "b", 2)) // returns [2, "a", 1, "b", 2, 0, 0]
}
