package day1

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

func Part2(input io.Reader) (int, error) {
	if input == nil {
		return 0, errInputNil
	}

	scanner := bufio.NewScanner(input)

	var increments int
	mesuraments := make([]int, 0)

	for scanner.Scan() {
		data := strings.TrimSuffix(scanner.Text(), "\n")

		depth, err := strconv.Atoi(data)
		if err != nil {
			return 0, fmt.Errorf("could not parse input data to int: %s - %w", err, errInputInvalid)
		}
		mesuraments = append(mesuraments, depth)
	}

	// Not enough to worth covering this time ;]
	if len(mesuraments) < 3 {
		return 0, errInputInvalid
	}

	currentSum := mesuraments[0] + mesuraments[1] + mesuraments[2]

	for i, j, k := 1, 2, 3; i <= len(mesuraments)-3; i, j, k = i+1, j+1, k+1 {
		newSum := mesuraments[i] + mesuraments[j] + mesuraments[k]
		if newSum > currentSum {
			increments++
		}
		currentSum = newSum
	}
	return increments, nil
}
