package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type record struct {
	time int
	dist int
}

var (
	records []record
	result  = 1
)

func main() {
	records = parseInput("input.txt")

	for i, race := range records {
		raceOptions := calcRaceOptions(race.time)
		wins := calcWinningCombinations(raceOptions, race.dist)
		result *= wins
		fmt.Printf("Race %d has %d winning options\n", i, wins)
	}
	fmt.Printf("Final result: %d\n", result)
}

func parseInput(fn string) (recs []record) {
	file, err := os.OpenFile(fn, os.O_RDONLY, 0)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		rec := record{}
		line := scanner.Text()
		vals := strings.Split(line, ",")
		t, err := strconv.Atoi(vals[0])
		if err != nil {
			log.Fatal("error converting time", vals[0])
		}
		rec.time = t
		d, err := strconv.Atoi(vals[1])
		if err != nil {
			log.Fatal("error converting distance", vals[1])
		}
		rec.dist = d
		recs = append(recs, rec)
	}
	return
}

func calcRaceOptions(dur int) (options [][2]int) {
	for t := 0; t < dur; t++ {
		var opt [2]int
		opt[0] = t
		opt[1] = (dur - t) * t
		options = append(options, opt)
	}
	return
}

func calcWinningCombinations(options [][2]int, record int) (wins int) {

	for _, opt := range options {
		if opt[1] > record {
			wins++
		}
	}

	return
}
