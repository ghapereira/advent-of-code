package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
)

const initialWidthPlaceholder = -1

type WordSearchBox struct {
    WordMatrix []string
    Height int
    Width int
}

func main() {
    wsb := processFile("inputs/4.input")
    xmasCount := countXmas(wsb)
    fmt.Println(xmasCount)
}

func countXmas(wsb *WordSearchBox) int {
    counter := 0

    for i := range(wsb.Width) {
        for j := range(wsb.Height) {
            c := wsb.WordMatrix[i][j]
            if c == 'X' { counter += wsb.CheckXmas(i, j) }
        }
    }

    return counter
}

func (wsb *WordSearchBox) CheckXmas(i, j int) int {
    xmasCounter := 0

    if (i >= wsb.Height) || (j >= wsb.Width) || (i < 0) || (j < 0) {
        log.Fatal("invalid dimensions")
        os.Exit(-1)
    }

    /*
        8 1 2 
        7 C 3
        6 5 4

    C is the center point, the coordinates (i,j). We run clockwise from the
    cell marked '1' then looking for 'XMAS' on all available directions

    Then we check which directions we can look to:
    UP   : 8, 1, 2
    RIGHT: 2, 3, 4
    DOWN : 6, 5, 4
    LEFT : 8, 7, 6
    */

    canLookUp := (i - 3) >= 0
    canLookRight := (j + 3) < wsb.Width
    canLookDown := (i + 3) < wsb.Height
    canLookLeft := (j - 3) >= 0
    
    // TODO: simplify this with a loop in (-n,...,0,...,n) X (-n,...,0,...,n)
    //       and embedding bounds checks to simplify and generalize for search
    //       string size
    if canLookUp {
        // check 1
        if (wsb.WordMatrix[i-1][j] == 'M') && (wsb.WordMatrix[i-2][j] == 'A') && (wsb.WordMatrix[i-3][j] == 'S') {
            xmasCounter++            
        }
        if canLookRight {
            // check 2
            if (wsb.WordMatrix[i-1][j+1] == 'M') && (wsb.WordMatrix[i-2][j+2] == 'A') && (wsb.WordMatrix[i-3][j+3] == 'S') {
                xmasCounter++            
            }
        }
        if canLookLeft {
            // check 8
            if (wsb.WordMatrix[i-1][j-1] == 'M') && (wsb.WordMatrix[i-2][j-2] == 'A') && (wsb.WordMatrix[i-3][j-3] == 'S') {
                xmasCounter++            
            }
        }
    }
    if canLookRight {
        // check 3
        if (wsb.WordMatrix[i][j+1] == 'M') && (wsb.WordMatrix[i][j+2] == 'A') && (wsb.WordMatrix[i][j+3] == 'S') {
            xmasCounter++            
        }
        if canLookDown {
            // check 4
            if (wsb.WordMatrix[i+1][j+1] == 'M') && (wsb.WordMatrix[i+2][j+2] == 'A') && (wsb.WordMatrix[i+3][j+3] == 'S') {
                xmasCounter++            
            }
        }
    }

    if canLookDown {
        // check 5
        if (wsb.WordMatrix[i+1][j] == 'M') && (wsb.WordMatrix[i+2][j] == 'A') && (wsb.WordMatrix[i+3][j] == 'S') {
            xmasCounter++            
        }
        if canLookLeft {
            // check 6
            if (wsb.WordMatrix[i+1][j-1] == 'M') && (wsb.WordMatrix[i+2][j-2] == 'A') && (wsb.WordMatrix[i+3][j-3] == 'S') {
                xmasCounter++            
            }
        }
    }

    if canLookLeft {
        // check 7
        if (wsb.WordMatrix[i][j-1] == 'M') && (wsb.WordMatrix[i][j-2] == 'A') && (wsb.WordMatrix[i][j-3] == 'S') {
            xmasCounter++            
        }
    }
    
    return xmasCounter
}

func processFile(fileName string) *WordSearchBox {
    f, err := os.Open(fileName) 
    if err != nil {
        log.Fatal("couldn't open file")
        os.Exit(-1)
    }
    defer f.Close()

    wsb := new(WordSearchBox)
    wsb.WordMatrix = make([]string, 0)

    scanner := bufio.NewScanner(f)
    wsb.Width = initialWidthPlaceholder
    for scanner.Scan() {
        wordLine := scanner.Text()

        if wsb.Width == initialWidthPlaceholder {
            wsb.Width = len(wordLine)
        } else if wsb.Width != len(wordLine) {
            log.Fatal("invalid line length") 
            os.Exit(-1)
        }

        wsb.WordMatrix = append(wsb.WordMatrix, wordLine)
    }

    wsb.Height = len(wsb.WordMatrix)

    return wsb
}

