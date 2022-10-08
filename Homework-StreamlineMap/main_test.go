package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var full = map[string][]string{
	"checkout": {},                          // Clone the repository
	"lint":     {"checkout"},                // Run a linter
	"build":    {"checkout"},                // Build binaries
	"test":     {"checkout"},                // Run tests
	"package":  {"build"},                   // Package the binaries into a package
	"publish":  {"test", "lint", "package"}, // Publish the packages to a server for storing build artifacts
	"deploy":   {"publish"},                 // Pick the package from where it was published and deploy it
}

func TestLoad(t *testing.T) {
	testCases := []struct {
		description string
		input       map[string][]string
		want        []string
	}{
		{"1. no dependencies - checkout only",
			map[string][]string{
				"checkout": {},
			},
			[]string{"checkout"}},
		{"2. easy - checkout - lint",
			map[string][]string{
				"checkout": {},
				"lint":     {"checkout"},
			},
			[]string{"checkout", "lint"}},
		{"3. hard - lint - checkout",
			map[string][]string{
				"lint":     {"checkout"},
				"checkout": {},
			},
			[]string{"checkout", "lint"}},
		{"4. hard - lint -build - checkout",
			map[string][]string{
				"lint":     {"checkout"},
				"build":    {"checkout"},
				"checkout": {},
			},
			[]string{"checkout", "lint", "build"}},
		{"5. remaining stages blocking",
			map[string][]string{
				"checkout": {},
				"deploy":   {"publish"},
			},
			[]string{"checkout"}},
		{"6. remaining stages unblocking", map[string][]string{
			"checkout": {},
			"deploy":   {"publish"},
			"lint":     {"checkout"},
		},
			[]string{"checkout", "lint"}},
		{"7. full",
			full,
			[]string{"checkout", "lint", "build", "test", "package", "publish", "deploy"},
		},
		{"8. dependencies not met",
			map[string][]string{
				"deploy": {"publish"},
				"lint":   {"checkout"},
			},
			[]string{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			state := newState()
			state.load(tc.input)

			t.Logf("state process: %s", state.process)

			assert.True(t, contains(state.process, tc.want))
		})
	}
}
