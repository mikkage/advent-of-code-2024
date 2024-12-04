package main

import (
  "bufio"
  "fmt"
  "os"
  "regexp"
  "strconv"
)

func main() {
  // Part 1
  matches := readInput("input.txt")
  total := 0

  // Each instruction will have the first and second digit stored in the second and third index
  // Multiply them and add it to the total
  for i := range matches {
    num1, _ := strconv.Atoi(matches[i][1])
    num2, _ := strconv.Atoi(matches[i][2])
    total += num1 * num2
  }

  fmt.Println("Part 1 total:", total)

  matchesP2 := readInputPart2("input.txt")
  total = 0
  r := regexp.MustCompile(`mul\((\d+),(\d+)\)`)

  // Instructions are enabled by default
  instructionsEnabled := true

  // Iterate through all instructions
  for i := range matchesP2 {
    // If the instruction is do() or don't(), enable or disable instructions accordingly then move on to the next iteration
    if matchesP2[i] == "don't()" || matchesP2[i] == "do()" {
      instructionsEnabled =matchesP2[i] == "do()"
      continue
    }

    // If the current instruction is not do() or don't(), then it's mul()
    // We only need to do the mul() operation if instructions are currently enabled
    if instructionsEnabled {
      // Extract the two mul() arguments from the instruction and do the operation
      regexMatches := r.FindAllStringSubmatch(matchesP2[i], -1)

      for i := range regexMatches {
        num1, _ := strconv.Atoi(regexMatches[i][1])
        num2, _ := strconv.Atoi(regexMatches[i][2])
        total += num1 * num2
      }
    }
  }

  fmt.Println("Part 2 total:", total)
}

// Part 1
// Reads the input file and finds all mul() instructions.
// Directly captures the arguments to each mul() instruction as part of the regex capture group result
func readInput(filename string) [][]string {
  file, err := os.Open(filename)
  if err != nil {
    panic(err)
  }
  defer file.Close()

  r := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
  var matches [][]string

  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    line := scanner.Text()
    regexMatches := r.FindAllStringSubmatch(line, -1)
    matches = append(matches, regexMatches...)
  }

  return matches
}

// Part 2
// Scans the input file for all mul(), do(), and don't() instructions
// Returns a list of each instruction in the order it appears in the file
func readInputPart2(filename string) []string {
  file, err := os.Open(filename)
  if err != nil {
    panic(err)
  }
  defer file.Close()

  r := regexp.MustCompile(`(mul\(\d+,\d+\)|do\(\)|don't\(\))`)
  var matches []string

  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    line := scanner.Text()
    regexMatches := r.FindAllString(line, -1)
    matches = append(matches, regexMatches...)
  }

  return matches
}
