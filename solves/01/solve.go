package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"sort"
	"strconv"
	"strings"
)

func main() {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("No caller information")
	}

	dirname := filepath.Dir(filename)

	dat, err := os.ReadFile(filepath.Join(dirname, "../../inputs/01.txt"))
	if err != nil {
		panic(err)
	}

	input := string(dat)

	fmt.Println(PartOne(input))
	fmt.Println(PartTwo(input))
}

func PartOne(input string) int {
	elves := strings.Split(input, "\n\n")

	maxCalories := 0
	for _, elv := range elves {
		calories := 0
		for _, food := range strings.Split(elv, "\n") {
			calory, _ := strconv.Atoi(food)
			calories += calory
		}

		if calories > maxCalories {
			maxCalories = calories
		}
	}

	return maxCalories
}

func PartTwo(input string) int {
	elves := strings.Split(input, "\n\n")

	caloriesSlice := make([]int, len(elves))
	for _, elv := range elves {
		calories := 0
		for _, food := range strings.Split(elv, "\n") {
			calory, _ := strconv.Atoi(food)
			calories += calory
		}

		caloriesSlice = append(caloriesSlice, calories)
	}

	sort.Ints(caloriesSlice)

	return caloriesSlice[len(caloriesSlice)-1] + caloriesSlice[len(caloriesSlice)-2] + caloriesSlice[len(caloriesSlice)-3]
}
