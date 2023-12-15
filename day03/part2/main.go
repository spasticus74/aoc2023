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
type coords struct {
	r int
	c int
}

var (
	reNumber    = regexp.MustCompile(`(\d{1,3})`)
	reGear      = regexp.MustCompile(`(\*)`)
	result      = 0
	inputMatrix = make([][]string, 0)
	inputSlice  = make([]string, 0)
	pageHeight  = 0
	pageWidth   = 0
	foundGears  map[coords]int //make(map[matrixRange][]int, 0)
)

func main() {
	file, err := os.OpenFile("test.txt", os.O_RDONLY, 0)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	foundGears = make(map[coords]int, 0)

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

	for lineNumber, line := range inputSlice {
		matches := reNumber.FindAllStringSubmatch(line, -1)
		if matches == nil {
			continue
		}
		for _, m := range matches {
			pos := strings.Index(line, m[0])
			if pos > -1 { //found a number
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
					gearPos := coords{r: prop.ln, c: prop.st + sym}
					vals, ok := foundGears[gearPos]
					if !ok {
						foundGears[gearPos] = val
					} else {
						foundGears[gearPos] = vals * val
						result += foundGears[gearPos]
					}

					//result += val
					var rep string
					for l := 0; l < len(m[0]); l++ {
						rep += "."
					}
					line = strings.Replace(line, m[0], rep, 1)
					//fmt.Printf("Found \"%s\" at position [%d,%d] to [%d,%d], found symbol (%s) in range %v, adding %s to total. RESULT: %d\n", m[0], lineNumber, pos, lineNumber, pos-1+len(m[0]), sym, prop, m[0], result)
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
					gearPos := coords{r: prop.ln, c: prop.st + sym}
					vals, ok := foundGears[gearPos]
					if !ok {
						foundGears[gearPos] = val
					} else {
						foundGears[gearPos] = vals * val
						result += foundGears[gearPos]
					}

					var rep string
					for l := 0; l < len(m[0]); l++ {
						rep += "."
					}
					line = strings.Replace(line, m[0], rep, 1)
					//fmt.Printf("Found \"%s\" at position [%d,%d] to [%d,%d], found symbol (%s) in range %v, adding %s to total. RESULT: %d\n", m[0], lineNumber, pos, lineNumber, pos-1+len(m[0]), sym, prop, m[0], result)
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
					gearPos := coords{r: prop.ln, c: prop.st + sym}
					vals, ok := foundGears[gearPos]
					if !ok {
						foundGears[gearPos] = val
					} else {
						foundGears[gearPos] = vals * val
						result += foundGears[gearPos]
					}

					var rep string
					for l := 0; l < len(m[0]); l++ {
						rep += "."
					}
					line = strings.Replace(line, m[0], rep, 1)
					//fmt.Printf("Found \"%s\" at position [%d,%d] to [%d,%d], found symbol (%s) in range %v, adding %s to total. RESULT: %d\n", m[0], lineNumber, pos, lineNumber, pos-1+len(m[0]), sym, prop, m[0], result)
					continue
				}
			}
		}
	}
	fmt.Println(foundGears)
	fmt.Printf("Final result: %d\n", result)

}

func FindSymbol(testString string) (bool, int) {
	x := reGear.FindIndex([]byte(testString))
	if x != nil {
		return true, x[0]
	} else {
		return false, 0
	}
}
