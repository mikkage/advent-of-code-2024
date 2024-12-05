package main

import (
  "bufio"
  "fmt"
  "os"
  "strings"
  "strconv"
)

func main() {
  rules, updates := readInput("input.txt")

  // Keep track of which updates are and are not in order
  var updatesInOrder [][]int
  var updatesNotInOrder [][]int

  // Go through each update, check if meets the rules, and add it to the appropriate list whether it's in order or not
  for i := range updates {
    if isInOrder(updates[i], rules) {
      updatesInOrder = append(updatesInOrder, updates[i])
    } else {
      updatesNotInOrder = append(updatesNotInOrder, updates[i])
    }
  }

  // Part 1 - Go through each update that is in order, find the middle element and sum all the values
  total := 0
  for i := range updatesInOrder {
    total += getMiddleElement(updatesInOrder[i])
  }

  fmt.Println("Part 1 total:", total)

  // Part 2
  // Go through each out of order update and apply the rules to them until they are in order
  for i := range updatesNotInOrder {
    inOrder := false

    // Brute force - keep applying all of the rules to the update until it's in order
    for {
      if inOrder {
        break
      }

      for j := range rules {
        updatesNotInOrder[i] = applyRule(updatesNotInOrder[i], rules[j])
      }
      inOrder = isInOrder(updatesNotInOrder[i], rules)
    }
  }

  total = 0
  // Now that each out of order update is in order, we just have to find the middle element of each update and sum the values
  for i := range updatesNotInOrder {
    total += getMiddleElement(updatesNotInOrder[i])
  }

  fmt.Println("Part 2 total:", total)
}

type OrderingRule struct {
  X int
  Y int
}

func readInput(filename string) ([]OrderingRule, [][]int) {
  file, err := os.Open(filename)
  if err != nil {
    panic(err)
  }
  defer file.Close()

  var rules []OrderingRule
  var updates [][]int

  isUpdate := false
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    line := scanner.Text()

    if line == "" {
      isUpdate = true
      continue
    }

    if isUpdate {
      pages := strings.Split(line, ",")
      var pageNumbers []int
      for i := range pages {
        number, _ := strconv.Atoi(pages[i])
        pageNumbers = append(pageNumbers, number)
      }
      updates = append(updates, pageNumbers)
    } else {
      r := strings.Split(line, "|")
      x, _ := strconv.Atoi(r[0])
      y, _ := strconv.Atoi(r[1])
      rule := OrderingRule{
        X: x,
        Y: y,
      }

      rules = append(rules, rule)
    }
  }

  return rules, updates
}

// Returns the index of the given value in the given array. Returns -1 if the value is not in the array
func getIndexOf(list []int, value int) int {
  for i := range list {
    if list[i] == value {
      return i
    }
  }

  return -1
}

// Returns the value of the middle element in the array. Note that this problem only uses array os odd lengths
func getMiddleElement(list []int) int {
  l := len(list)
  return list[l/2]
}

// Returns whether or not the given update is in order given a set of rules
func isInOrder(update []int, rules[]OrderingRule) bool {
  for i := range rules {
    xIndex := getIndexOf(update, rules[i].X)
    yIndex := getIndexOf(update, rules[i].Y)

    // If both elements exist in the update and the X element is after the Y element(X comes after Y)
    // then we know that this update does not follow at least one of the rules and is not in order
    if xIndex >= 0 && yIndex >= 0 && xIndex > yIndex {
      return false
    }
  }

  // If none of the rules were violated by the update, then it's in order
  return true
}

// Applies the given rule to the give update
// In this case, applying the update means that for any two values that do not meet the rules, the two elements get swapped
func applyRule(update []int, rule OrderingRule) []int {
  xIndex := getIndexOf(update, rule.X)
  yIndex := getIndexOf(update, rule.Y)

  // If both elements exist in the update and do not follow the rule, swap the two elements
  if xIndex >= 0 && yIndex >= 0 && xIndex > yIndex {
    update[xIndex], update[yIndex] = update[yIndex], update[xIndex]
  }

  return update
}
