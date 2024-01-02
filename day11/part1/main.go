package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

type Pos struct {
	row int
	col int
}

var (
	field   [][]string
	maxCols int
	maxRows int
	stars   []Pos
)

func main() {

	field = parseInput("input.txt")
	fmt.Println(field)

	stars = findStars()
	fmt.Println(stars)

	expandColumns()
	fmt.Println(stars)

	fmt.Println(calcPaths())
}

func parseInput(fn string) [][]string {
	file, err := os.OpenFile(fn, os.O_RDONLY, 0)
	if err != nil {
		log.Fatal("error opening input file", err)
	}
	defer file.Close()

	f := make([][]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		thisRow := make([]string, 0)
		thisRow = append(thisRow, strings.Split(scanner.Text(), "")...)
		maxCols = len(thisRow)
		f = append(f, thisRow)
		maxRows++

		// if this row has no stars we can add it twice
		if isEmptyRow(thisRow) {
			f = append(f, thisRow)
			maxRows++
		}
	}
	return f
}

func isEmptyRow(row []string) bool {
	for _, p := range row {
		if p == "#" {
			return false
		}
	}
	return true
}

func findStars() []Pos {
	s := make([]Pos, 0)
	for r := 0; r < maxRows; r++ {
		for c := 0; c < maxCols; c++ {
			if field[r][c] == "#" {
				thisPos := Pos{
					row: r,
					col: c,
				}
				s = append(s, thisPos)
			}
		}
	}
	return s
}

func expandColumns() {
	for c := maxCols - 1; c > 0; c-- {
		for r := 0; r < maxRows; r++ {
			if field[r][c] == "#" {
				goto THERE
			}
		}
		// reached here then need to oush column coordinate of stars out
		pushStars(c)
	THERE:
	}
}

func pushStars(col int) {
	for i, s := range stars {
		if s.col > col {
			s.col++
			stars[i] = s
		}
	}
}

func calcPaths() float64 {
	var pathLength float64
	for thisStar := 0; thisStar < len(stars); thisStar++ {
		for nextStar := (thisStar + 1); nextStar < len(stars); nextStar++ {
			pathLength += math.Abs(float64(stars[thisStar].row - stars[nextStar].row))
			pathLength += math.Abs(float64(stars[thisStar].col - stars[nextStar].col))
		}
	}
	return pathLength
}
