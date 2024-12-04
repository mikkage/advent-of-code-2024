package main

import (
  "bufio"
  "fmt"
  "os"
)

func main() {
  // Part 1
  matrix := readInput("input.txt")

  total := 0
  for y := range matrix {
    for x := range matrix[y] {
      total += search(matrix, x, y)
    }
  }
  fmt.Println("Part 1 total:", total)

  // Part 2
  total = 0
  for y := range matrix {
    for x := range matrix[y] {
      if searchP2(matrix, x, y) {
        total += 1
      }
    }
  }
  fmt.Println("Part 2 total:", total)
}

// Part 1
// Searches in all possible directions from the given coordinates and returns the number of times 'XMAS' was found
func search(matrix []string, x, y int) int {
  // Determine which directions we can search from the given coordinates
  // A direction is searchable if the coordinate(s) are at least 3 other characters between the current location and the edge of the array(s)
  canSearchUp := y >= 3
  canSearchDown := y <= (len(matrix) - 4)
  canSearchLeft := x >= 3
  canSearchRight := x <= (len(matrix[0]) - 4)

  words := []string{}

  // For each searchable direction, create a string with the four letters when searching in that direction and add it to a list
  if canSearchUp {
    word := ""
    for i := 0; i < 4; i += 1 {
      word += string(matrix[y-i][x])
    }
    words = append(words, word)
  }

  if canSearchDown {
    word := ""
    for i := 0; i < 4; i += 1 {
      word += string(matrix[y+i][x])
    }
    words = append(words, word)
  }

  if canSearchLeft {
    word := ""
    for i := 0; i < 4; i += 1 {
      word += string(matrix[y][x-i])
    }
    words = append(words, word)
  }

  if canSearchRight {
    word := ""
    for i := 0; i < 4; i += 1 {
      word += string(matrix[y][x+i])
    }
    words = append(words, word)
  }

  if canSearchUp && canSearchLeft {
    word := ""
    for i := 0; i < 4; i += 1 {
      word += string(matrix[y-i][x-i])
    }
    words = append(words, word)
  }

  if canSearchUp && canSearchRight {
    word := ""
    for i := 0; i < 4; i += 1 {
      word += string(matrix[y-i][x+i])
    }
    words = append(words, word)
  }

  if canSearchDown && canSearchLeft {
    word := ""
    for i := 0; i < 4; i += 1 {
      word += string(matrix[y+i][x-i])
    }
    words = append(words, word)
  }

  if canSearchDown && canSearchRight {
    word := ""
    for i := 0; i < 4; i += 1 {
      word += string(matrix[y+i][x+i])
    }
    words = append(words, word)
  }

  // Once all eight directions have been searched, we just need to see how many times 'XMAS' occurs in the list and return the count
  matches := 0
  for i := range words {
    if words[i] == "XMAS" {
      matches += 1
    }
  }

  return matches
}

// Part 2
// Searches around the given coordinates and returns whether or not an 'X-MAS' occurs at the given coordinates
func searchP2(matrix []string, x, y int) bool {
  matrixLength := len(matrix)

  // If the x or y coordinate is the first or last index, or the current letter is not an 'A', then this spot cannot be an X-MAS
  if x == 0 || y == 0 || x == matrixLength-1 || y == matrixLength-1 || string(matrix[y][x]) != "A" {
    return false
  }

  // Get each diagonal character from the center 'A'
  topLeft := string(matrix[y-1][x-1])
  topRight := string(matrix[y-1][x+1])
  bottomLeft := string(matrix[y+1][x-1])
  bottomRight := string(matrix[y+1][x+1])

  // Track whether each diagonal line spells out 'MAS'
  diagDownRightIsMAS := false
  diagUpLeftIsMAS := false

  // If the top left and bottom right diagonal characters are 'M' or 'S', and they are not equal, then 'MAS' is spelled in the diagonal
  if (topLeft == "M" || topLeft == "S") && (bottomRight == "M" || bottomRight == "S") {
    if topLeft != bottomRight {
      diagDownRightIsMAS = true
    }
  }
  // If the bottom left and top right diagonal characters are 'M' or 'S', and they are not equal, then 'MAS' is spelled in the diagonal
  if (bottomLeft == "M" || bottomLeft == "S") && (topRight == "M" || topRight == "S") {
    if bottomLeft != topRight {
      diagUpLeftIsMAS = true
    }
  }

  // If each diagonal line spells 'MAS' then the search in this location is an 'X-MAS'
  return diagDownRightIsMAS && diagUpLeftIsMAS
}

func readInput(filename string) []string {
  file, err := os.Open(filename)
  if err != nil {
    panic(err)
  }
  defer file.Close()

  var lines []string

  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    line := scanner.Text()
    lines = append(lines, line)
  }

  return lines
}
