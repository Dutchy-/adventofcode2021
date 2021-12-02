package main

import (
	"fmt"
	"regexp"
	"strings"
)

func lastyear_four() {
	lines := ReadLines("./2020/4")

	attrs := ParsePassportsAttrs(lines)
	valid := 0
	reallyValid := 0

	required := []string{"byr", "ecl", "eyr", "hcl", "hgt", "iyr", "pid"}

	for _, attr := range attrs {
		if attr.IsValid(required) {
			valid += 1
		}
		pass := attr.GetPassport()
		if pass.IsValid() {
			reallyValid += 1
			fmt.Println(pass)
		}
	}

	fmt.Printf("Valid passports: %d\n", valid)
	fmt.Printf("Really valid passports: %d\n", reallyValid)

}

type PassportAttrs map[string]string

type BirthYear struct {
	value int
}

func (i BirthYear) Init(v string) BirthYear {
	i.value = SafeNumber(v)
	return i
}

func (byr BirthYear) IsValid() bool {
	return byr.value > 1919 && byr.value < 2003
}

type IssueYear struct {
	value int
}

func (i IssueYear) Init(v string) IssueYear {
	i.value = SafeNumber(v)
	return i
}

func (iyr IssueYear) IsValid() bool {
	return iyr.value > 2009 && iyr.value < 2021
}

type ExpYear struct {
	value int
}

func (i ExpYear) Init(v string) ExpYear {
	i.value = SafeNumber(v)
	return i
}

func (eyr ExpYear) IsValid() bool {
	return eyr.value > 2019 && eyr.value < 2031
}

type Height struct {
	value int
	unit  string
}

func (hgt Height) Init(v string) Height {
	re := regexp.MustCompile(`(\d{2,3})(\w{2})`)
	m := re.FindStringSubmatch(v)
	if m != nil {
		hgt.value = Number(m[1])
		hgt.unit = m[2]
	}
	return hgt
}

func (hgt Height) IsValid() bool {
	if hgt.unit == "cm" {
		return hgt.value > 149 && hgt.value < 194
	} else if hgt.unit == "in" {
		return hgt.value > 58 && hgt.value < 77
	} else {
		return false
	}
}

type HairColor string

func (hcl HairColor) IsValid() bool {
	re := regexp.MustCompile(`#[0-9a-f]{6}`)
	return re.MatchString(string(hcl))
}

type EyeColor string

func (ecl EyeColor) IsValid() bool {
	re := regexp.MustCompile(`amb|blu|brn|gry|grn|hzl|oth`)
	return re.MatchString(string(ecl))
}

type PassportID string

func (pid PassportID) IsValid() bool {
	re := regexp.MustCompile(`^\d{9}$`)
	return re.MatchString(string(pid))
}

type Passport struct {
	// So... this should really be a list of PasswordProp
	// interfaces that implement IsValid() but this took me
	// way too long already so I'm going to leave it at this
	byr BirthYear
	iyr IssueYear
	eyr ExpYear
	hgt Height
	hcl HairColor
	ecl EyeColor
	pid PassportID
}

func (p *Passport) String() string {
	return fmt.Sprintf("%d %d %d %03d%s %s %s %s", p.byr.value, p.iyr.value, p.eyr.value, p.hgt.value, p.hgt.unit, p.hcl, p.ecl, p.pid)
}

func (p *Passport) IsValid() bool {
	return p.byr.IsValid() && p.iyr.IsValid() && p.eyr.IsValid() && p.hgt.IsValid() && p.hcl.IsValid() && p.ecl.IsValid() && p.pid.IsValid()
}

func (p *PassportAttrs) IsValid(required []string) bool {
	for _, req := range required {
		_, ok := (*p)[req]
		if !ok {
			return false
		}
	}
	return true
}

func (p PassportAttrs) GetPassport() *Passport {
	return &Passport{
		byr: BirthYear{}.Init(p["byr"]),
		iyr: IssueYear{}.Init(p["iyr"]),
		eyr: ExpYear{}.Init(p["eyr"]),
		hgt: Height{}.Init(p["hgt"]),
		hcl: HairColor(p["hcl"]),
		ecl: EyeColor(p["ecl"]),
		pid: PassportID(p["pid"]),
	}
}

func (p *PassportAttrs) GetKeys() []string {
	r := []string{}
	for k := range *p {
		r = append(r, k)
	}
	return r
}

func ParsePassportsAttrs(lines []string) []*PassportAttrs {
	result := []*PassportAttrs{}

	p := &PassportAttrs{}
	din := 0
	for _, line := range lines {
		if line == "" {
			(*p)["din"] = fmt.Sprintf("%d", din)
			result = append(result, p)
			p = &PassportAttrs{}
		} else {
			values := strings.Split(line, " ")
			for _, value := range values {
				attrs := strings.Split(value, ":")
				(*p)[attrs[0]] = attrs[1]
			}
		}
		din += 1
	}

	return result
}
