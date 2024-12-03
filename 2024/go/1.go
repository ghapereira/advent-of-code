package main

import (
    "bufio"
    "fmt"
    "log"
    "math"
    "os"
    "sort"
)

func main() {
    inputName := "1.input"
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

    sort.Slice(leftList, func (i, j int) bool {
        return leftList[i] < leftList[j]
    })
    sort.Slice(rightList, func (i, j int) bool {
        return rightList[i] < rightList[j]
    })

    diff := 0
    for i := range(len(leftList)) {
        diff += int(math.Abs(float64(leftList[i]) - float64(rightList[i])))
    }

    fmt.Println(diff)
}

