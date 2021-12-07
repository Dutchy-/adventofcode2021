package main

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

func thisyear_seven() {
	n := Numbers(strings.Split(ReadLines("./2021/7b")[0], ","))

	sort.Ints(n)
	// fmt.Println(n)
	median := Median(n)
	d := SumDistances(n, median, DistanceLinear)

	fmt.Printf("The median of the crab positions is %d, the distance (fuel) to move there is %d\n", median, d)
	meanF := math.Round(Mean(n))
	mean := int(meanF)
	// So.. the calculation below works for the demo data set but not for the actual data set
	d2 := SumDistances(n, mean, DistanceExp)
	fmt.Printf("The mean of the crab positions is %d, the distance (fuel) to move there is %d\n", mean, d2)
	// For the actual dataset the mean-1 (461) is the answer, even though the non-rounded mean is closer to 462
	// than to 461 (although barely). I haven't figured out yet why the mean is not the answer here, but it
	// may have to do with the distribution of the points in combination with the cost function.
	d3 := SumDistances(n, mean-1, DistanceExp)
	fmt.Printf("The mean-1 of the crab positions is %d, the distance (fuel) to move there is %d\n", mean-1, d3)
	// fmt.Println(mean)
}

type DistanceFunc func(int, int) int

func SumDistances(n []int, target int, f DistanceFunc) int {
	r := 0
	for _, v := range n {
		r += f(v, target)
	}
	return r
}

func DistanceExp(n int, target int) int {
	return Exp(DistanceLinear(n, target))
}

func DistanceLinear(n int, target int) int {
	return Abs(target - n)
}

func Exp(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return n + Exp(n-1)
}
