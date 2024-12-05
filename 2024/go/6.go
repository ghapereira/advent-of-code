package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "slices"
    "strconv"
    "strings"
    "regexp"
)

type Position struct {
    Start int
    End int
}

type TokenType int
const (
    TDo TokenType = iota
    TDont
    TMul
)

type Token struct {
    P Position
    T TokenType
}

type InstructionCollection struct {
    DoInstructions []Position
    DontInstructions []Position
    MulInstructions []Position
}

func main() {
    f, err := os.Open("inputs/3.input")
    // f, err := os.Open("inputs/dummy3.1.input")
    if err != nil {
        log.Fatal(err)
        os.Exit(-1)
    }
    
    defer f.Close()


    scanner := bufio.NewScanner(f)
    memoryContents := ""
    for scanner.Scan() {
        memoryContents += scanner.Text()
    }

    tokens := tokenizeInstructions(memoryContents)
    totalValue := runAccounting(tokens, memoryContents)

    fmt.Println(totalValue)
}

func runAccounting(tokens []Token, memoryContent string) int {
    total := 0

    shouldMul := true
    for _, t := range tokens {
        switch t.T {
            case TDont:
                shouldMul = false
            case TDo:
                shouldMul = true
            case TMul:
                if shouldMul {
                    total += mul(t, memoryContent[t.P.Start:t.P.End])        
                } 
            default:
                log.Fatal("invalid type")
                os.Exit(-1)
        }
    }

    return total
}

func mul(t Token, inst string) int {
   components := strings.Split(inst, ",")

   term1, _ := strconv.Atoi(components[0][4:])
   term2, _ := strconv.Atoi(components[1][:len(components[1])-1])
   return term1 * term2
}


func tokenizeInstructions(memoryContents string) []Token {
    insts := new(InstructionCollection)
    insts.MulInstructions = fillPositionSliceOnRegexAllIndexes(memoryContents, `mul\(\d+,\d+\)`)
    insts.DoInstructions = fillPositionSliceOnRegexAllIndexes(memoryContents, `do\(\)`)
    insts.DontInstructions = fillPositionSliceOnRegexAllIndexes(memoryContents, `don't\(\)`)

    tokens := make([]Token, len(insts.MulInstructions) + len(insts.DoInstructions) + len(insts.DontInstructions))

    cursor := 0
    for _, inst := range insts.MulInstructions {
        t := Token{P: inst, T: TMul,}

        tokens[cursor] = t
        cursor++
    }

    for _, inst := range insts.DoInstructions {
        t := Token{P: inst, T: TDo,}

        tokens[cursor] = t
        cursor++
    }

    for _, inst := range insts.DontInstructions {
        t := Token{P: inst, T: TDont,}

        tokens[cursor] = t
        cursor++
    }

    slices.SortFunc(tokens, func(i, j Token) int {
        if i.P.Start > j.P.Start { return 1 }
        if i.P.Start < j.P.Start { return -1 }
        log.Fatal("shouldn't have duplicated tokens")
        os.Exit(-1)
        return 0
    })

    return tokens
}

func fillPositionSliceOnRegexAllIndexes(baseContents string, regex string) []Position {
        targetSlice := make([]Position, 0)

        re := regexp.MustCompile(regex)
        indexes := re.FindAllIndex([]byte(baseContents), -1)
        for _, pair := range indexes {
            p := Position{Start: pair[0], End: pair[1]}
            targetSlice = append(targetSlice, p)
        }

        // TODO work directly on preallocated position slice pointer?
        return targetSlice
}

func showInstructions(memoryContents string, instructions *InstructionCollection) {
    fmt.Println("mul:")
    for _, p := range instructions.MulInstructions {
       fmt.Printf("\t%s\n", memoryContents[p.Start:p.End]) 
    }

    fmt.Println("do:")
    for _, p := range instructions.DoInstructions {
       fmt.Printf("\t%s\n", memoryContents[p.Start:p.End]) 
    }

    fmt.Println("don't:")
    for _, p := range instructions.DontInstructions {
       fmt.Printf("\t%s\n", memoryContents[p.Start:p.End]) 
    }
}


