package anagram

import (
	"slices"
	"strings"
)

// Helper function to turn a string into a sorted, lowercase "fingerprint"
func getFingerprint(s string) string {
	runes := []rune(strings.ToLower(s))
	slices.Sort(runes) // Sorts the letters alphabetically in one line
	return string(runes)
}

func Detect(subject string, candidates []string) []string {
	subjectLower := strings.ToLower(subject)
	subjectFingerprint := getFingerprint(subject)
	
	var anagrams []string

	for _, candidate := range candidates {
		candidateLower := strings.ToLower(candidate)

		// Rule 1: A word is not its own anagram
		if candidateLower == subjectLower {
			continue
		}

		// Rule 2: If their sorted letters match, it's an anagram!
		if getFingerprint(candidate) == subjectFingerprint {
			anagrams = append(anagrams, candidate)
		}
	}

	return anagrams
}