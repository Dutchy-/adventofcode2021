package main

import (
	"fmt"
	"strings"
)

/// THREE ///

func lastyear_three() {
	lines := ReadLines("./2020/3")

	grid := ParseGrid(lines)

	paths := []Path{
		{Right, Down},
		{Right, Right, Right, Down},
		{Right, Right, Right, Right, Right, Down},
		{Right, Right, Right, Right, Right, Right, Right, Down},
		{Right, Down, Down},
	}
	results := WalkPaths(grid, &paths)
	fmt.Printf("Trees multiplied: %d\n", mul(results))
}

func ParseGrid(lines []string) *Grid {
	// We make a grid
	grid := make(Grid, len(lines))

	// We loop over the input
	for i, line := range lines {
		// We make a row
		width := len(line)
		grid[i] = make([]*Position, width)

		// We loop over the input row
		for j, square := range line {
			// i, j are coordinates
			pos := &Position{Tree: square == '#'}
			var up *Position
			// connect to the line above
			if i != 0 {
				up = grid[i-1][j]
				pos.Up = up
				up.Down = pos
			}
			var left *Position
			// connect to the previous position
			if j != 0 {
				left = grid[i][j-1]
				pos.Left = left
				left.Right = pos
			}
			var right *Position
			// loop around at the end of the line
			if j == width-1 {
				right = grid[i][0]
				pos.Right = right
				right.Left = pos
			}

			grid[i][j] = pos
		}
	}

	return &grid
}

func WalkPaths(grid *Grid, paths *[]Path) []int {
	result := make([]int, 0)
	for i, path := range *paths {
		trees := grid.Walk(path)
		fmt.Printf("Encountered %d trees on path %d\n", trees, i)
		result = append(result, trees)
	}
	return result
}

/// Position

type Position struct {
	Up    *Position
	Down  *Position
	Left  *Position
	Right *Position
	Tree  bool
}

func (pos Position) String() string {
	if pos.Tree {
		return "#"
	}
	return "."
}

type Step func(*Position) *Position

type Path []Step

func Down(pos *Position) *Position {
	return pos.Down
}

func Right(pos *Position) *Position {
	return pos.Right
}

/// Grid

type Grid [][]*Position

func (grid *Grid) String() string {
	sb := strings.Builder{}
	for _, row := range *grid {
		for _, pos := range row {
			sb.WriteString(pos.String())
		}
		sb.WriteString("\n")
	}
	return sb.String()
}

func (grid Grid) Walk(path Path) int {
	pos := grid[0][0]
	result := 0
	for pos.Down != nil {
		for _, step := range path {
			pos = step(pos)
		}
		if pos.Tree {
			result += 1
		}
	}
	return result
}
