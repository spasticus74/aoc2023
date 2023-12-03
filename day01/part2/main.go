package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	var result int
	file, err := os.OpenFile("input.txt", os.O_RDONLY, 0)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data := scanner.Text()
		digits := make([]string, len(data))
		for x := range digits {
			digits[x] = "SKIP"
		}
		for p, ch := range data {
			if unicode.IsDigit(ch) {
				digits[p] = fmt.Sprintf("%c", ch)
			}

			found := strings.Index(data[p:], "one")
			if found > -1 {
				digits[found+p] = "1"
			}

			found = strings.Index(data[p:], "two")
			if found > -1 {
				digits[found+p] = "2"
			}

			found = strings.Index(data[p:], "three")
			if found > -1 {
				digits[found+p] = "3"
			}

			found = strings.Index(data[p:], "four")
			if found > -1 {
				digits[found+p] = "4"
			}

			found = strings.Index(data[p:], "five")
			if found > -1 {
				digits[found+p] = "5"
			}

			found = strings.Index(data[p:], "six")
			if found > -1 {
				digits[found+p] = "6"
			}

			found = strings.Index(data[p:], "seven")
			if found > -1 {
				digits[found+p] = "7"
			}

			found = strings.Index(data[p:], "eight")
			if found > -1 {
				digits[found+p] = "8"
			}

			found = strings.Index(data[p:], "nine")
			if found > -1 {
				digits[found+p] = "9"
			}

		}

		cs := stripSlice(digits)
		r, _ := strconv.Atoi(cs[0] + cs[len(cs)-1])
		fmt.Printf("Parsed %s to %d\n", scanner.Text(), r)
		result += r
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(result)
}

func stripSlice(input []string) []string {
	cleanedSlice := make([]string, 0)

	for _, v := range input {
		if v != "SKIP" {
			cleanedSlice = append(cleanedSlice, v)
		}
	}

	return cleanedSlice
}
