package anagram

import (
	"sort"
	"strings"
)

func SortString(s string) string {
	sortedChars := strings.Split(s, "")
	sort.Strings(sortedChars)
	return strings.Join(sortedChars, "")
}

func GroupAnagrams(words []string) [][]string {
	var result [][]string

	for i := 0; i < len(words); i++ {
		if words[i] != "" {
			anagramGroup := []string{words[i]}
			sortedWord := SortString(words[i])
			for j := i + 1; j < len(words); j++ {
				if words[j] != "" && sortedWord == SortString(words[j]) {
					anagramGroup = append(anagramGroup, words[j])
					words[j] = "" 
				}
			}
			result = append(result, anagramGroup)
		}
	}
	return result
}