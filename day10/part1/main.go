package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var (
	field      [][]string
	startPos   [2]int //row, col
	currentPos [2]int
	lastPos    [2]int
	stepCount  int
)

func main() {
	field, startPos = parseInput("input.txt")
	currentPos = startPos
	lastPos = startPos

	if currentPos == startPos && stepCount == 0 {
		for {
			switch field[currentPos[0]][currentPos[1]] {
			case "S":
				if !moveNorth() {
					if !moveEast() {
						if !moveSouth() {
							if !moveWest() {
								log.Fatal("no move available", currentPos)
							}
						}
					}
				}
				if currentPos == startPos && stepCount > 0 {
					goto END
				}
			case "F":
				if !moveEast() {
					if !moveSouth() {
						log.Fatal("no move available", currentPos)
					}
				}
				if currentPos == startPos && stepCount > 0 {
					goto END
				}
			case "-":
				if !moveEast() {
					if !moveWest() {
						log.Fatal("no move available", currentPos)
					}
				}
				if currentPos == startPos && stepCount > 0 {
					goto END
				}
			case "7":
				if !moveSouth() {
					if !moveWest() {
						log.Fatal("no move available", currentPos)
					}
				}
				if currentPos == startPos && stepCount > 0 {
					goto END
				}
			case "|":
				if !moveNorth() {
					if !moveSouth() {
						log.Fatal("no move available", currentPos)
					}
				}
				if currentPos == startPos && stepCount > 0 {
					goto END
				}
			case "L":
				if !moveNorth() {
					if !moveEast() {
						log.Fatal("no move available", currentPos)
					}
				}
				if currentPos == startPos && stepCount > 0 {
					goto END
				}
			case "J":
				if !moveNorth() {
					if !moveWest() {
						log.Fatal("no move available", currentPos)
					}
				}
				if currentPos == startPos && stepCount > 0 {
					goto END
				}
			default:
				log.Fatal("fuck knows!", currentPos, field[currentPos[0]][currentPos[1]])
			}
		}
	}
END:
	fmt.Printf("result: %d\n", stepCount/2)
}

func parseInput(fn string) (f [][]string, s [2]int) {
	file, err := os.OpenFile(fn, os.O_RDONLY, 0)
	if err != nil {
		log.Fatal("error opening input file", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var r = 1
	var dummyRow []string
	for scanner.Scan() {
		thisRow := make([]string, 1)
		thisRow[0] = "."
		thisRow = append(thisRow, strings.Split(scanner.Text(), "")...)
		thisRow = append(thisRow, ".")

		if len(dummyRow) == 0 {
			dummyRow = getDummyRow(len(thisRow))
			f = append(f, dummyRow)
		}

		for i, p := range thisRow {
			if p == "S" {
				s[0] = r
				s[1] = i
			}
		}
		f = append(f, thisRow)
		r++
	}
	f = append(f, dummyRow)
	return
}

func getDummyRow(size int) []string {
	r := make([]string, size)
	for i := range r {
		r[i] = "."
	}
	return r
}

func moveNorth() bool {
	// look at the point immediately above and if it's possible
	// to move there return the new coordinates
	newPos := currentPos
	newPos[0] = newPos[0] - 1

	if newPos == lastPos {
		return false
	}

	currentPipe := field[currentPos[0]][currentPos[1]]
	northPipe := field[newPos[0]][newPos[1]]

	if currentPipe == "|" || currentPipe == "L" || currentPipe == "J" || (currentPipe == "S" && (northPipe == "|" || northPipe == "7" || northPipe == "F")) {
		lastPos = currentPos
		currentPos = newPos
		stepCount++
		return true
	}

	return false
}

func moveEast() bool {
	// look at the point immediately right and if it's possible
	// to move there return the new coordinates
	newPos := currentPos
	newPos[1] = newPos[1] + 1

	if newPos == lastPos {
		return false
	}

	currentPipe := field[currentPos[0]][currentPos[1]]
	nextPipe := field[newPos[0]][newPos[1]]

	if currentPipe == "-" || currentPipe == "L" || currentPipe == "F" || (currentPipe == "S" && (nextPipe == "-" || nextPipe == "7" || nextPipe == "J")) {
		lastPos = currentPos
		currentPos = newPos
		stepCount++
		return true
	}

	return false
}

func moveSouth() bool {
	// look at the point immediately below and if it's possible
	// to move there return the new coordinates
	newPos := currentPos
	newPos[0] = newPos[0] + 1

	if newPos == lastPos {
		return false
	}

	currentPipe := field[currentPos[0]][currentPos[1]]
	nextPipe := field[newPos[0]][newPos[1]]

	if currentPipe == "|" || currentPipe == "F" || currentPipe == "7" || (currentPipe == "S" && (nextPipe == "J" || nextPipe == "L" || nextPipe == "|")) {
		lastPos = currentPos
		currentPos = newPos
		stepCount++
		return true
	}

	return false
}

func moveWest() bool {
	// look at the point immediately left and if it's possible
	// to move there return the new coordinates
	newPos := currentPos
	newPos[1] = newPos[1] - 1

	if newPos == lastPos {
		return false
	}

	currentPipe := field[currentPos[0]][currentPos[1]]
	nextPipe := field[newPos[0]][newPos[1]]

	if currentPipe == "-" || currentPipe == "J" || currentPipe == "7" || (currentPipe == "S" && (nextPipe == "L" || nextPipe == "F" || nextPipe == "-")) {
		lastPos = currentPos
		currentPos = newPos
		stepCount++
		return true
	}

	return false
}
