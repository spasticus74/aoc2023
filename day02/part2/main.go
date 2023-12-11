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
	reGame  = regexp.MustCompile(`^Game (\d{1,6}):.*$`)
	reBlue  = regexp.MustCompile(`(\d{1,6}) blue`)
	reGreen = regexp.MustCompile(`(\d{1,6}) green`)
	reRed   = regexp.MustCompile(`(\d{1,6}) red`)
	result  = 0
)

func main() {
	file, err := os.OpenFile("testdata.txt", os.O_RDONLY, 0)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		thisGame := parseGame(text)
		x := testGame(thisGame)
		fmt.Println(x)
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

func testGame(gameToTest game) int {
	gameMax := selection{}

	for _, selection := range gameToTest.selections {
		if selection.blue > gameMax.blue {
			gameMax.blue = selection.blue
		}
		if selection.green > gameMax.green {
			gameMax.green = selection.green
		}
		if selection.red > gameMax.red {
			gameMax.red = selection.red
		}
	}
	return gameMax.blue * gameMax.green * gameMax.red
}
