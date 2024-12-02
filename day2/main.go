package main

import (
  "bufio"
  "fmt"
  "math"
  "os"
  "strconv"
  "strings"
)

func main() {
  reports := readInput("input.txt")
  safeReports := 0
  for i := range reports {
    if isSafe(reports[i]) {
      safeReports += 1
    }
  }

  fmt.Println("Part 1 - Count of safe reports:", safeReports)

  safeReports = 0
  for i := range reports {
    if isSafePart2(reports[i]) {
      safeReports += 1
    }
  }

  fmt.Println("Part 2 - Count of safe reports:", safeReports)
}

func isSafe(report []int) bool {
  allIncreasing := true
  allDecreasing := true

  for i := 0; i < (len(report) -1); i += 1 {
    diff := report[i] - report[i+1]
    absDiff := int(math.Abs(float64(diff)))

    // If there is no change or the absolute value of the change is more than 3, we can immediately tell it's not safe
    if diff == 0 || absDiff > 3 {
      return false
    }

    // If the diff is positive, the values are decreasing
    // If the diff is negative, the values are increasing
    if diff > 0 {
      allIncreasing = false
    } else {
      allDecreasing = false
    }
  }

  return allIncreasing || allDecreasing
}

func isSafePart2(report []int) bool {
  if isSafe(report) {
    return true
  }

  // If the report is not already safe, we need to search more
  for i := range report {
    // Create a copy of the initial report as removing a reading will modify the original
    reportCopy := append([]int{}, report...)
    // Remove the reading at index i
    trimmedReport := append(reportCopy[:i], reportCopy[i+1:]...)

    // Check whether the report with the value removed is safe or not.
    // We can immediately return true if it's safe to prevent checking any more iterations
    if isSafe(trimmedReport) {
      return true
    }
  }

  // If it's gone through each possible iteration of the report with a reading missing without returning
  // by now, then we know that there were no possible iteration that was safe
  return false
}

func readInput(filename string) ([][]int) {
  file, err := os.Open(filename)
  if err != nil {
    panic(err)
  }
  defer file.Close()

  var reports [][]int

  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    line := scanner.Text()
    levels := strings.Fields(line)

    var report []int
    for i := range levels {
      level, _ := strconv.Atoi(levels[i])
      report = append(report, level)
    }
    reports = append(reports, report)
  }

  return reports
}
