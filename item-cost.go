package main

import (
	"fmt"
	"math"
)

type item struct {
	name     string
	quantity int
	price    float64
}

var itemsA = []item{
	item{name: "shoes", quantity: 2, price: 19.95},
	item{name: "handbag", quantity: 1, price: 39.95},
	item{name: "shirt", quantity: 3, price: 29.95},
}

var itemsB = []item{
	item{name: "shoes", quantity: -2, price: 19.95},
	item{name: "handbag", quantity: 1, price: 39.95},
	item{name: "shirt", quantity: -3, price: 29.95},
}

func total(items []item) float64 {
	var total float64
	for _, item := range items {
		total += math.Abs(float64(item.quantity)) * item.price
	}
	return total
}

func main() {
	fmt.Println(total(itemsA)) // returns 169.70
	fmt.Println(total(itemsB)) // returns 169.70
}
