package main

import (
	"fmt"
	"sort"
)

func thisyear_nine() {
	l := ReadLines("./2021/9b")

	cave := ParseCave(l)
	// fmt.Println(cave)
	fmt.Printf("Number of lowest points in the smoke cave: %d\n", cave.CountRiskLows())
	basins := cave.FindBasins()
	sort.Slice(basins, func(i, j int) bool {
		return len(basins[i]) > len(basins[j])
	})
	answer := 1
	for _, b := range basins[:3] {
		answer *= len(b)
	}
	fmt.Printf("The multiplied size of the smoke basins is: %d\n", answer)
}

type SmokePoint struct {
	Height      int
	Up          *SmokePoint
	Down        *SmokePoint
	Left        *SmokePoint
	Right       *SmokePoint
	PartOfBasin bool
}

type SmokeCave [][]*SmokePoint

func ParseCave(l []string) SmokeCave {
	cave := make(SmokeCave, len(l))
	for i, line := range l {
		// We make a row
		width := len(line)
		cave[i] = make([]*SmokePoint, width)

		// We loop over the input row
		for j, square := range line {
			// i, j are coordinates
			pos := &SmokePoint{Height: Number(string(square))}
			var up *SmokePoint
			// connect to the line above
			if i != 0 {
				up = cave[i-1][j]
				pos.Up = up
				up.Down = pos
			}
			var left *SmokePoint
			// connect to the previous position
			if j != 0 {
				left = cave[i][j-1]
				pos.Left = left
				left.Right = pos
			}

			cave[i][j] = pos
		}
	}

	return cave
}

func (s *SmokePoint) GetNeighbours() []*SmokePoint {
	return []*SmokePoint{s.Up, s.Down, s.Left, s.Right}
}
func (s *SmokePoint) IsLow() bool {
	result := true
	for _, nb := range s.GetNeighbours() {
		result = result && (nb == nil || s.Height < nb.Height)
	}
	return result
}

func (s *SmokePoint) RiskLevel() int {
	return s.Height + 1
}

func (b SmokeCave) String() string {
	result := ""
	for _, row := range b {
		for _, p := range row {
			result += fmt.Sprint(p.Height)
		}
		result += fmt.Sprintln()
	}
	return result
}

func (b SmokeCave) GetLows() []*SmokePoint {
	result := []*SmokePoint{}
	for _, row := range b {
		for _, p := range row {
			if p.IsLow() {
				result = append(result, p)
			}
		}
	}
	return result
}

func (b SmokeCave) CountRiskLows() int {
	result := 0
	for _, lp := range b.GetLows() {
		result += lp.RiskLevel()
	}
	return result
}

func (b *SmokeCave) Reset() {
	for _, row := range *b {
		for _, p := range row {
			p.PartOfBasin = false
		}
	}
}

type SmokeBasin []*SmokePoint

func (p *SmokePoint) Walk() SmokeBasin {
	p.PartOfBasin = true
	basin := SmokeBasin{p}
	for _, nb := range p.GetNeighbours() {
		if nb != nil && nb.Height > p.Height && nb.Height != 9 && !nb.PartOfBasin {
			basin = append(basin, nb.Walk()...)
		}
	}
	return basin
}

func (b *SmokeCave) FindBasins() []SmokeBasin {
	basins := []SmokeBasin{}
	for _, lp := range b.GetLows() {
		basins = append(basins, lp.Walk())
		b.Reset()
	}
	return basins
}
