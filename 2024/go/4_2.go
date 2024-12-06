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
    x_masCount := countX_Mas(wsb)
    fmt.Println(x_masCount)
}

func countX_Mas(wsb *WordSearchBox) int {
    counter := 0

    for i := 1; i < wsb.Width - 1; i++ {
        for j := 1; j < wsb.Height - 1; j++ {
            c := wsb.WordMatrix[i][j]
            if c == 'A' && wsb.CheckX_Mas(i, j) { counter++ }
        }
    }

    return counter
}

func (wsb *WordSearchBox) CheckX_Mas(i, j int) bool {
    // main diagonal
    mainDiagIsMAS := wsb.WordMatrix[i-1][j-1] == 'M' && wsb.WordMatrix[i+1][j+1] == 'S'
    mainDiagIsSAM := wsb.WordMatrix[i-1][j-1] == 'S' && wsb.WordMatrix[i+1][j+1] == 'M'
    mainDiagOk := mainDiagIsMAS || mainDiagIsSAM

    // secondary diagonal
    secDiagIsMAS := wsb.WordMatrix[i-1][j+1] == 'M' && wsb.WordMatrix[i+1][j-1] == 'S'
    secDiagIsSAM := wsb.WordMatrix[i-1][j+1] == 'S' && wsb.WordMatrix[i+1][j-1] == 'M'
    secDiagOk := secDiagIsMAS || secDiagIsSAM

    return mainDiagOk && secDiagOk
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

