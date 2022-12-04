package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

var UPPER_BASE = int('A')
var LOWER_BASE = int('a')

func Intersection(a, b []rune) (c []rune) {
	m := make(map[rune]bool)

	for _, item := range a {
		m[item] = true
	}

	for _, item := range b {
		if _, ok := m[item]; ok {
			c = append(c, item)
		}
	}
	return
}

func readBadges(path string) (badges []rune, err error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var i int = 0
	var groupSack [][]rune
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var l = scanner.Text()
		if l != "" {
			groupSack = append(groupSack, []rune(l))
			i++
			if i%3 == 0 {
				badges = append(badges, Intersection(Intersection(groupSack[0], groupSack[1]), groupSack[2])[0])
				groupSack = nil
			}
		}
	}
	return
}

func main() {
	var ruckscaks = "input.txt"
	badges, _ := readBadges(ruckscaks)
	var total int
	for _, b := range badges {
		if unicode.IsUpper(b) == true {
			total += 27 + int(b) - UPPER_BASE
		} else {
			total += 1 + int(b) - LOWER_BASE
		}
	}
	fmt.Printf("adges values total: %d\n", total)
}
