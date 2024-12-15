package problem

type Problem interface {
	Name() string
	Run() error
}

type BaseProblem struct {
	name     string
	solution func() error
}

func NewProblem(name string, solution func() error) Problem {
	return &BaseProblem{
		name:     name,
		solution: solution,
	}
}

func (b *BaseProblem) Name() string {
	return b.name
}

func (b *BaseProblem) Run() error {
	return b.solution()
}
