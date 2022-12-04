package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	file, _ := ioutil.ReadFile("input.txt")

	fmt.Println(results(string(file)))
}

func results(input string) int {
	total := 0
	overlapsTotal := 0
	splittedInput := strings.Split(string(input), "\n")

	var lots [][]string

	for i := 0; i < len(splittedInput); i++ {
		lots = append(lots, strings.Split(splittedInput[i], ","))
	}

	for i := 0; i < len(lots); i++ {
		lotOne := lots[i][0]
		lotTwo := lots[i][1]

		if isContained(lotOne, lotTwo) {
			total++
		}

		if overlaps(lotOne, lotTwo) {
			overlapsTotal++
		}
	}

	fmt.Println("part two", overlapsTotal)

	return total
}

func isContained(lotOne string, lotTwo string) bool {
	parsedLotOne := strings.Split(lotOne, "-")
	parsedLotTwo := strings.Split(lotTwo, "-")

	lotOneMin, _ := strconv.Atoi(parsedLotOne[0])
	lotOneMax, _ := strconv.Atoi(parsedLotOne[1])

	lotTwoMin, _ := strconv.Atoi(parsedLotTwo[0])
	lotTwoMax, _ := strconv.Atoi(parsedLotTwo[1])

	if lotOneMin >= lotTwoMin && lotOneMax <= lotTwoMax || lotOneMin <= lotTwoMin && lotOneMax >= lotTwoMax {
		return true
	}
	return false
}

func overlaps(lotOne string, lotTwo string) bool {
	parsedLotOne := strings.Split(lotOne, "-")
	parsedLotTwo := strings.Split(lotTwo, "-")

	lotOneMin, _ := strconv.Atoi(parsedLotOne[0])
	lotOneMax, _ := strconv.Atoi(parsedLotOne[1])

	lotTwoMin, _ := strconv.Atoi(parsedLotTwo[0])
	lotTwoMax, _ := strconv.Atoi(parsedLotTwo[1])

	if lotOneMin >= lotTwoMin && lotOneMin <= lotTwoMax || lotTwoMin >= lotOneMin && lotTwoMin <= lotOneMax {
		return true
	}

	if lotOneMax >= lotTwoMin && lotOneMax <= lotTwoMax || lotTwoMax >= lotOneMin && lotTwoMax <= lotOneMax {
		return true
	}

	return false
}
