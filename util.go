package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func ReadLines(path string) []string {
	result := make([]string, 0)

	file, err := os.Open(path)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}

	check(scanner.Err())
	return result
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func Numbers(lines []string) []int {
	result := make([]int, 0)

	for _, line := range lines {
		value, err := strconv.Atoi(line)
		check(err)
		result = append(result, value)
	}

	return result
}

func nextProduct(a []int, r int) func() []int {
	p := make([]int, r)
	x := make([]int, len(p))
	return func() []int {
		p := p[:len(x)]
		for i, xi := range x {
			p[i] = a[xi]
		}
		for i := len(x) - 1; i >= 0; i-- {
			x[i]++
			if x[i] < len(a) {
				break
			}
			x[i] = 0
			if i <= 0 {
				x = x[0:0]
				break
			}
		}
		return p
	}
}

func sum(a []int) int {
	result := 0
	for _, v := range a {
		result += v
	}
	return result
}

func mul(a []int) int {
	result := 1
	for _, v := range a {
		result *= v
	}
	return result
}

func ParseInstruction(line string) (string, int) {
	x := strings.Split(line, " ")
	i, err := strconv.Atoi(x[1])
	check(err)
	return x[0], i
}
