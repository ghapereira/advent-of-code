package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"golang.org/x/exp/constraints"
)

type Coord struct {
	X int
	Y int
}

type AntennaPositions map[rune][]Coord

type Problem struct {
	AP        AntennaPositions
	Layout    []string
	Antinodes [][]rune
}

// GetMaxDimensions provides the dimensions of the map as a coordinate
func (p *Problem) GetMaxDimensions() Coord {
	if p.Layout == nil || len(p.Layout) <= 0 {
		return Coord{X: 0, Y: 0}
	}
	return Coord{X: len(p.Layout), Y: len(p.Layout[0])}
}

func (p *Problem) ShowAntinodes() {
	dims := p.GetMaxDimensions()
	for i := 0; i < dims.X; i++ {
		total := 0
		for j := 0; j < dims.Y; j++ {
			node := p.Antinodes[i][j]
			if node == '#' {
				total++
			}
			fmt.Printf("%c", node)
		}

		fmt.Println(" | ", total)
	}
}

func (p *Problem) IsNodeValid(c Coord) bool {
	maxDimensions := p.GetMaxDimensions()
	if (c.X < 0) || (c.X >= maxDimensions.X) {
		return false
	}
	if (c.Y < 0) || (c.Y >= maxDimensions.Y) {
		return false
	}

	return true
}

func main() {
	processFile("../inputs/8.input")
	// processFile("../inputs/8.dummy.input")
	// processFile("../inputs/8.dummy.2.input")
}

func processFile(fileName string) {
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	p := new(Problem)

	p.AP = make(AntennaPositions)
	p.Layout = make([]string, 0)

	scanner := bufio.NewScanner(f)
	i := 0
	for scanner.Scan() {
		line := scanner.Text()
		p.Layout = append(p.Layout, line)
		p.FillAntennaCoords(i, line)
		i += 1
	}

	totalAntinodes := p.SearchAntinodes()
	p.ShowAntinodes()

	fmt.Println(totalAntinodes)
}

func Abs[T constraints.Integer](x T) T {
	if x < 0 {
		return -x
	}
	return x
}

func (p *Problem) SetAntinodeSymbolAtPosition(r rune, c Coord) {
	p.Antinodes[c.X][c.Y] = r
}

func (p *Problem) SearchAntinodes() int {
	dims := p.GetMaxDimensions()
	p.Antinodes = make([][]rune, dims.X)
	for i := 0; i < dims.X; i++ {
		p.Antinodes[i] = make([]rune, dims.Y)
		for j := 0; j < dims.Y; j++ {
			p.Antinodes[i][j] = '.'
		}
	}

	fmt.Println(p.AP)

	for _, antennae := range p.AP {
		for i := 0; i < len(antennae); i++ {
			// For this algorithm to work I depend on the fact that I populate
			// the antennae list topwise, so it is naturally ordered w.r.t X;
			// if the need arises to have an unordered list then it needs to
			// be sorted in the crescent order of X
			cur := antennae[i]
			for j := i + 1; j < len(antennae); j++ {
				next := antennae[j]
				fmt.Printf("comparing %v to %v -> ", cur, next)

				isMainDiag := cur.Y < next.Y
				if isMainDiag {
					p.FillMainDiag(cur, next)
				} else {
					p.FillSecondaryDiag(cur, next)
				}
			}
		}
	}

	uniqueAntinodes := 0

	for i := 0; i < dims.X; i++ {
		for j := 0; j < dims.Y; j++ {
			if p.Antinodes[i][j] == '#' {
				uniqueAntinodes++
			}
		}
	}

	return uniqueAntinodes
}

func (p *Problem) FillMainDiag(cur Coord, next Coord) {
	// Get the 'box difference'
	boxDiff := Coord{X: next.X - cur.X, Y: next.Y - cur.Y}
	// Fill in from cur left-top
	for aux := cur; p.IsNodeValid(aux); {
		p.SetAntinodeSymbolAtPosition('#', aux)
		aux.X -= boxDiff.X
		aux.Y -= boxDiff.Y
	}

	// Fill in from next right-down
	for aux := next; p.IsNodeValid(aux); {
		p.SetAntinodeSymbolAtPosition('#', aux)
		aux.X += boxDiff.X
		aux.Y += boxDiff.Y
	}
}
func (p *Problem) FillSecondaryDiag(cur Coord, next Coord) {
	// Get the 'box difference'
	boxDiff := Coord{X: next.X - cur.X, Y: cur.Y - next.Y}
	// Fill in from cur right-top
	for aux := cur; p.IsNodeValid(aux); {
		p.SetAntinodeSymbolAtPosition('#', aux)
		aux.X -= boxDiff.X
		aux.Y += boxDiff.Y
	}
	// Fill in from next left-down
	for aux := cur; p.IsNodeValid(aux); {
		p.SetAntinodeSymbolAtPosition('#', aux)
		aux.X += boxDiff.X
		aux.Y -= boxDiff.Y
	}
}

func (p *Problem) FillAntennaCoords(i int, line string) {
	for j, v := range line {
		addAntenna := v != '.'
		if !addAntenna {
			continue
		}
		_, ok := p.AP[v]
		if !ok {
			p.AP[v] = make([]Coord, 0)
		}
		p.AP[v] = append(p.AP[v], Coord{X: i, Y: j})
	}
}
