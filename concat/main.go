package main

import (
	"fmt"
	"strings"
)

func concat(words ...string) string {
	return strings.Join(words, "")
}

func main() {
	fmt.Println(concat("Amazing", "Waffles"))
	fmt.Println(concat("Shine", "On", "You", "Crazy", "Diamond"))
}
