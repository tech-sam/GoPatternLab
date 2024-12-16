package twopointers

import (
	"fmt"
	"github.com/tech-sam/grokking-patterns-go-blind-75/patterns/twopointers/problems"
	"github.com/tech-sam/grokking-patterns-go-blind-75/pkg/problem"
)

type TwoPointers struct {
	problems map[string]problem.Problem
}

func New() *TwoPointers {
	tp := &TwoPointers{
		problems: make(map[string]problem.Problem),
	}
	tp.registerProblems()
	return tp
}

func (t *TwoPointers) Name() string {
	return "Two Pointers Pattern"
}

func (t *TwoPointers) registerProblems() {
	t.problems["validpalindrome"] = problems.NewValidPalindrome()
	t.problems["sortcolors"] = problems.NewSortColors()
	t.problems["min_moves_for_palindrome"] = problems.NewMinMovesForPalindrome()
}

func (t *TwoPointers) RunProblem(name string) error {
	p, exists := t.problems[name]
	if !exists {
		return fmt.Errorf("problem %s not found", name)
	}
	return p.Run()
}

func (t *TwoPointers) ListProblems() []string {
	var p []string
	for name := range t.problems {
		p = append(p, name)
	}
	return p
}
