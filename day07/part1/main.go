package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type Hand struct {
	cards    []string
	cardVal  []int
	bid      int
	handType int
	rank     int
}

var (
	reHand = regexp.MustCompile(`(.*) (\d{1,4})`)
)

func main() {
	x := parseInput("input.txt")
	y := rankHands(x)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Printf("Total winnings : %d\n", calcScore(y))
}

func parseInput(fn string) (hands []Hand) {
	file, err := os.OpenFile(fn, os.O_RDONLY, 0)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		matches := reHand.FindAllStringSubmatch(scanner.Text(), -1)
		if matches != nil {
			thisHand := Hand{}
			thisHand.cards = strings.Split(matches[0][1], "")
			thisHand.cardVal = getCardValues(thisHand.cards)
			b, err := strconv.Atoi(matches[0][2])
			if err != nil {
				log.Fatal("error converting bid", matches[0][2])
			}
			thisHand.bid = b
			thisHand.handType = getHandType(thisHand.cards)
			hands = append(hands, thisHand)
		}
	}
	return
}

// return the 'type' of a hand as int
func getHandType(cards []string) int {
	index := make(map[string]int)
	for _, c := range cards {
		index[c] += 1
	}

	switch len(index) {
	case 1: // five-of-a-kind!
		return 7
	case 2: // could be four-of-a-kind, or a full house
		if index[cards[0]] == 4 || index[cards[0]] == 1 {
			return 6
		} else {
			return 5
		}
	case 3: // could be three-of-a-kind, or two pair
		for i := 0; i < 5; i++ {
			if index[cards[i]] == 3 {
				return 4
			}
		}
		return 3
	case 4: // one pair
		return 2
	case 5: // nothing
		return 1
	}
	fmt.Printf("%v: %d\n", cards, len(index))
	return -1
}

func getCardValues(cards []string) []int {
	cv := make([]int, len(cards))
	for i, c := range cards {
		switch c {
		case "A":
			cv[i] = 13
		case "K":
			cv[i] = 12
		case "Q":
			cv[i] = 11
		case "J":
			cv[i] = 10
		case "T":
			cv[i] = 9
		case "9":
			cv[i] = 8
		case "8":
			cv[i] = 7
		case "7":
			cv[i] = 6
		case "6":
			cv[i] = 5
		case "5":
			cv[i] = 4
		case "4":
			cv[i] = 3
		case "3":
			cv[i] = 2
		case "2":
			cv[i] = 1
		}
	}
	return cv
}

func rankHands(hands []Hand) (rankedHands []Hand) {
	currentRank := 1
	// look at 'High Card' hands first
	hc := make([]Hand, 0)
	for _, h := range hands {
		if h.handType == 1 {
			hc = append(hc, h)
		}
	}
	sort.Slice(hc, func(i, j int) bool {
		for p := 0; p < 5; p++ {
			if hc[i].cardVal[p] < hc[j].cardVal[p] {
				return true
			} else if hc[i].cardVal[p] > hc[j].cardVal[p] {
				return false
			}
		}
		return true
	})
	for _, x := range hc {
		x.rank = currentRank
		rankedHands = append(rankedHands, x)
		currentRank++
	}

	// next 'One Pair' hands
	op := make([]Hand, 0)
	for _, h := range hands {
		if h.handType == 2 {
			op = append(op, h)
		}
	}
	sort.Slice(op, func(i, j int) bool {
		for p := 0; p < 5; p++ {
			if op[i].cardVal[p] < op[j].cardVal[p] {
				return true
			} else if op[i].cardVal[p] > op[j].cardVal[p] {
				return false
			}
		}
		return true
	})
	for _, x := range op {
		x.rank = currentRank
		rankedHands = append(rankedHands, x)
		currentRank++
	}

	// next 'Two Pair' hands
	tp := make([]Hand, 0)
	for _, h := range hands {
		if h.handType == 3 {
			tp = append(tp, h)
		}
	}
	sort.Slice(tp, func(i, j int) bool {
		for p := 0; p < 5; p++ {
			if tp[i].cardVal[p] < tp[j].cardVal[p] {
				return true
			} else if tp[i].cardVal[p] > tp[j].cardVal[p] {
				return false
			}
		}
		return true
	})
	for _, x := range tp {
		x.rank = currentRank
		rankedHands = append(rankedHands, x)
		currentRank++
	}

	// next 'Three-of-a-kind' hands
	toak := make([]Hand, 0)
	for _, h := range hands {
		if h.handType == 4 {
			toak = append(toak, h)
		}
	}
	sort.Slice(toak, func(i, j int) bool {
		for p := 0; p < 5; p++ {
			if toak[i].cardVal[p] < toak[j].cardVal[p] {
				return true
			} else if toak[i].cardVal[p] > toak[j].cardVal[p] {
				return false
			}
		}
		return true
	})
	for _, x := range toak {
		x.rank = currentRank
		rankedHands = append(rankedHands, x)
		currentRank++
	}

	// next 'Full house' hands
	fh := make([]Hand, 0)
	for _, h := range hands {
		if h.handType == 5 {
			fh = append(fh, h)
		}
	}
	sort.Slice(fh, func(i, j int) bool {
		for p := 0; p < 5; p++ {
			if fh[i].cardVal[p] < fh[j].cardVal[p] {
				return true
			} else if fh[i].cardVal[p] > fh[j].cardVal[p] {
				return false
			}
		}
		return true
	})
	for _, x := range fh {
		x.rank = currentRank
		rankedHands = append(rankedHands, x)
		currentRank++
	}

	// next 'Four-of-a-kind' hands
	foak := make([]Hand, 0)
	for _, h := range hands {
		if h.handType == 6 {
			foak = append(foak, h)
		}
	}
	sort.Slice(foak, func(i, j int) bool {
		for p := 0; p < 5; p++ {
			if foak[i].cardVal[p] < foak[j].cardVal[p] {
				return true
			} else if foak[i].cardVal[p] > foak[j].cardVal[p] {
				return false
			}
		}
		return true
	})
	for _, x := range foak {
		x.rank = currentRank
		rankedHands = append(rankedHands, x)
		currentRank++
	}

	// finally 'Five-of-a-kind' hands
	five := make([]Hand, 0)
	for _, h := range hands {
		if h.handType == 7 {
			five = append(five, h)
		}
	}
	sort.Slice(five, func(i, j int) bool {
		for p := 0; p < 5; p++ {
			if five[i].cardVal[p] < five[j].cardVal[p] {
				return true
			} else if five[i].cardVal[p] > five[j].cardVal[p] {
				return false
			}
		}
		return true
	})
	for _, x := range five {
		x.rank = currentRank
		rankedHands = append(rankedHands, x)
		currentRank++
	}

	return
}

/*func getHandScore(h Hand) int {
	score := 0
	score += (5 * h.cardVal[0])
	score += (4 * h.cardVal[1])
	score += (3 * h.cardVal[2])
	score += (2 * h.cardVal[3])
	score += h.cardVal[4]
	return score
}*/

func calcScore(h []Hand) int {
	result := 0
	for _, hand := range h {
		fmt.Printf("Cards: %v, rank: %d, bid: %d, score: %d\n", hand.cards, hand.rank, hand.bid, hand.rank*hand.bid)
		result += hand.bid * hand.rank
	}
	return result
}
