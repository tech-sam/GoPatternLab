package problems

import (
	"fmt"
	"github.com/tech-sam/GoPatternLab/pkg/problem"
)

func NewValidPalindrome() problem.Problem {
	return problem.NewProblem("Valid Palindrome (Blind 75 #125)", runValidPalindrome)
}

func runValidPalindrome() error {
	testCases := []struct {
		input    string
		expected bool
	}{
		{"A man, a plan, a canal: Panama", true},
		{"race a car", false},
	}

	for _, tc := range testCases {
		result := isPalindrome(tc.input)
		fmt.Printf("\nInput: %q\n", tc.input)
		fmt.Printf("Output: %v\n", result)
	}
	return nil

}
func isPalindrome(s string) bool {
	return true
}
