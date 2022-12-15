package main_test

import (
	"os"
	"testing"

	solve01 "github.com/plastic041/aoc-2022-go/solves/01"
	"github.com/stretchr/testify/assert"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func TestPartOne(t *testing.T) {
	dat, err := os.ReadFile("../../examples/01.txt")
	check(err)

	assert.Equal(t, 24000, solve01.PartOne(string(dat)))
}

func TestPartTwo(t *testing.T) {
	dat, err := os.ReadFile("../../examples/01.txt")
	check(err)

	assert.Equal(t, 45000, solve01.PartTwo(string(dat)))
}
