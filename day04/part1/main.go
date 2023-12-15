package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var (
	parseCardRE = regexp.MustCompile(`(\d+):(.*)\|(.*)`)
)

func main() {
	var result float64
	file, err := os.OpenFile("input.txt", os.O_RDONLY, 0)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		card := parseCardRE.FindAllStringSubmatch(scanner.Text(), -1)
		if card != nil {
			ourNums := parseNumbers(card[0][2])
			drawnNums := parseNumbers(card[0][3])
			wins := getWinCount(ourNums, drawnNums)
			fmt.Printf("Card %s gave %d wins - %f points\n", card[0][1], wins, math.Pow(2, (float64(wins)-1)))
			if wins > 0 {
				result += math.Pow(2, (float64(wins) - 1))
			}
		}
	}

	fmt.Println(result)
}

func parseNumbers(in string) (out []int) {
	n := strings.Split(strings.TrimSpace(strings.ReplaceAll(in, "  ", " ")), " ")
	for _, v := range n {
		x, err := strconv.Atoi(v)
		if err != nil {
			log.Fatal(v, err)
		}
		out = append(out, x)
	}
	return
}

func getWinCount(ours, drawn []int) (wins int) {
	for _, o := range ours {
		for _, d := range drawn {
			if o == d {
				wins++
				continue
			}
		}
	}
	return
}
