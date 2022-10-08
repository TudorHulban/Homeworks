package main

type state struct {
	process      []string
	dependencies map[string][]string
}

func newState() *state {
	return &state{
		process:      make([]string, 0),
		dependencies: make(map[string][]string),
	}
}

func (s *state) extractDependency(name string) {
	if dependencies, exists := s.dependencies[name]; exists {
		for _, stage := range dependencies {
			s.process = append(s.process, stage)

			s.extractDependency(stage)
		}

		return
	}
}

func (s *state) load(stages map[string][]string) {
	for stage, dependencies := range stages {
		if len(dependencies) == 0 {
			s.process = append(s.process, stage)

			s.extractDependency(stage)

			continue
		}

		if contains(s.process, dependencies) {
			s.process = append(s.process, stage)

			s.extractDependency(stage)

			continue
		}

		for _, dependency := range dependencies {
			s.dependencies[dependency] = append(s.dependencies[dependency], stage)
		}
	}
}

func main() {}
