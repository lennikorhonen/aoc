package main

import (
	"bufio"
	"fmt"
	"os"
  "strconv"
)

func main() {
  readFile, err := os.Open("input.txt")

  if err != nil {
    fmt.Println(err)
  }
  defer readFile.Close()

  fileScanner := bufio.NewScanner(readFile)
  elves := make([][]int, 0)
  elf := make([]int, 0)

  current := 0
  biggest := 0

  for fileScanner.Scan() {
    fs := fileScanner.Text()
    if fs == "" {
      elves = append(elves, elf)
      elf = make([]int, 0)
      if current > biggest {
        biggest = current
      }
      current = 0
    } else {
      i, err := strconv.Atoi(fs)
      if err != nil {
        fmt.Println(err)
      }
      elf = append(elf, i)
      if line, err := strconv.Atoi(fs); err == nil {
        current += line
      }
    }  
  }

  fmt.Println("Most cals:", mostCals(elves))
  fmt.Println("Sum of top three:", topThree(elves))
}

func mostCals(elves [][]int) int {
  biggest := 0

  for _, elf := range elves {
    sum := 0
    for _, cal := range elf {
      sum += cal
    }
    if sum > biggest {
      biggest = sum
    }
  }
  return biggest
}

func topThree(elves [][]int) int {
  top_three := [3]int{0, 0, 0}
  for _, elf := range elves {
    sum := 0
    for _, cal := range elf {
      sum += cal
    }
    x := sum
    for i, n := range top_three {
      if x > top_three[i] {
        top_three[i] = x
        x = n
        continue
      }
    }
  }
  sum := 0
  for _, n := range top_three {
    sum += n
  }
  return sum
}
