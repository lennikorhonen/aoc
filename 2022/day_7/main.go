package main

import (
	"fmt"
	"os"
	"path"
	"strings"
)

func main() {
	input, err := os.ReadFile("input.prod")

	if err != nil {
		fmt.Println(err)
	}

	filesys, cd := map[string]int{}, ""
	for _, s := range strings.Split(strings.TrimSpace(string(input)), "\n") {
		size := 0
		name := ""

		if strings.HasPrefix(s, "$ cd") {
			cd = path.Join(cd, strings.Fields(s)[2])
		} else if _, err := fmt.Sscanf(s, "%d %s", &size, &name); err == nil {
			for dir := cd; dir != "/"; dir = path.Dir(dir) {
				filesys[dir] += size
			}
			filesys["/"] += size
		}
	}

	fmt.Println("Part one: ", sumOfDirs(filesys))
	fmt.Println("Part one: ", dirToRemove(filesys))
}

func sumOfDirs(filesys map[string]int) int {
	maxsize := 100000
	dirsizes := make([]int, 0)

	for _, size := range filesys {
		if size <= maxsize {
			dirsizes = append(dirsizes, size)
		}
	}

	sum := 0
	for _, i := range dirsizes {
		sum += i
	}

	return sum
}

func dirToRemove(filesys map[string]int) int {
	delDirSize := 0
	systemSize := 70000000
	sizeNeededToUp := 30000000

	for _, dirSize := range filesys {
		if dirSize+systemSize-filesys["/"] > sizeNeededToUp {
			fmt.Println(dirSize)
		}
	}

	return delDirSize
}
