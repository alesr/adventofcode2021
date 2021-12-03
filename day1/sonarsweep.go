package day1

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

var (
	errInputInvalid = fmt.Errorf("invalid input")
	errInputNil     = fmt.Errorf("input is nil")
)

func SonarSweep(input io.Reader) (int, error) {
	if input == nil {
		return 0, errInputNil
	}

	var previousDepth, increments int

	scanner := bufio.NewScanner(input)

	for i := 0; scanner.Scan(); i++ {
		data := strings.TrimSuffix(scanner.Text(), "\n")

		depth, err := strconv.Atoi(data)
		if err != nil {
			return 0, fmt.Errorf("could not parse input data to int: %s - %w", err, errInputInvalid)
		}

		// No increment, as we just got our first depth input
		if i == 0 {
			previousDepth = depth
			continue
		}

		if depth > previousDepth {
			increments++
		}

		previousDepth = depth
	}

	return increments, nil
}
