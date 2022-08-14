package main

import (
	"fmt"
	"strings"
)

type stage struct {
	name         string
	dependencies []string
}

type process []*stage

var full = map[string][]string{
	"checkout": {},                          // Clone the repository
	"lint":     {"checkout"},                // Run a linter
	"build":    {"checkout"},                // Build binaries
	"test":     {"checkout"},                // Run tests
	"package":  {"build"},                   // Package the binaries into a package
	"publish":  {"test", "lint", "package"}, // Publish the packages to a server for storing build artifacts
	"deploy":   {"publish"},                 // Pick the package from where it was published and deploy it
}

func (p process) String() string {
	var res []string

	for _, stage := range p {
		res = append(res, stage.name+"-dependencies{"+strings.Join(stage.dependencies, ",")+"}")
	}

	return strings.Join(res, "\n")
}

func main() {
	p := identifySteps(full)

	fmt.Println(p)

	fmt.Println("streamlined:", parse(p))
}

func identifySteps(stages map[string][]string) process {
	var res process

	for name, dependencies := range stages {
		res = append(res, &stage{
			name:         name,
			dependencies: dependencies,
		})
	}

	return res
}

func parse(p process) []string {
	if len(p) == 0 {
		return nil
	}

	var res []string

	var i int
	for i <= len(p)-1 {
		if len(p[i].dependencies) == 0 {
			res = append(res, p[i].name)

			p = append(p[:i], p[i+1:]...)

			if len(p) == 0 {
				break
			}

			i = 0

			continue
		}

		if contains(res, p[i].dependencies) {
			res = append(res, p[i].name)

			p = append(p[:i], p[i+1:]...)

			if len(p) == 0 {
				break
			}

			i = 0

			continue
		}

		i++
	}

	return res
}

func parseStages(stages map[string][]string) []string {
	steps := identifySteps(stages)

	return parse(steps)
}

func contains(source, what []string) bool {
	for _, elementWhat := range what {
		var elementIsContained bool

		for _, elementSource := range source {
			if elementSource == elementWhat {
				elementIsContained = true

				break
			}
		}

		if !elementIsContained {
			return false
		}
	}

	return true
}
