package main

import (
    "bufio"
    "fmt"
    "log"
    "math"
    "os"
    "strconv"
    "strings"
)

type Direction int
const (
    empty Direction = iota
    flat
    increasing
    decreasing
)

const diffThreshold = 3

func main() {
    inputName := "inputs/2.input"
    f, err := os.Open(inputName)
    if err != nil {
        log.Fatal(err)
        os.Exit(-1)
    }

    defer f.Close()

    scanner := bufio.NewScanner(f)

    safeLevels := 0
    for scanner.Scan() {
        report := scanner.Text()
        sLevels := strings.Split(report, " ")
        levels := make([]int, len(sLevels))
        for i, l := range sLevels {
            v, err := strconv.Atoi(l)
            if err != nil {
                log.Fatal(err)
                os.Exit(1)
            }
            levels[i] = v
        }
        isSafe := areLevelsSafe(levels)
        if isSafe {
            safeLevels++
        }
    }

    fmt.Println(safeLevels)
}

func areLevelsSafe(levels []int) bool {
    if len(levels) <= 1 {
        return true
    }

    initialDirection := determineDirection(levels[0], levels[1])
    if initialDirection == flat {
        return false
    }
    diff := int(math.Abs(float64(levels[0]) - float64(levels[1])))
    if diff > diffThreshold {
        return false
    }

    for i := 1; i < len(levels) - 1; i++ {
        currentDirection := determineDirection(levels[i], levels[i+1])
        if currentDirection != initialDirection {
            return false
        }

        diff = int(math.Abs(float64(levels[i]) - float64(levels[i+1])))
        if diff > diffThreshold {
            return false
        }
    }
        
    return true
}

func determineDirection(a, b int) Direction {
    if a == b {
        return flat
    }

    if a < b {
        return increasing
    } 

    return decreasing
}

