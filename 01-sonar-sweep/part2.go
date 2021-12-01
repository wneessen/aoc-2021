package main

import (
	"fmt"
)

func part2(i []int) {
	var t [][]int
	tl := len(i)
	var a []int
	for c := 0; c < len(i); c++ {
		if c+2 >= tl {
			break
		}
		a = []int{i[c+0], i[c+1], i[c+2]}
		t = append(t, a)
	}

	ps := 0
	tl = len(t)
	ic := 0
	for c := 0; c < tl; c++ {
		sum := t[c][0] + t[c][1] + t[c][2]
		if ps == 0 {
			ps = sum
			continue
		}
		if sum > ps {
			ic++
		}
		ps = sum
	}
	fmt.Printf("Part 2 => Total number of increases: %d\n", ic)
}
