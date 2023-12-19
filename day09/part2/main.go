package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type history []int

var (
	values []history
	result int
)

func main() {
	values = parseInput("input.txt")

	for _, v := range values {
		result += predict(v)
	}

	fmt.Printf("result: %d\n", result)
}

func parseInput(fn string) (h []history) {
	file, err := os.OpenFile(fn, os.O_RDONLY, 0)
	if err != nil {
		log.Fatal("error opening input file", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		row := history{}
		hx := strings.Split(scanner.Text(), " ")
		for _, x := range hx {
			v, err := strconv.Atoi(x)
			if err != nil {
				log.Fatal("error converting input", x)
			}
			row = append(row, v)
		}
		h = append(h, row)
	}
	return
}

func predict(h history) int {
	expandedVals := make([]history, 0)
	expandedVals = append(expandedVals, h)
	thisRow := h

	for {
		nextRow := history{}

		for i := 1; i < len(thisRow); i++ {
			nextRow = append(nextRow, thisRow[i]-thisRow[i-1])
		}
		expandedVals = append(expandedVals, nextRow)
		thisRow = nextRow
		if checkAllZero(nextRow) {
			break
		}
	}

	// start from the second last row, subtract the first value
	// of the previous row from the first value of this row
	for r := len(expandedVals) - 2; r >= 0; r-- {
		expandedVals[r] = insert(expandedVals[r], expandedVals[r][0]-expandedVals[r+1][0], 0)
	}

	return expandedVals[0][0]
}

func insert(a []int, c int, i int) []int {
	return append(a[:i], append([]int{c}, a[i:]...)...)
}

func checkAllZero(z history) bool {
	for _, v := range z {
		if v != 0 {
			return false
		}
	}
	return true
}
