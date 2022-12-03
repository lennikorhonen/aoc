package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	input, _ := ioutil.ReadFile("input.txt")

	fmt.Println(sumOfMispalces(string(input)))
	fmt.Println(badgePrioSum(string(input)))
}

func getPriority(c rune) int {
	if c <= 'Z' {
		return int(c - 'A' + 27)
	}
	return int(c - 'a' + 1)
}

func sumOfMispalces(rucksack string) int {
	sum := 0

	for _, compartment := range strings.Split(rucksack, "\n") {
		half := len(compartment) / 2
		firstHalf, secondHalf := compartment[:half], compartment[half:]

		runeMap := make(map[rune]bool)

		for _, c := range firstHalf {
			runeMap[c] = false
		}

		for _, c := range secondHalf {
			if done, ok := runeMap[c]; ok && !done {
				sum += getPriority(c)
				runeMap[c] = true
			}
		}
	}

	return sum
}

func badgePrioSum(rucksack string) int {
	elves := strings.Split(rucksack, "\n")
	sum := 0
	for i := 0; i < len(elves); i += 3 {
		prios := make([]int, 53)
		elf1 := elves[i]
		elf2 := elves[i+1]
		elf3 := elves[i+2]

		for _, c := range elf1 {
			prio := getPriority(c)
			if prios[prio] == 0 {
				prios[prio] = 1
			}
		}

		for _, c := range elf2 {
			prio := getPriority(c)
			if prios[prio] == 1 {
				prios[prio] = 2
			}
		}

		for _, c := range elf3 {
			prio := getPriority(c)
			if prios[prio] == 2 {
				sum += prio
				break
			}
		}
	}

	return sum
}
