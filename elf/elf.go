package elf

type Elf struct {
	Id       int
	Calories []int
}

func (e Elf) GetTotalCalories() int {
	total := 0
	for _, c := range e.Calories {
		total += c
	}
	return total
}
