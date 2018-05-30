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
	{[]string{"Amazing", "Waffles"}, "Crap Pankaces"}, // failing test
}

func TestConcat(t *testing.T) {
	for i, testCase := range testData {
		actual := concat(testCase.args...)
		if actual != testCase.expected {
			error := fmt.Sprintf("%d) Expected: %s, Got: %s", i, testCase.expected, actual)
			t.Error(error)
		}
	}
}
