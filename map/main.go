package main

import "fmt"

// mapping function, iterates over ints, applies passed-in fn, returns array of floats
func mapper(nums []int, fn func(int) float64) []float64 {
	result := make([]float64, len(nums))
	for i, num := range nums {
		result[i] = fn(num)
	}
	return result
}

// return multiplying function (stores 'x' multiplier in closure)
func multi(x interface{}) func(int) float64 {
	return func(num int) float64 {
		switch x.(type) {
		case int:
			return float64(num) * float64(x.(int))
		case float64:
			return float64(num) * x.(float64)
		default:
			return 0
		}
	}
}

func main() {
	fmt.Println(mapper([]int{1, 2, 3}, multi(2)))    // [3 6 9]
	fmt.Println(mapper([]int{1, 2, 3}, multi(2.5)))  // [2.5 5 7.5]
	fmt.Println(mapper([]int{1, 2, 3}, multi(10.1))) // [10.1 20.2 30.299999999999997]
}
