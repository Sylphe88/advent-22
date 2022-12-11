package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type Section struct {
	Start int
	End   int
}

func (s Section) getRange() int {
	return s.End - s.Start + 1
}

func isInnerSection(s Section, s2 Section) bool {
	return (s.Start <= s2.Start && s.End >= s2.End) || (s2.Start <= s.Start && s2.End >= s.End)
}

func isOverlappedPair(s Section, s2 Section) bool {
	return (s.End >= s2.Start && s.Start <= s2.Start) || (s2.End >= s.Start && s2.Start <= s.Start)
}

func getDuplicatePairs(path string) (int, int, error) {
	file, err := os.Open(path)
	if err != nil {
		return 0, 0, err
	}
	defer file.Close()

	duplicates, overlaps := 0, 0
	var pair [2]Section
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var l = scanner.Text()
		r := regexp.MustCompile(`(\d+)-(\d+),(\d+)-(\d+)`)
		matches := r.FindAllStringSubmatch(l, -1)
		if len(matches) > 0 {
			var sections [4]int
			for i := 1; i < len(matches[0]); i++ {
				sections[i-1], err = strconv.Atoi(matches[0][i])
			}
			pair[0] = Section{
				Start: sections[0],
				End:   sections[1],
			}
			pair[1] = Section{
				Start: sections[2],
				End:   sections[3],
			}
			if isInnerSection(pair[0], pair[1]) {
				duplicates++
			}
			if isOverlappedPair(pair[0], pair[1]) {
				overlaps++
			}
		}

	}
	return duplicates, overlaps, err
}

func main() {
	var pairsfile = "input.txt"
	dups, overlaps, _ := getDuplicatePairs(pairsfile)
	fmt.Printf("Total duplicate pairs: %d\n", dups)
	fmt.Printf("Total overlapping pairs: %d\n", overlaps)
}
