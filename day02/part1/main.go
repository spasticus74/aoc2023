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

type selection struct {
	blue  int
	green int
	red   int
}

type game struct {
	id         int
	selections []selection
}

var (
	reGame        = regexp.MustCompile(`^Game (\d{1,6}):.*$`)
	reBlue        = regexp.MustCompile(`(\d{1,6}) blue`)
	reGreen       = regexp.MustCompile(`(\d{1,6}) green`)
	reRed         = regexp.MustCompile(`(\d{1,6}) red`)
	testSelection = selection{red: 12, green: 13, blue: 14}
	result        = 0
)

func main() {
	file, err := os.OpenFile("input.txt", os.O_RDONLY, 0)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		thisGame := parseGame(text)
		x := testGame(thisGame, testSelection)
		result += x
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Final result: %d\n", result)
}

func parseGame(text string) game {
	g := game{}

	// Capture game id
	matches := reGame.FindStringSubmatch(text)
	z, err := strconv.Atoi(matches[1])
	if err != nil {
		fmt.Printf("Broke on [%s]\n", text)
	}
	g.id = z

	// split out selections
	text = text[strings.Index(text, ":")+1:]
	s := strings.Split(text, ";")

	for _, sample := range s {
		thisSelection := selection{}
		blueMatches := reBlue.FindStringSubmatch(sample)
		if blueMatches != nil {
			thisSelection.blue, _ = strconv.Atoi(blueMatches[1])
		}
		greenMatches := reGreen.FindStringSubmatch(sample)
		if greenMatches != nil {
			thisSelection.green, _ = strconv.Atoi(greenMatches[1])
		}
		redMatches := reRed.FindStringSubmatch(sample)
		if redMatches != nil {
			thisSelection.red, _ = strconv.Atoi(redMatches[1])
		}
		g.selections = append(g.selections, thisSelection)
	}

	return g
}

func testGame(gameToTest game, testSelection selection) int {

	for _, selection := range gameToTest.selections {
		if selection.blue > testSelection.blue || selection.green > testSelection.green || selection.red > testSelection.red {
			return 0
		}
	}
	return gameToTest.id
}
