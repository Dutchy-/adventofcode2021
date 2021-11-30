package main

import "fmt"

/// ONE ///

func lastyear_one() {
	values := Numbers(ReadLines("./2020/1b"))
	// fmt.Println(values)

	lastyear_one_part1(values)
	lastyear_one_part2(values)

}

func lastyear_one_part1(values []int) {
	next := nextProduct(values, 2)
	for {
		test := next()
		if isAnswer(test, sum, 2020) {
			fmt.Printf("The answer is %d\n", mul(test))
			break
		}
	}
}
func lastyear_one_part2(values []int) {
	next2 := nextProduct(values, 3)
	for {
		test := next2()
		if isAnswer(test, sum, 2020) {
			fmt.Printf("The answer is %d\n", mul(test))
			break
		}
	}
}

func isAnswer(product []int, f func(a []int) int, value int) bool {
	return f(product) == value
}
