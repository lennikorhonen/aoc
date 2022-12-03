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

  i := 0
  current := 0
  biggest := 0

  for fileScanner.Scan() {
    if fileScanner.Text() == "" {
      fmt.Println("Elf number: ", i)
      if current > biggest {
        biggest = current
      }
      current = 0
      i++
    } else {
      if line, err := strconv.Atoi(fileScanner.Text()); err == nil {
        current += line
      }
    }  
  }
  fmt.Println(biggest)

  readFile.Close()
}
