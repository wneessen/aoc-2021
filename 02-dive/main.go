package main

import (
	"github.com/wneessen/aoc-2021/common"
	"log"
	"os"
)

func main() {
	i, err := common.GetInputLinesAsString("input.data")
	if err != nil {
		log.Printf("failed to read input: %s", err)
		os.Exit(1)
	}
	part1(i)
	part2(i)
}
