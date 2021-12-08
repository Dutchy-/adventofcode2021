package main

import (
	"fmt"
	"math"
	"strings"
)

func thisyear_eight() {
	l := ReadLines("./2021/8b")

	ds := ParseDisplays(l)
	fmt.Printf("The number of known 7-segment digits is %d\n", CountKnownDigits(ds))
	r := GetDisplaySum(ds)
	fmt.Printf("The sum of all 7-segment displays is %d\n", r)

}

func Signals(v []string) []Signal {
	signals := []Signal{}
	for _, s := range v {
		signals = append(signals, Signal(s))
	}
	return signals
}

func ParseDisplays(l []string) []*Display {
	r := []*Display{}
	for _, line := range l {
		v := strings.Split(line, " | ")
		d := Display{
			Map:     map[int]Signal{},
			Hashes:  map[int]int{},
			Signals: Signals(strings.Split(v[0], " ")),
			Outputs: Signals(strings.Split(v[1], " ")),
		}
		d.Decode()
		r = append(r, &d)
	}
	return r
}

type Signal string

type Display struct {
	Signals []Signal
	Outputs []Signal
	Map     map[int]Signal
	Hashes  map[int]int
}

func (d *Display) FindKnownDigits() {
	for _, signal := range d.Signals {
		switch len(signal) {
		case 2:
			d.Map[1] = signal
		case 3:
			d.Map[7] = signal
		case 4:
			d.Map[4] = signal
		case 7:
			d.Map[8] = signal
		}
	}
}

func GetDisplaySum(ds []*Display) int {
	r := 0
	for _, d := range ds {
		r += d.GetOutput()
	}
	return r
}

func CountKnownDigits(ds []*Display) int {
	count := 0
	for _, d := range ds {
		for _, o := range d.Outputs {
			switch d.Hashes[o.Hash()] {
			case 1, 4, 7, 8:
				count += 1
			}
		}
	}
	return count
}

func (d *Display) CalculateRemainingDigits() {
	for _, signal := range d.Signals {
		o8 := signal.CountOverlap(d.Map[8])
		o4 := signal.CountOverlap(d.Map[4])
		o7 := signal.CountOverlap(d.Map[7])

		// We select on length so we can skip the ones we
		// already know (1, 4, 7, 8)
		switch len(signal) {
		// 0, 6 and 9 have 6 segments
		case 6:
			if o8 == 6 && o4 == 3 && o7 == 3 {
				d.Map[0] = signal
			} else if o8 == 6 && o4 == 3 && o7 == 2 {
				d.Map[6] = signal
			} else {
				d.Map[9] = signal
			}
		// 2, 3 and 5 have 5 segments
		case 5:
			if o8 == 5 && o4 == 2 {
				d.Map[2] = signal
			} else if o8 == 5 && o4 == 3 && o7 == 3 {
				d.Map[3] = signal
			} else {
				d.Map[5] = signal
			}
		}
	}
}

func (d *Display) Hash() {
	for value, s := range d.Map {
		d.Hashes[s.Hash()] = value
	}
}

func (d *Display) GetOutput() int {
	r := 0
	l := len(d.Outputs)
	for i, o := range d.Outputs {
		pos := float64(l - i - 1)
		r += int(d.Hashes[o.Hash()] * int(math.Pow(10, pos)))
	}
	return r
}

func (d *Display) Decode() {
	d.FindKnownDigits()
	d.CalculateRemainingDigits()
	d.Hash()
}

func (s Signal) CountOverlap(s2 Signal) int {
	c := 0
	for _, r := range s {
		if strings.ContainsRune(string(s2), r) {
			c += 1
		}
	}
	return c
}

func (s Signal) Hash() int {
	h := 0
	for _, r := range s {
		offset := r - 'a'
		h += 1 << offset
	}
	return h
}
