package twopointers

import (
	"fmt"
	"github.com/tech-sam/grokking-patterns-go-blind-75/patterns/twopointers/problems"
)

type TwoPointers struct {
	problems map[string]problems.Problem
}

func New() *TwoPointers {
	tp := &TwoPointers{
		problems: make(map[string]problems.Problem),
	}
	tp.registerProblems()
	return tp
}

func (t *TwoPointers) Name() string {
	return "Two Pointers Pattern"
}

func (t *TwoPointers) registerProblems() {
	t.problems["validpalindrome"] = problems.NewValidPalindrome()
}

func (t *TwoPointers) RunProblem(name string) error {
	problem, exists := t.problems[name]
	if !exists {
		return fmt.Errorf("problem %s not found", name)
	}
	return problem.Run()
}

func (t *TwoPointers) ListProblems() []string {
	var problems []string
	for name := range t.problems {
		problems = append(problems, name)
	}
	return problems
}
