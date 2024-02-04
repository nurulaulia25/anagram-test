package main

import (
	"anagram-backend-test/anagram"
	"fmt"
	"strings"
)

func main() {
	input := []string{"cook", "save", "taste", "aves", "vase", "state", "map"}
	result := anagram.GroupAnagrams(input)

	fmt.Println("[")
	for _, group := range result {
		fmt.Printf("  [ '%s' ],\n", strings.Join(group, "', '"))
	}
	fmt.Println("]")
}
