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

func SafeNumber(line string) int {
	value, err := strconv.Atoi(line)
	if err != nil {
		return 0
	}
	return value
}

func Number(line string) int {
	return NumberBase(line, 10)
}

func Numbers(lines []string) []int {
	return NumbersBase(lines, 10)
}

func NumberBase(line string, base int) int {
	value, err := strconv.ParseInt(line, base, 32)
	check(err)
	return int(value)
}

func NumbersBase(lines []string, base int) []int {
	result := make([]int, 0)

	for _, line := range lines {
		result = append(result, NumberBase(line, base))
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

func BitMask(bits int) int {
	return 1<<bits - 1
}

func Bit(n int, pos int) int {
	return n >> (pos - 1) & 1
}

func Mean(n []int) float64 {
	return float64(sum(n)) / float64(len(n))
}

func Median(n []int) int {
	l := len(n)
	if l%2 == 0 {
		return int(Mean(n[l/2-1 : l/2+1]))
	} else {
		return n[l/2]
	}
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
