package main

import (
	"fmt"
	"log"
	"os"

	"github.com/alesr/adventofcode2021/day1"
)

func main() {
	sonarDataInput, err := os.Open("day1/sonar_sweep_input.txt")
	if err != nil {
		log.Fatalln("failed to open sonar data input file:", err)
	}
	defer closeFile(sonarDataInput)

	result, err := day1.SonarSweep(sonarDataInput)
	if err != nil {
		log.Fatalln("failed to process sonar sweep data:", err)
	}

	fmt.Println("Sonar sweep result:", result)
}

func closeFile(f *os.File) {
	err := f.Close()
	if err != nil {
		log.Fatalln("failed to close file:", err)
	}
}
