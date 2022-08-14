package main

import (
	"fmt"
)

type rover struct {
	x int
	y int

	direction string
}

var directions = map[string]bool{
	"NORTH": true,
	"WEST":  true,
	"SOUTH": true,
	"EAST":  true,
}

func validDirection(s string) bool {
	_, exists := directions[s]

	return exists
}

func NewRover(x, y int, direction string) *rover {
	if !validDirection(direction) {
		return nil
	}

	return &rover{
		x:         x,
		y:         y,
		direction: direction,
	}
}

// ActualPos computes actual position.
func (r *rover) ActualPos() string {
	return fmt.Sprintf("(%d, %d) %s", r.x, r.y, r.direction)
}

// Receive receives commands
func (r *rover) Receive(c string) {
	if len(c) == 0 {
		return
	}

	for i := 0; i < len(c); i++ {
		action := c[i : i+1]

		switch action {
		case "F":
			{
				r.MoveForward()
			}

		case "B":
			{
				r.MoveBackward()
			}

		case "L":
			{
				r.RotateLeft()
			}

		case "R":
			{
				r.RotateRight()
			}
		}
	}
}

// MoveForward as per current direction.
func (r *rover) MoveForward() {
	if r.direction == "NORTH" {
		r.y++
	}

	if r.direction == "WEST" {
		r.x--
	}

	if r.direction == "SOUTH" {
		r.y--
	}

	if r.direction == "EAST" {
		r.x++
	}
}

// MoveBackward as per current direction.
func (r *rover) MoveBackward() {
	if r.direction == "NORTH" {
		r.y--
	}

	if r.direction == "WEST" {
		r.x++
	}

	if r.direction == "SOUTH" {
		r.y++
	}

	if r.direction == "EAST" {
		r.x--
	}
}

// RotateLeft as per current position.
func (r *rover) RotateLeft() {
	if r.direction == "NORTH" {
		r.direction = "WEST"
	}

	if r.direction == "WEST" {
		r.direction = "SOUTH"
	}

	if r.direction == "SOUTH" {
		r.direction = "EAST"
	}

	if r.direction == "EAST" {
		r.direction = "NORTH"
	}
}

// RotateRight as per current position.
func (r *rover) RotateRight() {
	if r.direction == "NORTH" {
		r.direction = "EAST"
	}

	if r.direction == "WEST" {
		r.direction = "NORTH"
	}

	if r.direction == "SOUTH" {
		r.direction = "WEST"
	}

	if r.direction == "EAST" {
		r.direction = "SOUTH"
	}
}
