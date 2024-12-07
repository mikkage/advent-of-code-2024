package main

import (
  "bufio"
  "fmt"
  "math"
  "os"
  "strings"
  "strconv"
)

func main() {
  equations := readInput("input.txt")
  sum := 0
  for i := range equations {
    if equations[i].IsSolvable() {
      sum += equations[i].Result
    }
  }

  fmt.Println("Sum of solvable equations:", sum)
}

type Equation struct {
  Numbers []int
  Result int
}

func (e *Equation) IsSolvable() bool {
  // The number of operators for the equation is one less than the total count of numbers
  operatorCount := len(e.Numbers) - 1
  // max stores the total number of iterations to try every combination of + and * for the given number
  // of operators. For example:
  // 1 2 3 4
  //  + + +
  //  + + *
  //  + * +
  // .......
  // There will be 2^(operator count) combinations
  max := math.Pow(2, float64(operatorCount))

  // Try each combination of operators
  for i := 0; i < int(max); i += 1 {
    result := -999 // Starting off with dummy value

    // Covert the current iteration into the string representation of its binary value
    operatorsBinary := fmt.Sprintf("%0" + strconv.Itoa(operatorCount) +"b", i)

    // Using the 0 or 1 of the binary, apply the additon or multiplication respectively
    for j := range operatorsBinary {
      if operatorsBinary[j] == '0' {
        // For the first operation, we need to apply the operator to the first two numbers
        if j == 0 {
          result = e.Numbers[j] + e.Numbers[j+1]
          continue
        }

        // Otherwise, apply the operator the current result and the next number in the list
        result += e.Numbers[j+1]
      } else {
        if j == 0 {
          result = e.Numbers[j] * e.Numbers[j+1]
          continue
        }

        result *= e.Numbers[j+1]
      }
    }

    // We should return true as soon as we find the first combination of operators that works
    if e.Result == result {
      return true
    }
  }

  // If we've reached this point, no combinations of operators have ended up with the correct result
  return false
}

func readInput(filename string) ([]Equation) {
  file, err := os.Open(filename)
  if err != nil {
    panic(err)
  }
  defer file.Close()

  var equations []Equation

  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    line := scanner.Text()
    var e Equation

    eqSplit := strings.Split(line, ":")
    e.Result, _ = strconv.Atoi(eqSplit[0])
    numbers := strings.Split(strings.TrimSpace(eqSplit[1]), " ")
    for i := range numbers {
      n, _ := strconv.Atoi(numbers[i])
      e.Numbers = append(e.Numbers, n)
    }
    equations = append(equations, e)
  }

  return equations
}
