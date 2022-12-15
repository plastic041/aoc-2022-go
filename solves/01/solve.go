package main

import (
	"fmt"
	"os"
)

func main() {
	dat, err := os.ReadFile("../../inputs/01.txt")
	if err != nil {
		panic(err)
	}

	input := string(dat)

	fmt.Println(PartOne(input))
	fmt.Println(PartTwo(input))
}

func PartOne(input string) int {
	return 0
}

func PartTwo(input string) int {
	return 0
}
