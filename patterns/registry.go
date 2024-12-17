package patterns

import (
	"fmt"
	"github.com/tech-sam/GoPatternLab/patterns/twopointers"
)

type Pattern interface {
	Name() string
	RunProblem(name string) error
	ListProblems() []string
}

type Registry struct {
	patterns map[string]Pattern
}

func NewRegistry() *Registry {
	r := &Registry{
		patterns: make(map[string]Pattern),
	}
	r.registerPatterns()
	return r
}

func (r *Registry) registerPatterns() {
	r.patterns["twopointers"] = twopointers.New()
}

func (r *Registry) RunProblem(pattern, problem string) error {
	p, exists := r.patterns[pattern]
	if !exists {
		return fmt.Errorf("pattern %s not found", pattern)
	}
	return p.RunProblem(problem)
}

func (r *Registry) ListPatterns() map[string][]string {
	result := make(map[string][]string)
	for name, pattern := range r.patterns {
		result[name] = pattern.ListProblems()
	}
	return result
}
