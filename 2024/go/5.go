package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strconv"
    "strings"
    "regexp"
)

func main() {
    f, err := os.Open("inputs/3.input")
    // f, err := os.Open("inputs/dummy3.input")
    if err != nil {
        log.Fatal(err)
        os.Exit(-1)
    }
    
    defer f.Close()

    totalValue := 0

    instructions := getInstructions(f)
    for _, inst := range instructions {
       // fmt.Println(inst)
       components := strings.Split(inst, ",")

       // Not checking for error because we already matched an integer
       term1, _ := strconv.Atoi(components[0][4:])
       term2, _ := strconv.Atoi(components[1][:len(components[1])-1])
       totalValue += (term1 * term2)

       // fmt.Printf("\t%d * %d = %d - total %d\n", term1, term2, term1*term2, totalValue)
    }

    fmt.Println(totalValue)
}

func getInstructions(f *os.File) []string {
    instructions := make([]string, 0)

    scanner := bufio.NewScanner(f)
    for scanner.Scan() {
        memoryContents := scanner.Text()
        re := regexp.MustCompile(`mul\(\d+,\d+\)`)
        instructions = append(instructions, re.FindAllString(memoryContents, -1)...)
    }

    return instructions
}

