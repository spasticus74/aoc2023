package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
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
		digits := make([]string, 0)
		for _, ch := range scanner.Text() {
			if unicode.IsDigit(ch) {
				digits = append(digits, fmt.Sprintf("%c", ch))
			}
		}
		r, _ := strconv.Atoi(digits[0] + digits[len(digits)-1])
		fmt.Printf("Parsed %s to %d\n", scanner.Text(), r)
		result += r
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(result)
}
