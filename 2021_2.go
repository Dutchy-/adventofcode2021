package main

import (
	"fmt"
)

func thisyear_two() {
	lines := ReadLines("./2021/2")

	dirs := ParseDirectionsSum(lines)

	depth := dirs["down"] - dirs["up"]
	answer := dirs["forward"] * depth
	fmt.Printf("The answer for part 1 is %d\n", answer)

	pos := ParseDirections(lines)
	answer2 := pos.Distance * pos.Depth
	fmt.Printf("The answer for part 2 is %d\n", answer2)
}

func ParseDirectionsSum(lines []string) map[string]int {
	result := map[string]int{}
	for _, line := range lines {
		dir, dist := ParseInstruction(line)
		result[dir] += dist
	}

	return result
}

func ParseDirections(lines []string) *SubmarinePosition {
	pos := SubmarinePosition{}
	for _, line := range lines {
		dir, dist := ParseInstruction(line)
		pos.Execute(dir, dist)
	}
	return &pos
}

type SubmarinePosition struct {
	Distance int
	Depth    int
	aim      int
}

func (pos *SubmarinePosition) AdjustAim(value int) {
	pos.aim += value
}

func (pos *SubmarinePosition) Forward(value int) {
	pos.Distance += value
	pos.Depth += pos.aim * value
}

func (pos *SubmarinePosition) Execute(dir string, value int) {
	switch dir {
	case "forward":
		pos.Forward(value)
	case "down":
		pos.AdjustAim(value)
	case "up":
		pos.AdjustAim(-value)
	}
}
