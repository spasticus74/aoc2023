package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type matrixRange struct {
	ln  int
	st  int
	end int
}

var (
	reNumber         = regexp.MustCompile(`(\d{1,3})`)
	reNotNumberOrDot = regexp.MustCompile(`([^0-9.])`)
	result           = 0
	inputMatrix      = make([][]string, 0)
	inputSlice       = make([]string, 0)
	pageHeight       = 0
	pageWidth        = 0
)

func main() {
	file, err := os.OpenFile("input.txt", os.O_RDONLY, 0)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var dummyString []byte
	for scanner.Scan() {
		text := "." + scanner.Text() + "."
		pageWidth = len(text)
		if pageHeight == 0 { // prepend a dummy line of dots
			dummyString = make([]byte, pageWidth)
			for i := 0; i < pageWidth; i++ {
				dummyString[i] = '.'
			}
			inputSlice = append(inputSlice, string(dummyString))
			pageHeight++
		}
		pageHeight++
		inputSlice = append(inputSlice, text)
		inputMatrix = append(inputMatrix, strings.Split(text, ""))
	}
	inputSlice = append(inputSlice, string(dummyString))
	pageHeight++

	// parse the file line by line looking for symbols
	/*for rowNum, row := range inputMatrix {
		for colNum, col := range row {
			if reNotNumberOrDot.FindAllString(col, 1) != nil {
				fmt.Printf("found \"%s\" at [%d,%d]\n", col, rowNum, colNum)
			}
		}
	}*/
	for lineNumber, line := range inputSlice {
		matches := reNumber.FindAllStringSubmatch(line, -1)
		if matches == nil {
			continue
		}
		for _, m := range matches {
			pos := strings.Index(line, m[0])
			if pos > -1 { //found a number
				//fmt.Printf("Found \"%s\" at position [%d,%d] to [%d,%d], ", m[0], lineNumber, pos, lineNumber, pos-1+len(m[0]))
				// look for symbol in the range above
				prop := matrixRange{ln: lineNumber - 1, st: pos - 1, end: pos + len(m[0]) + 1}
				ts := inputSlice[prop.ln][prop.st:prop.end]
				found, sym := FindSymbol(ts)
				if found {
					val, err := strconv.Atoi(m[0])
					if err != nil {
						fmt.Printf("Error converting %s to number, dying", m[0])
						log.Fatal()
					}
					result += val
					var rep string
					for l := 0; l < len(m[0]); l++ {
						rep += "."
					}
					line = strings.Replace(line, m[0], rep, 1)
					fmt.Printf("Found \"%s\" at position [%d,%d] to [%d,%d], found symbol (%s) in range %v, adding %s to total. RESULT: %d\n", m[0], lineNumber, pos, lineNumber, pos-1+len(m[0]), sym, prop, m[0], result)
					continue
				}

				prop = matrixRange{ln: lineNumber, st: pos - 1, end: pos + len(m[0]) + 1}
				ts = inputSlice[prop.ln][prop.st:prop.end]
				found, sym = FindSymbol(ts)
				if found {
					val, err := strconv.Atoi(m[0])
					if err != nil {
						fmt.Printf("Error converting %s to number, dying", m[0])
						log.Fatal()
					}
					result += val
					var rep string
					for l := 0; l < len(m[0]); l++ {
						rep += "."
					}
					line = strings.Replace(line, m[0], rep, 1)
					fmt.Printf("Found \"%s\" at position [%d,%d] to [%d,%d], found symbol (%s) in range %v, adding %s to total. RESULT: %d\n", m[0], lineNumber, pos, lineNumber, pos-1+len(m[0]), sym, prop, m[0], result)
					continue
				}

				prop = matrixRange{ln: lineNumber + 1, st: pos - 1, end: pos + len(m[0]) + 1}
				ts = inputSlice[prop.ln][prop.st:prop.end]
				found, sym = FindSymbol(ts)
				if found {
					val, err := strconv.Atoi(m[0])
					if err != nil {
						fmt.Printf("Error converting %s to number, dying", m[0])
						log.Fatal()
					}
					result += val
					var rep string
					for l := 0; l < len(m[0]); l++ {
						rep += "."
					}
					line = strings.Replace(line, m[0], rep, 1)
					fmt.Printf("Found \"%s\" at position [%d,%d] to [%d,%d], found symbol (%s) in range %v, adding %s to total. RESULT: %d\n", m[0], lineNumber, pos, lineNumber, pos-1+len(m[0]), sym, prop, m[0], result)
					continue
				}
			}
		}
	}

	fmt.Printf("Final result: %d\n", result)

}

/*func LookAround(r, c int) []string {
	lookRow := r
	lookCol := c
	// look upper left
	if reNumber.FindAllString(inputMatrix[lookRow][lookCol], 1) != nil { //found a number
		// look left from here
	}
}*/

func FindSymbol(testString string) (bool, string) {
	r := reNotNumberOrDot.FindAllStringSubmatch(testString, -1)
	if reNotNumberOrDot.FindAllStringSubmatch(testString, -1) != nil {
		return true, r[0][0]
	} else {
		return false, ""
	}
}

/*func SanitizeCoords(proposed matrixRange) (fixed matrixRange) {
	fixed = proposed
	if proposed.ln < 0 {
		fixed.ln = 0
	} else if proposed.ln >= pageHeight {
		fixed.ln = pageHeight - 1
	}

	if proposed.st < 0 {
		fixed.st = 0
	}

	if proposed.end >= pageWidth {
		fixed.end = pageWidth - 1
	}
	return
}*/
