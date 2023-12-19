package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

type node [2]string

var (
	pattern     []int
	nodes       map[string]node
	reNode      = regexp.MustCompile(`([A-Z]{3}) = \(([A-Z]{3}), ([A-Z]{3})\)`)
	currentNode = "AAA"
)

func main() {
	pattern, nodes = parseInput("input.txt")

	result := 0
	for {
		for _, step := range pattern {
			result++
			currentNode = nodes[currentNode][step]
			if currentNode == "ZZZ" {
				fmt.Printf("Reached 'ZZZ' in %d steps\n", result)
				goto END
			}
		}
	}
END:
}

func parseInput(fn string) ([]int, map[string]node) {
	steps := make([]int, 0)
	locs := make(map[string]node)

	file, err := os.OpenFile(fn, os.O_RDONLY, 0)
	if err != nil {
		log.Fatal("error opening input file", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// first line is the step pattern
	scanner.Scan()
	stps := strings.Split(scanner.Text(), "")
	for _, s := range stps {
		if s == "L" {
			steps = append(steps, 0)
		} else {
			steps = append(steps, 1)
		}
	}

	// second line is blank
	scanner.Scan()

	for scanner.Scan() {
		nodeMatches := reNode.FindAllStringSubmatch(scanner.Text(), -1)
		if nodeMatches != nil {
			locs[nodeMatches[0][1]] = node{nodeMatches[0][2], nodeMatches[0][3]}
		}
	}

	return steps, locs
}
