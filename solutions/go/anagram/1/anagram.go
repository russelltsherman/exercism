package anagram

import (
    "slices"
	"strings"
)

func Downcase(word string) string {
    // Convert to lowercase
    return strings.ToLower(word)
}

func SortWord(word string) string {
    // Convert the string to a slice of runes (handles multi-byte Unicode characters)
	runes := []rune(Downcase(word))
	// Sort the slice in-place
	slices.Sort(runes)
	// Convert the sorted runes back to a string
	sortedStr := string(runes)  
    return sortedStr
}

func IsAnagram(subject, candidate string) bool {
    if Downcase(subject) == Downcase(candidate) {
        return false // same words are not anagrams
    }
    if SortWord(subject) == SortWord(candidate) {
        return true
    }  
    return false
}

func Detect(subject string, candidates []string) []string {
	anagrams := []string{}
    for _, candidate := range candidates {
        if IsAnagram(subject, candidate) {
        	anagrams = append(anagrams, candidate) 
        }
    }
	return anagrams
}
