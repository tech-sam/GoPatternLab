package problems

import (
	"fmt"
	"github.com/tech-sam/grokking-patterns-go-blind-75/pkg/problem"
)

func NewMinMovesForPalindrome() problem.Problem {
	return problem.NewProblem("Minimum Number of Moves to Make Palindrome", TestMinMovesToMakePalindrome)
}

func TestMinMovesToMakePalindrome() error {
	testCases := []struct {
		input    string
		expected int
	}{
		{"ccxx", 2},
		{"mamad", 3},
		{"eggeekgbbeg", 8},
	}

	for _, tc := range testCases {
		result := minMovesToMakePalindrome(tc.input)
		fmt.Printf("\nInput: %q\n", tc.input)
		fmt.Printf("Output: %v\n", result)
		if tc.expected != result {
			fmt.Printf("expected %d got %d\n ", tc.expected, result)
		}
	}
	return nil
}

func minMovesToMakePalindrome(s string) int {
	runes := []rune(s)
	moves := 0
	for i, j := 0, len(runes)-1; i < j; i++ {
		k := j
		for k > i {
			if runes[i] == runes[k] {
				for k < j {
					runes[k], runes[k+1] = runes[k+1], runes[k]
					moves++
					k++
				}
				j--
				break
			}
			k--
		}
		if k == i {
			moves += len(runes)/2 - i
		}
	}
	return moves
}
