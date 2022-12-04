package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	input, _ := ioutil.ReadFile("input.txt")

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

		parsedLotOne := strings.Split(lotOne, "-")
		parsedLotTwo := strings.Split(lotTwo, "-")

		lotOneMin, _ := strconv.Atoi(parsedLotOne[0])
		lotOneMax, _ := strconv.Atoi(parsedLotOne[1])

		lotTwoMin, _ := strconv.Atoi(parsedLotTwo[0])
		lotTwoMax, _ := strconv.Atoi(parsedLotTwo[1])

		if isContained(lotOneMin, lotOneMax, lotTwoMin, lotTwoMax) {
			total++
		}

		if overlappedLots(lotOneMin, lotOneMax, lotTwoMin, lotTwoMax) {
			overlapsTotal++
		}
	}

	fmt.Println("Part one:", total)
	fmt.Println("Part two:", overlapsTotal)
}

func isContained(lotOneMin int, lotOneMax int, lotTwoMin int, lotTwoMax int) bool {
	if lotOneMin >= lotTwoMin && lotOneMax <= lotTwoMax || lotOneMin <= lotTwoMin && lotOneMax >= lotTwoMax {
		return true
	}
	return false
}

func overlappedLots(lotOneMin int, lotOneMax int, lotTwoMin int, lotTwoMax int) bool {
	if lotOneMin >= lotTwoMin && lotOneMin <= lotTwoMax || lotTwoMin >= lotOneMin && lotTwoMin <= lotOneMax {
		return true
	}

	if lotOneMax >= lotTwoMin && lotOneMax <= lotTwoMax || lotTwoMax >= lotOneMin && lotTwoMax <= lotOneMax {
		return true
	}

	return false
}
