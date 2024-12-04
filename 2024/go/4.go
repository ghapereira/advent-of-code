package main

import (
    "bufio"
    "fmt"
    "log"
    "math"
    "os"
    "strconv"
    "strings"
    "slices"
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

        // TODO: run async and compare performance/memory usage
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
        } else {
            isSafeWithRetry := retryLevelsSafety(levels)
            if isSafeWithRetry {
                safeLevels++
            }
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

func retryLevelsSafety(levels []int) bool {
    for i := 0; i < len(levels); i++ {
       retrySlice := getRetrySlice(levels, i) 
       // fmt.Printf("\t%v\n", retrySlice)
       if areLevelsSafe(retrySlice) {
            // fmt.Printf("%v is safe removing element %v\n", levels, i)
            return true
       }
    }

    return false
}

func getRetrySlice(levels []int, breakPoint int) []int {
    if breakPoint == 0 {
        return levels[1:]
    } 
    if breakPoint == (len(levels) - 1) {
        return levels[:breakPoint]
    }
    s1 := levels[:breakPoint]
    s2 := levels[breakPoint+1:]
    s3 := slices.Concat(nil, s1, s2)

    // TODO: compare s3 vs sx performance/memory usage (copy vs slices.concat)
    /*
    sc1 := make([]int, len(s1))
    sc2 := make([]int, len(s2))
    copy(sc1, s1)
    copy(sc2, s2)
    sx := append(sc1, sc2...)
    */

    return s3
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

