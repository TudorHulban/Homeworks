package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var root = map[string][]string{
	"checkout": {},
}

var easy = map[string][]string{
	"checkout": {},
	"lint":     {"checkout"},
}

var block1 = map[string][]string{
	"checkout": {},
	"deploy":   {"publish"},
}

var block2 = map[string][]string{
	"checkout": {},
	"deploy":   {"publish"},
	"lint":     {"checkout"},
}

func TestParseStages(t *testing.T) {
	testCases := []struct {
		description string
		input       map[string][]string
		want        []string
	}{
		{"root", root, []string{"checkout"}},
		{"easy", easy, []string{"checkout", "lint"}},
		{"remaining stages blocking", block1, []string{"checkout"}},
		{"remaining stages unblocking", block2, []string{"checkout", "lint"}},
		{"full", full, []string{"checkout", "lint", "build", "test", "package", "publish", "deploy"}},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			assert.True(t, contains(parseStages(tc.input), tc.want))
		})
	}
}

func TestWErrorCorrection(t *testing.T) {
	stage1 := stage{
		name: "checkout",
	}

	stage2 := stage{
		name:         "lint",
		dependencies: []string{"checkout"},
	}

	stage3 := stage{
		name:         "build",
		dependencies: []string{"checkout"},
	}

	stage4 := stage{
		name:         "test",
		dependencies: []string{"checkout"},
	}

	stage5 := stage{
		name:         "package",
		dependencies: []string{"build"},
	}

	stage6 := stage{
		name:         "publish",
		dependencies: []string{"test", "lint", "package"},
	}

	stage7 := stage{
		name:         "deploy",
		dependencies: []string{"publish"},
	}

	testCases := []struct {
		description string
		input       process
		want        []string
	}{
		{"checkout", process{&stage1}, []string{"checkout"}},
		{"deploy, checkout", process{&stage7, &stage1}, []string{"checkout"}},
		{"lint, checkout", process{&stage2, &stage1}, []string{"checkout", "lint"}},
		{"lint, deploy, checkout", process{&stage2, &stage7, &stage1}, []string{"checkout", "lint"}},
		{"full1", process{&stage4, &stage5, &stage6, &stage7, &stage1, &stage2, &stage3}, []string{"checkout", "test", "lint", "build", "package", "publish", "deploy"}},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			result := parse(tc.input)
			t.Log("result: ", result)

			assert.Equal(t, tc.want, result)
		})
	}
}
