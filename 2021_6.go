package main

import (
	"fmt"
	"strings"
)

func thisyear_six() {
	state := Numbers(strings.Split(ReadLines("./2021/6")[0], ","))

	s := School{}.Init(state)
	s.Age(80)
	fmt.Printf("Lanternfish school size after 80 days: %d\n", s.GetTotal())
	s.Age(256 - 80)
	fmt.Printf("Lanternfish school size after 256 days: %d\n", s.GetTotal())
}

type School struct {
	Count map[int]int
}

func (s School) Init(state []int) School {
	s.Count = map[int]int{}
	for _, v := range state {
		s.Count[v] += 1
	}
	return s
}

func (s *School) Age(days int) {
	for d := 0; d < days; d++ {
		newCount := map[int]int{}
		for age := 8; age > 0; age-- {
			newCount[age-1] = s.Count[age]
		}
		newCount[8] = s.Count[0]
		newCount[6] += s.Count[0]
		s.Count = newCount
	}
}

func (s *School) GetTotal() int {
	r := 0
	for _, v := range s.Count {
		r += v
	}
	return r
}
