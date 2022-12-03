package main

import (
  "bufio"
  "fmt"
  "os"
)

func main() {
  file, err := os.Open("input.txt")

  if err != nil {
    fmt.Println("Error occured")
  }
  defer file.Close()

  scanner := bufio.NewScanner(file)
  rounds := make([][]string, 0)
  round := make([]string, 0)

  for scanner.Scan() {
    s := scanner.Text()
    round = append(round, s)
    rounds = append(rounds, round)
    round = make([]string, 0)
  }

  points := 0
  pointsPart2 := 0
  for _, round := range rounds {
    for _, r := range round {
      first := r[0:1]
      last := r[2:3]
      roundRes := calculatePoints(first, last)
      roundResPart2 := calculatePointsPart2(first, last)
      points += roundRes
      pointsPart2 += roundResPart2
    }
  }

  fmt.Println(points)
  fmt.Println(pointsPart2)
}

// Rock = X = 1pt, Paper = Y = 2pt, Scissors = Z = 3pt
// Rock = A,       Paper = B,       Scissors = C
// Lost = 0pt, Draw = 3pt, Win = 6pt
func calculatePoints(first string, last string) int {
  result := 0
  if first == "A" && last == "X" {
    result = 3 + 1
  } else if first == "A" && last == "Y" {
    result = 6 + 2
  } else if first == "A" && last == "Z" {
    result = 0 + 3
  } else if first == "B" && last == "X" {
    result = 0 + 1
  } else if first == "B" && last == "Y" {
    result = 3 + 2
  } else if first == "B" && last == "Z" {
    result = 6 + 3
  } else if first == "C" && last == "X" {
    result = 6 + 1
  } else if first == "C" && last == "Y" {
    result = 0 + 2
  } else if first == "C" && last == "Z" {
    result = 3 + 3
  } else {
    result = 0
  }

  return result
}

// Last = X need to lose, Last = Y need to draw, Last = Z need to win
func calculatePointsPart2(first string, last string) int {
  result := 0

  if last == "X" {
    if first == "A" {
      result = 0 + 3
    } else if first == "B" {
      result = 0 + 1
    } else if first == "C" {
      result = 0 + 2
    }
  } else if last == "Y" {
    if first == "A" {
      result = 3 + 1
    } else if first == "B" {
      result = 3 + 2
    } else if first == "C" {
      result = 3 + 3
    }
  } else if last == "Z" {
    if first == "A" {
      result = 6 + 2
    } else if first == "B" {
      result = 6 + 3
    } else if first == "C" {
      result = 6 + 1
    }
  }

  return result
}
