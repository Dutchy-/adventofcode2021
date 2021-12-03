package main

import (
	"fmt"
)

func thisyear_three() {
	lines := ReadLines("./2021/3b")
	n := NumbersBase(lines, 2)
	bits := len(lines[0])

	gamma, epsilon := GammaEpsilon(n, bits)
	fmt.Printf("Gamma is %d and the inverse (epsilon) is %d\n", gamma, epsilon)
	fmt.Printf("Multiplied they are %d\n", gamma*epsilon)

	o := FindOxygen(n, bits, gamma)
	c := FindCO2(n, bits, epsilon)
	fmt.Printf("Oxygen is %d and CO2 is %d\n", o, c)
	fmt.Printf("Multiplied they are %d\n", o*c)
}

func Gamma(numbers []int, bits int) int {
	l := len(numbers)
	result := 0
	for bits > 0 {
		count := 0
		for _, value := range numbers {
			count += Bit(value, bits)
		}
		add := 0
		if count*2 >= l {
			add = 1
		}
		result += add << (bits - 1)
		bits -= 1
	}
	return result
}

func Epsilon(gamma int, bits int) int {
	return gamma ^ BitMask(bits)
}

func GammaEpsilon(numbers []int, bits int) (int, int) {
	gamma := Gamma(numbers, bits)
	epsilon := Epsilon(gamma, bits)
	return gamma, epsilon
}

func ReduceCommon(numbers []int, pos int, gamma int) []int {
	result := []int{}
	gv := Bit(gamma, pos)
	for _, n := range numbers {
		if gv == Bit(n, pos) {
			result = append(result, n)
		}
	}
	return result
}

func FindSelector(numbers []int, bits int, selector func(int, int) int) int {
	// We do not need to calculate the whole gamma in this function,
	// but it's easier to reuse the code
	s := selector(GammaEpsilon(numbers, bits))
	for len(numbers) > 1 {
		numbers = ReduceCommon(numbers, bits, s)
		s = selector(GammaEpsilon(numbers, bits))
		bits -= 1
	}
	return numbers[0]
}

func FindOxygen(numbers []int, bits int, gamma int) int {
	return FindSelector(numbers, bits, func(gamma int, epsilon int) int {
		return gamma
	})
}

func FindCO2(numbers []int, bits int, epsilon int) int {
	return FindSelector(numbers, bits, func(gamma int, epsilon int) int {
		return epsilon
	})
}
