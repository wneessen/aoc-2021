package main

import (
	"fmt"
	"strconv"
	"strings"
)

func part2(i []string) {
	var fpx int
	var fpy int
	var aim int
	for _, ci := range i {
		ps := strings.SplitN(ci, " ", 2)
		if len(ps) != 2 {
			break
		}
		p := ps[0]
		n, err := strconv.Atoi(ps[1])
		if err != nil {
			fmt.Printf("could not convert to int: %s", err)
			break
		}
		switch p {
		case "forward":
			fpy += n * aim
			fpx += n
		case "up":
			aim -= n
		case "down":
			aim += n
		}
	}
	fmt.Printf("New position: %d\n", fpx*fpy)
}
