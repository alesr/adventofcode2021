package day2

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type locationImproved struct {
	horizontalPos, depth, aim int
}

func (l *locationImproved) move(action string, units int) {
	switch action {
	case "up":
		l.aim -= units
	case "down":
		l.aim += units
	case "forward":
		l.horizontalPos = l.horizontalPos + units
		l.depth += l.aim * units
	}
}

func (l *locationImproved) position() int { return l.horizontalPos * l.depth }

func Part2(input io.Reader) (int, error) {
	if input == nil {
		return 0, errInputNil
	}

	scanner := bufio.NewScanner(input)
	var loc locationImproved

	for scanner.Scan() {
		data := strings.TrimSuffix(scanner.Text(), "\n")
		inputs := strings.Split(data, " ")

		action := inputs[0]
		units, err := strconv.Atoi(inputs[1])
		if err != nil {
			return 0, fmt.Errorf("could not parse input data to int: %s - %w", err, errInputInvalid)
		}

		loc.move(action, units)
	}
	return loc.position(), nil
}
