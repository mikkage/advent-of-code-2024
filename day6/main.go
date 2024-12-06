package main

import (
  "bufio"
  "fmt"
  "os"
)

func main() {
  areaMap := readInput("input.txt")
  guardX, guardY := findGuard(areaMap)
  guard := Guard{
    X: guardX,
    Y: guardY,
    XVelocity: 0,
    YVelocity: -1,
  }

  for {
    // Mark current spot at visited
    areaMap[guard.Y][guard.X] = 'X'

    // Check if we can move
    canMove, atEdge := canMove(areaMap, guard)
    if canMove && atEdge {
      break
    }

    if canMove {
      guard.X += guard.XVelocity
      guard.Y += guard.YVelocity
    } else {
      guard.XVelocity, guard.YVelocity = nextVelocities(guard.XVelocity, guard.YVelocity)
    }

  }

  visitedCount := 0
  for i := range areaMap {
    for j := range areaMap[i] {
      if areaMap[i][j] == 'X' {
        visitedCount += 1
      }
    }
  }


  fmt.Println(visitedCount)
}

type Guard struct {
  X int
  Y int
  XVelocity int
  YVelocity int
}

func readInput(filename string) ([][]byte) {
  file, err := os.Open(filename)
  if err != nil {
    panic(err)
  }
  defer file.Close()

  var m [][]byte

  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    line := scanner.Text()
    m = append(m, []byte(line))
  }


  return m
}


func nextVelocities(x, y int) (int, int) {
  // Down to left
  if x == 0 && y == 1 {
    return -1, 0
  }

  // Left to up
  if x == -1 && y == 0 {
    return 0, -1
  }

  // Up to right
  if x == 0 && y == -1 {
    return 1, 0
  }

  // Right to down
  return 0, 1
}

func findGuard(areaMap [][]byte) (int, int) {
  for y := range areaMap {
    for x := range areaMap[y] {
      if areaMap[y][x] == '^' {
        return x, y
      }
    }
  }

  // This can't happen with valid input
  return -1, -1
}

func canMove(areaMap [][]byte, guard Guard) (bool, bool) {
  if guard.XVelocity == -1 && guard.X == 0 {
    return true, true
  }
  if guard.XVelocity == 1 && guard.X == len(areaMap)-1 {
    return true, true
  }

  if guard.YVelocity == -1 && guard.Y == 0 {
    return true, true
  }
  if guard.YVelocity == 1 && guard.Y == len(areaMap)-1 {
    return true, true
  }

  if guard.XVelocity != 0 {
    if areaMap[guard.Y][guard.X + guard.XVelocity] != '#' {
      return true, false
    }
  }

  if guard.YVelocity != 0 {
    if areaMap[guard.Y + guard.YVelocity][guard.X] != '#' {
      return true, false
    }
  }

  return false, false
}
