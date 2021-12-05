package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/alesr/adventofcode2021/day1"
	"github.com/alesr/adventofcode2021/day2"
)

func main() {
	day1P1Input := readFile("day1/data.in")
	day1P2Input := *day1P1Input

	day1P1Result, err := day1.Part1(day1P1Input)
	if err != nil {
		log.Fatalln("failed to process day 1 part 1:", err)
	}
	fmt.Println("Day 1 part 1:", day1P1Result)

	day1P2Result, err := day1.Part2(&day1P2Input)
	if err != nil {
		log.Fatalln("failed to process day 1 part 2:", err)
	}
	fmt.Println("Day 1 part 2:", day1P2Result)

	day2P1Input := readFile("day2/data.in")

	day2P1Result, err := day2.Part1(day2P1Input)
	if err != nil {
		log.Fatalln("failed to process day 2 part 1:", err)
	}
	fmt.Println("Day 2 part 1:", day2P1Result)
}

func readFile(filename string) *bytes.Buffer {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("failed to open '%s' data input file: %s\n", filename, err)
	}
	return bytes.NewBuffer(data)
}
