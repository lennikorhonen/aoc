package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	input, err := os.ReadFile("input.txt")

	if err != nil {
		panic(err)
	}

	fmt.Println(dataReciever(string(input), 1))
	fmt.Println(dataReciever(string(input), 2))
}

func dataReciever(input string, partNum int) int {
	splitInput := strings.Split(string(input), "")

	var marker []string
	var lastChar int
	var incrementer int
	if partNum == 1 {
		incrementer = 4
	} else {
		incrementer = 14
	}
	for i := 0; i+incrementer <= len(splitInput); i++ {
		lastChar = 0
		var checkSlice []string
		checker := false
		marker = splitInput[i : i+incrementer]
		for j := 0; j < len(marker); j++ {
			if strings.Contains(strings.Join(checkSlice, ""), marker[j]) {
				checker = false
				break
			} else {
				checkSlice = append(checkSlice, marker[j])
				checker = true
			}
		}
		if checker {
			lastChar = i + incrementer
			break
		}
	}

	return lastChar
}
