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

type mapping [3]int //destination,start,range

var (
	reSeeds        = regexp.MustCompile(`seeds:(.*)`)
	reSeedsToSoil  = regexp.MustCompile(`seed-to-soil map:(.*)`)
	reSoilToFert   = regexp.MustCompile(`soil-to-fertilizer map:(.*)`)
	reFertToWater  = regexp.MustCompile(`fertilizer-to-water map:(.*)`)
	reWaterToLight = regexp.MustCompile(`water-to-light map:(.*)`)
	reLightTemp    = regexp.MustCompile(`light-to-temperature map:(.*)`)
	reTempHumd     = regexp.MustCompile(`temperature-to-humidity map:(.*)`)
	reHumdLoc      = regexp.MustCompile(`humidity-to-location map:(.*)`)
	seeds          []int
	seedToSoil     []mapping
	soilToFert     []mapping
	fertToWater    []mapping
	waterToLight   []mapping
	lightToTemp    []mapping
	tempToHumd     []mapping
	humdToLoc      []mapping
	results        []int
)

func main() {
	parseInput("input.txt")

	for _, s := range seeds {
		x := findInMapping(s, seedToSoil)
		x = findInMapping(x, soilToFert)
		x = findInMapping(x, fertToWater)
		x = findInMapping(x, waterToLight)
		x = findInMapping(x, lightToTemp)
		x = findInMapping(x, tempToHumd)
		x = findInMapping(x, humdToLoc)
		results = append(results, x)
		fmt.Printf("Seed %d should be planted at location %d\n", s, x)
	}
	sort.Ints(results)
	fmt.Printf("Lowest Location Number: %d\n", results[0])
}

func parseInput(fn string) {
	file, err := os.OpenFile(fn, os.O_RDONLY, 0)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		seedsMatch := reSeeds.FindAllStringSubmatch(line, 1)
		seedSoilMatch := reSeedsToSoil.FindAllStringSubmatch(line, 1)
		soilFertMatch := reSoilToFert.FindAllStringSubmatch(line, 1)
		fertWaterMatch := reFertToWater.FindAllStringSubmatch(line, 1)
		waterLightMatch := reWaterToLight.FindAllStringSubmatch(line, 1)
		lightTempMatch := reLightTemp.FindAllStringSubmatch(line, 1)
		tempHumdMatch := reTempHumd.FindAllStringSubmatch(line, 1)
		humdLocMatch := reHumdLoc.FindAllStringSubmatch(line, 1)
		if seedsMatch != nil { //first line, has seed Ids
			n := strings.Split(strings.TrimSpace(seedsMatch[0][1]), " ")
			for _, v := range n {
				x, err := strconv.Atoi(v)
				if err != nil {
					log.Fatal(v, err)
				}
				seeds = append(seeds, x)
			}
		} else if seedSoilMatch != nil { //has seed-soil mappings
			n := strings.Split(strings.TrimSpace(seedSoilMatch[0][1]), ";")
			for _, x := range n {
				seedToSoil = append(seedToSoil, processMap(x)...)
			}
		} else if soilFertMatch != nil { //has soil-fertilizer mappings
			n := strings.Split(strings.TrimSpace(soilFertMatch[0][1]), ";")
			for _, x := range n {
				soilToFert = append(soilToFert, processMap(x)...)
			}
		} else if fertWaterMatch != nil { //has fertilizer-water mappings
			n := strings.Split(strings.TrimSpace(fertWaterMatch[0][1]), ";")
			for _, x := range n {
				fertToWater = append(fertToWater, processMap(x)...)
			}
		} else if waterLightMatch != nil { //has water-light mappings
			n := strings.Split(strings.TrimSpace(waterLightMatch[0][1]), ";")
			for _, x := range n {
				waterToLight = append(waterToLight, processMap(x)...)
			}
		} else if lightTempMatch != nil { //has light-temp mappings
			n := strings.Split(strings.TrimSpace(lightTempMatch[0][1]), ";")
			for _, x := range n {
				lightToTemp = append(lightToTemp, processMap(x)...)
			}
		} else if tempHumdMatch != nil { //has temp-humidity mappings
			n := strings.Split(strings.TrimSpace(tempHumdMatch[0][1]), ";")
			for _, x := range n {
				tempToHumd = append(tempToHumd, processMap(x)...)
			}
		} else if humdLocMatch != nil { //has humidity-location mappings
			n := strings.Split(strings.TrimSpace(humdLocMatch[0][1]), ";")
			for _, x := range n {
				humdToLoc = append(humdToLoc, processMap(x)...)
			}
		}
	}
}

func processMap(s string) (thisMapping []mapping) {
	components := strings.Split(s, " ")
	destBase, err := strconv.Atoi(components[0])
	if err != nil {
		log.Fatal("error converting destination", components[0])
	}
	srcBase, err := strconv.Atoi(components[1])
	if err != nil {
		log.Fatal("error converting source", components[1])
	}
	mapRange, err := strconv.Atoi(components[2])
	if err != nil {
		log.Fatal("error converting range", components[2])
	}

	m := mapping{}
	m[0] = destBase
	m[1] = srcBase
	m[2] = mapRange
	thisMapping = append(thisMapping, m)

	return
}

func findInMapping(in int, m []mapping) int {
	for _, thisMap := range m {
		if in >= thisMap[1] && in < (thisMap[1]+thisMap[2]) {
			return (in - thisMap[1]) + thisMap[0]
		}
	}
	return in
}
