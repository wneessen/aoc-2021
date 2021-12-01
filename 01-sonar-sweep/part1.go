package main

import (
	"fmt"
)

func part1(i []int) {
	first := true
	ic := 0
	pv := 0
	for _, cv := range i {
		if first {
			pv = cv
			first = false
			continue
		}
		if cv > pv {
			ic++
		}
		pv = cv
	}
	fmt.Printf("Part 1 => Total number of increases: %d\n", ic)
}
