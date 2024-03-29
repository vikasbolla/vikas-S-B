package solutions

import (
	"sort"
	"strings"
)

// func GroupAnagrams(words []string) [][]string {
// 	anagramsMap := make(map[string][]string)

// 	for _, word := range words {
// 		// Convert the word to a sorted string to identify anagrams
// 		sortedWord := sortString(word)

// 		// Append the word to the corresponding group in the map
// 		anagramsMap[sortedWord] = append(anagramsMap[sortedWord], word)
// 	}

// 	// Convert the map values to a slice of slices
// 	var result [][]string
// 	for _, group := range anagramsMap {
// 		result = append(result, group)
// 	}

// 	return result
// }

// func sortString(s string) string {
// 	// Convert the string to a slice of characters
// 	characters := strings.Split(s, "")
// 	// Sort the slice of characters
// 	sort.Strings(characters)
// 	// Join the sorted characters back into a string
// 	return strings.Join(characters, "")
// }

func GroupAnagrams(strs []string) [][]string {

	anagramsMap := make(map[string][]string)
	var result [][]string
	for _, word := range strs {
		characters := strings.Split(word, "")
		sort.Strings(characters)
		anagramsMap[strings.Join(characters, "")] = append(anagramsMap[strings.Join(characters, "")], word)

	}
	for _, groups := range anagramsMap {
		result = append(result, groups)
	}

	return result
}
