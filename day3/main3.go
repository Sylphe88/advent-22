package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

var UPPER_BASE = int('A')
var LOWER_BASE = int('a')

func Intersection(a, b []rune) (c rune) {
	m := make(map[rune]bool)

	for _, item := range a {
		m[item] = true
	}

	for _, item := range b {
		if _, ok := m[item]; ok {
			c = item
			break
		}
	}
	return
}

func readCompartmentDuplicate(path string) ([]rune, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var duplicates []rune
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var l = scanner.Text()
		if l != "" {
			liner := []rune(l)
			duplicates = append(duplicates, Intersection(liner[0:len(l)/2], liner[len(l)/2:len(l)]))
		}
	}
	return duplicates, scanner.Err()
}

func main() {
	var ruckscaks = "input.txt"
	dups, _ := readCompartmentDuplicate(ruckscaks)
	var total int
	for _, d := range dups {
		if unicode.IsUpper(d) == true {
			total += 27 + int(d) - UPPER_BASE
		} else {
			total += 1 + int(d) - LOWER_BASE
		}
	}
	fmt.Printf("Duplicated snacks total: %d\n", total)
}
