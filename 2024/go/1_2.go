package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
)

func main() {
    inputName := "inputs/1.input"
    f, err := os.Open(inputName)
    if err != nil {
        log.Fatal(err)
        os.Exit(-1)
    }

    defer f.Close()

    leftList := make([]int, 0)
    rightList := make([]int, 0)

    scanner := bufio.NewScanner(f)

    var left int
    var right int

    for scanner.Scan() {
        fmt.Sscanf(scanner.Text(), "%d %d", &left, &right)
        leftList = append(leftList, left)
        rightList = append(rightList, right)
    }

    if len(leftList) != len(rightList) {
        log.Fatal("different list sizes!")
        os.Exit(-1)
    }

    // fmt.Println(diff)

    numCounter := make(map[int]int)
    for _, val := range leftList {
        numCounter[val] = 0
    }
    for _, val := range rightList {
        _, ok := numCounter[val]
        if ok {
            numCounter[val] += 1
        }
    }

    similarityScore := 0
    for k, v := range numCounter {
        similarityScore += (k * v)
    }

    fmt.Println(similarityScore)

}

