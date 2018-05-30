package main

import (
	"fmt"
	"testing"
)

var testData = []struct {
	args     []string
	expected string
}{
	{[]string{"Hello", "World"}, "HelloWorld"}, // why do we need to type args again??
	{[]string{"Shine", "On", "You", "Crazy", "Diamond"}, "ShineOnYouCrazyDiamond"},
}

func TestConcat(t *testing.T) {
	for _, testCase := range testData {
		actual := concat(testCase.args...)
		if actual != testCase.expected {
			error := fmt.Sprintf("Expected: %s, Got: %s", testCase.expected, actual)
			t.Error(error)
		}
	}
}
