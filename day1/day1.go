package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

const test = `L68
L30
R48
L5
R60
L55
L1
L99
R14
L82`

func main() {
	useTest := flag.Bool("t", false, "Use test input")
	flag.Parse()
	var reader io.Reader
	if *useTest {
		// Use the test string (TrimSpace removes leading/trailing newlines typical in backticks)
		reader = strings.NewReader(strings.TrimSpace(test))
	} else {
		// Use the real file
		file, err := os.Open("day1/day1.txt")
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		reader = file
	}
	input := parseInput(reader)

	start := time.Now()
	log.Printf("Part 1: %v\n", part1(input))
	log.Printf("Part 1 duration: %v\n", time.Since(start))

	start = time.Now()
	log.Printf("Part 2: %v\n", part2(input))
	log.Printf("Part 2 duration: %v\n", time.Since(start))
}

func part1(lines []string) int {
	pos := 50
	numZero := 0
	for i, l := range lines {
		direction := l[0]
		ticks, err := strconv.Atoi(l[1:])
		if err != nil {
			log.Fatal(fmt.Sprintf("Failed part 1 ticks parsing at row %v:", i), err)
		}
		var new int
		switch string(direction) {
		case "L":
			new = pos - ticks%100
			if new < 0 {
				new = 100 + new
			}

		case "R":
			new = pos + ticks%100
			if new > 99 {
				new = new - 100
			}
		}
		pos = new
		if pos == 0 {
			numZero++
		}
	}
	return numZero
}

func part2(lines []string) int {
	pos := 50
	numZero := 0
	for i, l := range lines {
		direction := l[0]
		ticks, err := strconv.Atoi(l[1:])
		if err != nil {
			log.Fatal(fmt.Sprintf("Failed part 2 ticks parsing at row %v:", i), err)
		}
		var new int
		zeroClick := 0
		switch string(direction) {
		case "L":
			new = pos - ticks%100
			if new < 0 {
				new = 100 + new
				if pos != 0 {
					zeroClick++
				}
			}

		case "R":
			new = pos + ticks%100
			if new > 99 {
				new = new - 100
				zeroClick++

			}
		}
		pos = new
		if pos == 0 && zeroClick == 0 {
			numZero++
		}

		if ticks > 0 {
			zeroClick += (ticks / 100)
		}

		numZero += zeroClick
	}
	return numZero
}

func parseInput(r io.Reader) []string {
	scanner := bufio.NewScanner(r)
	var lines []string

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return lines
}
