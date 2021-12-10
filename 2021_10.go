package main

import (
	"fmt"
	"regexp"
	"sort"

	"github.com/alecthomas/participle/v2"
)

func thisyear_ten() {
	l := ReadLines("./2021/10b")

	scoreCorrupt, scoreComplete := ParseSyntax(l)
	fmt.Printf("The score of corrupt lines is %d\n", scoreCorrupt)
	sort.Ints(scoreComplete)
	middleScore := scoreComplete[len(scoreComplete)/2]
	fmt.Printf("The score of the middle autocompleted line is %d\n", middleScore)
}

var corruptMap map[string]int = map[string]int{")": 3, "]": 57, "}": 1197, ">": 25137}
var opposite map[string]string = map[string]string{"(": ")", "{": "}", "[": "]", "<": ">"}
var completeMap map[string]int = map[string]int{")": 1, "]": 2, "}": 3, ">": 4}

func ParseSyntax(l []string) (int, []int) {
	parser, err := participle.Build(&ChunkList{})
	check(err)

	scoreCorrupt := 0
	completeScores := []int{}

	for _, line := range l {
		err = parser.ParseString("", line, &ChunkList{})
		unexp, _ := ParseErrorMissingChar(err)
		if IsClose(unexp) {
			scoreCorrupt += corruptMap[unexp]
		} else {
			add := NaiveComplete(line)
			err = parser.ParseString("", line+add, &ChunkList{})
			if err != nil {
				add += Autocomplete(parser, line+add, err)
			}
			completeScores = append(completeScores, CalcScore(add))
		}
	}
	return scoreCorrupt, completeScores
}

func CalcScore(add string) int {
	score := 0
	for _, c := range add {
		score *= 5
		score += completeMap[string(c)]
	}
	return score
}

func NaiveComplete(line string) string {
	add := ""
	final := string(line[len(line)-1])
	for i := 2; IsOpen(final); i++ {
		add += opposite[final]
		final = string(line[len(line)-i])
	}
	return add
}

func Autocomplete(parser *participle.Parser, line string, err error) string {
	add := ""
	for ; err != nil; err = parser.ParseString("", line+add, &ChunkList{}) {
		_, exp := ParseErrorMissingChar(err)
		add += exp
	}
	return add
}

func IsClose(unexp string) bool {
	return unexp == "}" || unexp == "]" || unexp == ")" || unexp == ">"
}

func IsOpen(unexp string) bool {
	return unexp == "{" || unexp == "[" || unexp == "(" || unexp == "<"
}

var re *regexp.Regexp = regexp.MustCompile(`unexpected token "(.*?)"( \(expected "(.*?)"\))?`)

func ParseErrorMissingChar(err error) (string, string) {
	m := re.FindStringSubmatch(err.Error())
	return m[1], m[3]
}

type ChunkList struct {
	Chunks []*Chunk `parser:"@@*"`
}

type Chunk struct {
	Round  *ChunkList `parser:"('(' @@? ')')*"`
	Square *ChunkList `parser:"('[' @@? ']')*"`
	Curly  *ChunkList `parser:"('{' @@? '}')*"`
	Hook   *ChunkList `parser:"('<' @@? '>')*"`
}
