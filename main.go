package main

import (
	"os"

	"github.com/plastic041/aoc-2022-go/commands"
)

func main() {
	args := os.Args[1:]

	// go run main.go new [problem number]
	if len(args) == 2 && args[0] == "scaffold" {
		commands.Scaffold(args[1])
		return
	}
}
