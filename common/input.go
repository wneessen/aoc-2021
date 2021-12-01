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
