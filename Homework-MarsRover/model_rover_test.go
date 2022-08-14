package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValidationDirection(t *testing.T) {
	c := NewRover(6, 4, "EASt")

	require.Nil(t, c)
}

// validate the inital state
func TestNewRover(t *testing.T) {
	c := NewRover(6, 4, "EAST")

	require.NotNil(t, c)

	require.Equal(t, c.x, 6)
	require.Equal(t, c.y, 4)
	require.Equal(t, c.direction, "EAST")
}

func TestActualPos(t *testing.T) {
	c := NewRover(6, 4, "NORTH")

	require.Equal(t, c.ActualPos(), "(6, 4) NORTH")
}

func TestReceive(t *testing.T) {
	c := NewRover(6, 4, "NORTH")

	c.Receive("")
	require.Equal(t, c.ActualPos(), "(6, 4) NORTH")

	c.Receive("F")
	require.Equal(t, c.ActualPos(), "(6, 5) NORTH", "moves forward")
}

func TestMoveForwardNorth(t *testing.T) {
	c := NewRover(6, 4, "NORTH")

	c.Receive("F")
	require.Equal(t, c.ActualPos(), "(6, 5) NORTH", "moves forward to NORTH")
}

func TestMoveForwardWest(t *testing.T) {
	c := NewRover(6, 4, "WEST")

	c.Receive("F")
	require.Equal(t, c.ActualPos(), "(5, 4) WEST", "moves forward to WEST")
}

func TestMoveForwardSouth(t *testing.T) {
	c := NewRover(6, 4, "SOUTH")

	c.Receive("F")
	require.Equal(t, c.ActualPos(), "(6, 3) SOUTH", "moves forward to SOUTH")
}

func TestMoveForwardEast(t *testing.T) {
	c := NewRover(6, 4, "EAST")

	c.Receive("F")
	require.Equal(t, c.ActualPos(), "(7, 4) EAST", "moves forward to EAST")
}

func TestBackwards(t *testing.T) {
	testCases := []struct {
		description string
		x, y        int
		direction   string
		want        string
	}{
		{"Move Backward from NORTH", 6, 4, "NORTH", "(6, 3) NORTH"},
		{"Move Backward from WEST", 6, 4, "WEST", "(7, 4) WEST"},
		{"Move Backward from SOUTH", 6, 4, "SOUTH", "(6, 5) SOUTH"},
		{"Move Backward from EAST", 6, 4, "EAST", "(5, 4) EAST"},
	}

	for _, tt := range testCases {
		t.Run(tt.description, func(t *testing.T) {
			c := NewRover(tt.x, tt.y, tt.direction)

			c.Receive("B")
			require.Equal(t, c.ActualPos(), tt.want, tt.description)
		})
	}
}

func TestRotateLeft(t *testing.T) {
	c := NewRover(6, 4, "EAST")

	c.Receive("L")
	require.Equal(t, c.ActualPos(), "(6, 4) NORTH", "rotate left to NORTH")
}

func TestRotateRight(t *testing.T) {
	c := NewRover(6, 4, "EAST")

	c.Receive("R")
	require.Equal(t, c.ActualPos(), "(6, 4) SOUTH", "rotate right to SOUTH")
}
