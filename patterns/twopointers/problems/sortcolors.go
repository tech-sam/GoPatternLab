package problems

import (
	"fmt"
	"github.com/tech-sam/GoPatternLab/pkg/problem"
)

func NewSortColors() problem.Problem {
	return problem.NewProblem("Sort color (Blind 75 #111)", runSortColors)
}

func runSortColors() error {
	fmt.Println("sort colors called...")
	return nil
}
