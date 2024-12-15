package problems

import (
	"fmt"
	"strings"
	"unicode"
)

type Problem interface {
	Name() string
	Run() error
}

type ValidPalindrome struct{}

func NewValidPalindrome() *ValidPalindrome {
	return &ValidPalindrome{}
}

func (v *ValidPalindrome) Name() string {
	return "Valid Palindrome (Blind 75 #125)"
}

func (v *ValidPalindrome) Run() error {
	testCases := []struct {
		input    string
		expected bool
	}{
		{"A man, a plan, a canal: Panama", true},
		{"race a car", false},
		{"", true},
	}

	for _, tc := range testCases {
		result := v.isPalindrome(tc.input)
		fmt.Printf("\nInput: %q\n", tc.input)
		fmt.Printf("Output: %v\n", result)
		fmt.Printf("Expected: %v\n", tc.expected)
	}
	return nil
}

func (v *ValidPalindrome) isPalindrome(s string) bool {
	// Convert to lowercase
	s = strings.ToLower(s)

	// Two pointers approach
	left, right := 0, len(s)-1

	for left < right {
		// Skip non-alphanumeric characters from left
		for left < right && !unicode.IsLetter(rune(s[left])) && !unicode.IsNumber(rune(s[left])) {
			left++
		}
		// Skip non-alphanumeric characters from right
		for left < right && !unicode.IsLetter(rune(s[right])) && !unicode.IsNumber(rune(s[right])) {
			right--
		}

		// Compare characters
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}
