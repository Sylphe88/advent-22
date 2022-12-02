package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

var opponentPlay = map[string]int{"A": 0b001, "B": 0b010, "C": 0b100}
var myTargetScore = map[string]int{"X": 0, "Y": 3, "Z": 6}

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
			var matchVal int
			switch plays[1] {
			case "X":
				// Circular right shift
				matchVal = ((opponentPlay[plays[0]] >> 1) | (opponentPlay[plays[0]] << 2)) % 8
				break
			case "Y":
				matchVal = opponentPlay[plays[0]]
				break
			case "Z":
				// Circular left shift
				matchVal = ((opponentPlay[plays[0]] << 1) | (opponentPlay[plays[0]] >> 2)) % 8
			}
			myPlayScore := 1 + int(math.Log2(float64(matchVal))) // Get the XYZ value by checking the '1' position in the value
			score += myPlayScore + myTargetScore[plays[1]]
			fmt.Printf("Current score: %d\n", score)
		}
	}
	return score, err
}

func main() {
	score, _ := readCheatSheet("input.txt")
	fmt.Printf("Actual cheat score: %d\n", score)
}
