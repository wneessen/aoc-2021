package common

import (
	"bufio"
	"os"
	"strconv"
)

// GetInputLinesAsInt reads a input file line by line and returns a int slice
func GetInputLinesAsInt(f string) ([]int, error) {
	fh, err := os.Open(f)
	if err != nil {
		return []int{}, err
	}
	defer func() { _ = fh.Close() }()

	id := make([]int, 0)
	br := bufio.NewScanner(fh)
	for br.Scan() {
		atoi, err := strconv.Atoi(br.Text())
		if err != nil {
			return id, err
		}
		id = append(id, atoi)
	}
	return id, br.Err()
}

// GetInputLinesAsString reads a input file line by line and returns a string slice
func GetInputLinesAsString(f string) ([]string, error) {
	fh, err := os.Open(f)
	if err != nil {
		return []string{}, err
	}
	defer func() { _ = fh.Close() }()

	id := make([]string, 0)
	br := bufio.NewScanner(fh)
	for br.Scan() {
		id = append(id, br.Text())
	}
	return id, br.Err()
}
