package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(topOfStack(string(file)))
	fmt.Println(topOfStack9001(string(file)))
}

func topOfStack(input string) string {
	inputLines := strings.Split(input, "\n")
	stacks := [][]byte{
		[]byte("FRW"),
		[]byte("PWVDCMHT"),
		[]byte("LNZMP"),
		[]byte("RHCJ"),
		[]byte("BTQHGPC"),
		[]byte("ZFLWCG"),
		[]byte("CGJZQLVW"),
		[]byte("CVTWFRNP"),
		[]byte("VSRGHWJ"),
	}

	for i := 0; i < 10; i++ {
		stacks = append(stacks, []byte{})
	}

	digitsRegex := regexp.MustCompile(`\d+`)

	instructionsIndex := 0

	for i, line := range inputLines {
		if digitsRegex.MatchString(line) {
			instructionsIndex = i + 2
			break
		}

		for s, c := 0, 1; c < len(line); c, s = c+4, s+1 {
			if line[c] == ' ' {
				continue
			}
			stacks[s] = append(stacks[s], line[c])
		}
	}

	for i := instructionsIndex; i < len(inputLines); i++ {
		if inputLines[i] == "" {
			break
		}

		instructions := digitsRegex.FindAllString(inputLines[i], 3)
		n, _ := strconv.Atoi(instructions[0])
		from, _ := strconv.Atoi(instructions[1])
		dest, _ := strconv.Atoi(instructions[2])

		for ; n > 0; n-- {
			stacks[dest-1] = append([]byte{stacks[from-1][0]}, stacks[dest-1]...)
			stacks[from-1] = stacks[from-1][1:]
		}
	}

	result := []byte{}

	for _, stack := range stacks {
		if len(stack) == 0 {
			break
		}

		result = append(result, stack[0])
	}
	return string(result)
}

// Does not work with my input and not sure why. Also doesn't work with example input
// Had to read and copy other peoples
// code to even get the first answer
func topOfStack9001(input string) string {
	inputLines := strings.Split(input, "\n")
	stacks := [][]byte{
		[]byte("FRW"),
		[]byte("PWVDCMHT"),
		[]byte("LNZMP"),
		[]byte("RHCJ"),
		[]byte("BTQHGPC"),
		[]byte("ZFLWCG"),
		[]byte("CGJZQLVW"),
		[]byte("CVTWFRNP"),
		[]byte("VSRGHWJ"),
	}

	for i := 0; i < 10; i++ {
		stacks = append(stacks, []byte{})
	}

	digitsRegex := regexp.MustCompile(`\d+`)

	instructionsIndex := 0

	for i, line := range inputLines {
		if digitsRegex.MatchString(line) {
			instructionsIndex = i + 2
			break
		}

		for s, c := 0, 1; c < len(line); c, s = c+4, s+1 {
			if line[c] == ' ' {
				continue
			}
			stacks[s] = append([]byte{line[c]}, stacks[s]...)
		}
	}

	for i := instructionsIndex; i < len(inputLines); i += 1 {
		if inputLines[i] == "" {
			break
		}

		instructions := digitsRegex.FindAllString(inputLines[i], 3)
		n, _ := strconv.Atoi(instructions[0])
		from, _ := strconv.Atoi(instructions[1])
		dest, _ := strconv.Atoi(instructions[2])

		fromStack := stacks[from-1]
		stacks[dest-1] = append(stacks[dest-1], fromStack[len(fromStack)-n:]...)
		stacks[from-1] = stacks[from-1][:len(fromStack)-n]
	}

	result := []byte{}
	fmt.Println(stacks)

	for _, s := range stacks {
		if len(s) == 0 {
			break
		}

		result = append(result, s[len(s)-1])
	}
	return string(result)
}
