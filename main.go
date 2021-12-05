package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"

	"github.com/alesr/adventofcode2021/day1"
	"github.com/alesr/adventofcode2021/day2"
	"github.com/alesr/adventofcode2021/day3"
)

func main() {
	run("Day 1 Part 1", "day1/data.in", day1.Part1)
	run("Day 1 Part 2", "day1/data.in", day1.Part2)
	run("Day 2 Part 1", "day2/data.in", day2.Part1)
	run("Day 2 Part 2", "day2/data.in", day2.Part2)
	run("Day 3 Part 1", "day3/data.in", day3.Part1)
}

func readFile(filename string) *bytes.Buffer {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("failed to open '%s' data input file: %s\n", filename, err)
	}
	return bytes.NewBuffer(data)
}

func run(title, inputPath string, fn func(io.Reader) (int, error)) {
	input := readFile(inputPath)

	res, err := fn(input)
	if err != nil {
		log.Fatalln("could not execute function:", err)
	}
	fmt.Printf("%s: %d\n", title, res)
}
