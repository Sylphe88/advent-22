package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var opponentPlay = map[string]int{"A": 0, "B": 1, "C": 2}
var myPlay = map[string]int{"X": 0, "Y": 1, "Z": 2}
var playValue = map[string]int{"X": 1, "Y": 2, "Z": 3}
var matchScore = map[int]int{-2: 6, -1: 0, 0: 3, 1: 6, 2: 0}

func readCheatSheet(path string) (int, error) {
	file, err := os.Open(path)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	var score int = 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var l = scanner.Text()
		if l != "" {
			plays := strings.Fields(l)
			score += playValue[plays[1]] + matchScore[myPlay[plays[1]]-opponentPlay[plays[0]]]
			fmt.Printf("Current score: %d\n", score)
		}
	}
	return score, err
}

func main() {
	score, _ := readCheatSheet("input.txt")
	fmt.Printf("Cheat score: %d\n", score)
}
