package day3

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

var (
	errInputNil = fmt.Errorf("input is nil")
)

func Part1(input io.Reader) (int, error) {
	if input == nil {
		return 0, errInputNil
	}

	var leastCommon, mostCommon string

	// We want to create a new matrix rotated 90 degrees
	// From:
	//  1,2,3
	//  4,5,6
	//
	// To:
	// 	1,4
	//  2,5
	//  3,6
	columns := make([]string, 0)

	scanner := bufio.NewScanner(input)

	for i := 0; scanner.Scan(); i++ {
		input := strings.TrimSuffix(scanner.Text(), "\n")

		// Assuming that each row has the same lenght,
		// we initizalize len(rows) columns
		if i == 0 {
			for range input {
				columns = append(columns, "")
			}
		}

		for j, v := range input {
			columns[j] += string(v)
		}
	}

	for _, v := range columns {
		zeros := strings.Count(v, "0")
		ones := strings.Count(v, "1")

		if zeros > ones {
			mostCommon += "0"
			leastCommon += "1"
		} else {
			mostCommon += "1"
			leastCommon += "0"
		}
	}

	ε, err := strconv.ParseInt(leastCommon, 2, 64)
	if err != nil {
		return 0, fmt.Errorf("could not parse epsilon: %s", err)
	}

	γ, err := strconv.ParseInt(mostCommon, 2, 64)
	if err != nil {
		return 0, fmt.Errorf("could not parse gamma: %s", err)
	}

	return int(ε * γ), nil
}
