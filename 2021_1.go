package main

import "fmt"

func thisyear_one() {
	numbers := Numbers(ReadLines("2021/1"))

	increases := CountCompare(numbers, 1, IsLarger)
	fmt.Printf("There are %d increases for sliding window 1\n", increases)
	increases3 := CountCompare(numbers, 3, IsLarger)
	fmt.Printf("There are %d increases for sliding window 3\n", increases3)
}

func CountCompare(numbers []int, windowSize int, compare CompareFunc) int {
	result := 0
	i := windowSize
	for i < len(numbers) {
		window := numbers[i-windowSize+1 : i+1]
		prevWindow := numbers[i-windowSize : i]
		if compare(sum(window), sum(prevWindow)) {
			result += 1
		}
		i += 1
	}
	return result
}

type CompareFunc func(int, int) bool

func IsLarger(a int, b int) bool {
	return a > b
}
