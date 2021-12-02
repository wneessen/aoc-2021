package main

import (
	"fmt"
	"strconv"
	"strings"
)

func part1(i []string) {
	var fpx int
	var fpy int
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
			fpx += n
		case "up":
			fpy -= n
		case "down":
			fpy += n
		}
	}
	fmt.Printf("New position: %d\n", fpx*fpy)
}
