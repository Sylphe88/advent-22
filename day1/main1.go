package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"

	"github.com/sylphe88/advent22/elf"
)

func readElves(path string) ([]elf.Elf, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var elves []elf.Elf

	scanner := bufio.NewScanner(file)
	e := elf.Elf{
		Id: 0,
	}
	for scanner.Scan() {
		var l = scanner.Text()
		if l != "" {
			snack, _ := strconv.Atoi(l)
			e.Calories = append(e.Calories, snack)
			continue
		} else {
			newElf := e
			elves = append(elves, newElf)
			e = elf.Elf{
				Id: newElf.Id + 1,
			}
		}

	}
	return elves, scanner.Err()
}

func main() {
	var calories_files = "input.txt"
	elves, _ := readElves(calories_files)
	fattestElf := elf.Elf{}
	for _, e := range elves {
		fmt.Printf("Elf %d: %d calories\n", e.Id, e.GetTotalCalories())
		if e.GetTotalCalories() > fattestElf.GetTotalCalories() {
			fattestElf = e
		}
	}
	fmt.Printf("The fatest elf is #%d (%d calories!)\n", fattestElf.Id, fattestElf.GetTotalCalories())

	var eSlice []elf.Elf = elves[0:len(elves)]
	sort.Slice(eSlice, func(i, j int) bool {
		return eSlice[i].GetTotalCalories() > eSlice[j].GetTotalCalories()
	})
	var top3snacks int
	fmt.Println("Fattest Elf Contest!")
	for i, es := range eSlice[0:3] {
		top3snacks += es.GetTotalCalories()
		fmt.Printf("Rank %d --> Elf #%d (%d kcal)\n", i, es.Id, es.GetTotalCalories())
	}
	fmt.Printf("Top 3 fattest elves carry %d calories\n", top3snacks)
}
