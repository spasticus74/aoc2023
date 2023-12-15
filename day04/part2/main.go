package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type gameCard struct {
	wins  int
	count int
}

var (
	parseCardRE = regexp.MustCompile(`(\d+):(.*)\|(.*)`)
	deck        = make(map[int]gameCard)
	hand        = make(map[int]gameCard)
	wins        = make(map[int]gameCard)
)

func main() {
	//var result float64
	file, err := os.OpenFile("input.txt", os.O_RDONLY, 0)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		card := parseCardRE.FindAllStringSubmatch(scanner.Text(), -1)
		if card != nil {
			id, gc := processCard(card)
			deck[id] = gc
			hand[id] = gc
		}
	}

	// start processing
	for {
		wins = processHand(hand)
		fmt.Printf("Win: %v\n", wins)
		for i, v := range hand {
			gc := deck[i]
			gc.count += v.count
			deck[i] = gc
		}
		if len(wins) == 0 {
			break
		}
		hand = wins
	}

	fmt.Printf("Result: %d\n", calcResult())
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

func processCard(card [][]string) (int, gameCard) {
	gc := gameCard{}
	id, err := strconv.Atoi(card[0][1])
	if err != nil {
		log.Fatal(card[0][1])
	}

	ourNums := parseNumbers(card[0][2])
	drawnNums := parseNumbers(card[0][3])

	gc.wins = getWinCount(ourNums, drawnNums)
	gc.count = 1

	return id, gc
}

func processHand(thisHand map[int]gameCard) (thisWins map[int]gameCard) {
	thisWins = make(map[int]gameCard)
	for index, thisCard := range thisHand { // for every card we currently hold
		for x := 1; x <= thisCard.wins; x++ {
			// put 'count' copies into the win map
			gc, err := getCardFromDeck(index + x)
			if err == nil {
				if winCard, existsinWin := thisWins[index+x]; existsinWin {
					winCard.count += thisCard.count
					thisWins[index+x] = winCard
				} else {
					gc.count = thisCard.count
					thisWins[index+x] = gc
				}
			}
		}
	}
	return
}

func getCardFromDeck(index int) (gameCard, error) {
	if gc, exists := deck[index]; exists {
		return gc, nil
	}
	return gameCard{}, errors.New("does not exist")
}

func calcResult() int {
	result := 0
	for _, v := range deck {
		result += (v.count - 1)
	}
	return result
}
