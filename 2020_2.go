package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

/// TWO ///

type Policy struct {
	min    int
	max    int
	letter string
}

func (policy *Policy) permittedCount(password string) bool {
	count := strings.Count(password, policy.letter)
	return count >= policy.min && count <= policy.max
}

func (policy *Policy) permittedPosition(password string) bool {
	result := 0
	if string(password[policy.min-1]) == policy.letter {
		result += 1
	}
	if string(password[policy.max-1]) == policy.letter {
		result += 1
	}
	return result == 1
}

func lastyear_two() {
	values := ReadLines("./2020/2")

	validPasswords_part1 := 0
	validPasswords_part2 := 0

	for _, value := range values {
		policy, password := ParsePasswordPolicy(value)
		if policy.permittedCount(password) {
			validPasswords_part1 += 1
		}
		if policy.permittedPosition(password) {
			validPasswords_part2 += 1
		}
	}

	fmt.Printf("Valid passwords part 1: %d\n", validPasswords_part1)
	fmt.Printf("Valid passwords part 2: %d\n", validPasswords_part2)
}

func ParsePasswordPolicy(line string) (Policy, string) {
	re := regexp.MustCompile(`(\d+)-(\d+) (\w{1}): (\w+)`)
	matches := re.FindStringSubmatch(line)
	min, err := strconv.Atoi(matches[1])
	check(err)
	max, err := strconv.Atoi(matches[2])
	check(err)
	return Policy{min, max, matches[3]}, matches[4]
}
