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
  fileScanner := bufio.NewScanner(readFile)

  fileScanner.Split(bufio.ScanLines)

  current := 0
  top_three := [3]int{0, 0, 0}

  for fileScanner.Scan() {
    if fileScanner.Text() == "" {
      for i, v:= range top_three {
        if current > v {
          top_three[i] = current
          break
        }
      } 
      current = 0
    } else {
      if line, err := strconv.Atoi(fileScanner.Text()); err == nil {
        current += line
      }
    }  
  }

  sum_of_top_three := 0
  fmt.Println(top_three)

  for _, v := range top_three {
    sum_of_top_three += v
  }
  fmt.Println(sum_of_top_three)

  readFile.Close()
}
