package day2

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

type location struct {
	horizontalPos, depth int
}

func (l *location) move(action string, units int) {
	switch action {
	case "up":
		l.depth = l.depth - units
	case "down":
		l.depth = l.depth + units
	case "forward":
		l.horizontalPos = l.horizontalPos + units
	}
}

func (l *location) position() int { return l.horizontalPos * l.depth }

func Part1(input io.Reader) (int, error) {
	if input == nil {
		return 0, errInputNil
	}

	scanner := bufio.NewScanner(input)
	var loc location

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
