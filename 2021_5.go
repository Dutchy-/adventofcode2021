package main

import (
	"fmt"
	"regexp"
)

func thisyear_five() {
	l := ReadLines("./2021/5b")

	lines, maxX, maxY := ParseLines(l)
	grid := NewVentGrid(maxX, maxY)
	grid.AddStraightLines(lines)
	count := grid.CountHigherThan(1)
	fmt.Printf("%d overlapping straight vent lines\n", count)
	grid2 := NewVentGrid(maxX, maxY)
	grid2.AddLines(lines)
	// for _, row := range *grid2 {
	// 	fmt.Println(row)
	// }
	count2 := grid2.CountHigherThan(1)
	fmt.Printf("%d overlapping vent lines\n", count2)
}

func ParseLines(lines []string) ([]*Line, int, int) {
	maxX := 0
	maxY := 0
	result := []*Line{}
	for _, l := range lines {
		line := ParseLine(l)
		result = append(result, line)
		if line.x1 >= maxX {
			maxX = line.x1
		}
		if line.x2 >= maxX {
			maxX = line.x2
		}
		if line.y1 >= maxY {
			maxY = line.y1
		}
		if line.y2 >= maxY {
			maxY = line.y2
		}
	}
	return result, maxX + 1, maxY + 1
}

func ParseLine(l string) *Line {
	re := regexp.MustCompile(`(\d+),(\d+) -> (\d+),(\d+)`)
	m := re.FindStringSubmatch(l)
	line := Line{
		x1: Number(m[1]),
		y1: Number(m[2]),
		x2: Number(m[3]),
		y2: Number(m[4]),
	}
	return &line
}

type Line struct {
	x1 int
	y1 int
	x2 int
	y2 int
}

func (l *Line) IsStraight() bool {
	return l.x1 == l.x2 || l.y1 == l.y2
}

func (l *Line) AbsCoords() (int, int, int, int) {
	x1, x2 := l.x1, l.x2
	if l.x1 > l.x2 {
		x2, x1 = l.x1, l.x2
	}
	y1, y2 := l.y1, l.y2
	if l.y1 > l.y2 {
		y2, y1 = l.y1, l.y2
	}
	return x1, x2, y1, y2
}

type VentGrid [][]int

func NewVentGrid(xSize int, ySize int) *VentGrid {
	grid := make(VentGrid, ySize)
	for i := range grid {
		grid[i] = make([]int, xSize)
	}
	return &grid
}

func (grid *VentGrid) AddStraightLine(l Line) {
	x1, x2, y1, y2 := l.AbsCoords()
	for x := x1; x <= x2; x++ {
		for y := y1; y <= y2; y++ {
			(*grid)[y][x] += 1
		}
	}
}

func (grid *VentGrid) AddDiagonalLine(l Line) {
	xInc := 1
	if l.x1 > l.x2 {
		xInc = -1
	}
	yInc := 1
	if l.y1 > l.y2 {
		yInc = -1
	}
	for x, y := l.x1, l.y1; x != l.x2+xInc && y != l.y2+yInc; {
		(*grid)[y][x] += 1
		x += xInc
		y += yInc
	}
}

func (grid *VentGrid) AddLines(lines []*Line) {
	for _, l := range lines {
		if l.IsStraight() {
			grid.AddStraightLine(*l)
		} else {
			grid.AddDiagonalLine(*l)
		}
	}
}

func (grid *VentGrid) AddStraightLines(lines []*Line) {
	for _, l := range lines {
		if l.IsStraight() {
			grid.AddStraightLine(*l)
		}
	}
}

func (grid *VentGrid) CountHigherThan(x int) int {
	count := 0
	for _, row := range *grid {
		for _, value := range row {
			if value > x {
				count += 1
			}
		}
	}
	return count
}
