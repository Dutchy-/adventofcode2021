package main

import (
	"fmt"
	"regexp"
	"strings"
)

func thisyear_four() {
	l := ReadLines("./2021/4")

	n := Numbers(strings.Split(l[0], ","))

	cards := ParseCards(l[2:])
	winningNr, card := Play(cards, n)
	unmarked := card.GetUnmarked()
	fmt.Printf("Answer for bingo card is %d\n", winningNr*sum(unmarked))

	for _, c := range cards {
		c.Reset()
	}

	lastWinningNr, lastCard := FindLastWinning(cards, n)
	lastUnmarked := lastCard.GetUnmarked()

	fmt.Printf("Answer for last winning bingo card is %d\n", lastWinningNr*sum(lastUnmarked))

}

type Field struct {
	Value  int
	Marked bool
}

type Card [][]*Field

func (c *Card) Mark(n int) {
	for _, row := range *c {
		for _, field := range row {
			if field.Value == n {
				field.Marked = true
			}
		}
	}
}

func (c *Card) IsWinner() bool {
	cardsize := len(*c)
	// horizontal
	for i := 0; i < cardsize; i++ {
		result := 0
		for j := 0; j < cardsize; j++ {
			// fmt.Printf("i: %d, j: %d, card: %v\n", i, j, (*c))
			if (*c)[i][j].Marked {
				result += 1
			}
		}
		// fmt.Println(result)
		if result == cardsize {
			return true
		}
	}

	// vertical
	for j := 0; j < cardsize; j++ {
		result := 0
		for i := 0; i < cardsize; i++ {
			if (*c)[i][j].Marked {
				result += 1
			}
		}
		if result == cardsize {
			return true
		}
	}

	// no winner
	return false
}

func (c *Card) GetUnmarked() []int {
	result := []int{}
	for _, row := range *c {
		for _, field := range row {
			if !field.Marked {
				result = append(result, field.Value)
			}
		}
	}
	return result
}

func (c *Card) Reset() {
	for _, row := range *c {
		for _, field := range row {
			field.Marked = false
		}
	}
}

func NewCard(lines []string) *Card {
	cardsize := 5
	re := regexp.MustCompile(`\s+`)
	c := make(Card, cardsize)
	for i, l := range lines {
		c[i] = make([]*Field, cardsize)
		values := re.Split(strings.Trim(l, " "), -1)
		for j, value := range values {
			c[i][j] = &Field{Value: Number(value)}
		}
	}
	return &c
}

func ParseCards(lines []string) []*Card {
	cards := []*Card{}
	i := 0
	for i < len(lines) {
		c := NewCard(lines[i : i+5])
		i += 6
		cards = append(cards, c)
	}
	return cards
}

func Play(cards []*Card, numbers []int) (int, *Card) {
	for _, n := range numbers {
		for _, c := range cards {
			c.Mark(n)
			if c.IsWinner() {
				return n, c
			}
		}
	}
	return 0, nil
}

func FindLastWinning(cards []*Card, numbers []int) (int, *Card) {
	var lastWinner *Card = nil
	lastN := -1
	lastIndex := -1
	for _, c := range cards {
		for i, n := range numbers {
			c.Mark(n)
			if c.IsWinner() {
				if i > lastIndex {
					lastWinner = c
					lastN = n
					lastIndex = i
				}
				break
			}
		}
	}
	return lastN, lastWinner
}
