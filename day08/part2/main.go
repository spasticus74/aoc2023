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
	pattern      []int
	nodes        map[string]node
	reNode       = regexp.MustCompile(`([A-Z0-9]{3}) = \(([A-Z0-9]{3}), ([A-Z0-9]{3})\)`)
	currentNodes []string
	result       int
)

func main() {
	pattern, nodes = parseInput("input.txt")
	currentNodes = findStartNodes()

	result = getPeriod(currentNodes[0])

	// each path is cyclical with a fixed period, once we have the period of each path the solution is just the Least Common Multiple of them
	for p := 1; p < len(currentNodes); p++ {
		result = LCM(result, getPeriod(currentNodes[p]))
	}

	fmt.Println(result)
}

func getPeriod(startNode string) int {
	period := 0
	for {
		for _, step := range pattern {
			period++
			startNode = nodes[startNode][step]

			if startNode[2:] == "Z" {
				return period
			}
		}
	}
}

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

func areAtEnd(currentNodes []string) bool {
	for _, k := range currentNodes {
		if k[2:] != "Z" {
			return false
		}
	}
	return true
}

func findStartNodes() []string {
	startNodes := make([]string, 0)
	for k := range nodes {
		if k[2:] == "A" {
			startNodes = append(startNodes, k)
		}
	}
	return startNodes
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
