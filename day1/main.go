package main

import (
  "bufio"
  "fmt"
  "math"
  "os"
  "sort"
  "strconv"
  "strings"
)

func main() {
  l1, l2 := readInput("input.txt")

  sort.Ints(l1)
  sort.Ints(l2)

  // Part 1
  totalDistance := 0

  for i := range l1 {
    difference := float64(l1[i] - l2[i])
    totalDistance += int(math.Abs(difference))
  }

  fmt.Println(totalDistance)

  // Part 2
  similarity := 0

  for i := range l1 {
    count := 0
    for j := range l2 {
      if l1[i] == l2[j] {
        count += 1
      }
    }

    similarity += l1[i] * count
  }

  fmt.Println(similarity)
}

func readInput(filename string) ([]int, []int) {
  file, err := os.Open(filename)
  if err != nil {
    panic(err)
  }
  defer file.Close()

  var list1 []int
  var list2 []int

  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    line := scanner.Text()
    ids := strings.Fields(line)
    id1, _ := strconv.Atoi(ids[0])
    id2, _ := strconv.Atoi(ids[1])
    list1 = append(list1, id1)
    list2 = append(list2, id2)
  }

  return list1, list2
}
