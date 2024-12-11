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
			cur := antennae[i]
			for j := i + 1; j < len(antennae); j++ {
				next := antennae[j]
				fmt.Printf("comparing %v to %v -> ", cur, next)
				diffX := Abs(cur.X - next.X)
				diffY := Abs(cur.Y - next.Y)

				maxX := max(cur.X, next.X)
				minX := min(cur.X, next.X)
				maxY := max(cur.Y, next.Y)
				minY := min(cur.Y, next.Y)

				minAntinode := Coord{X: minX - diffX, Y: minY - diffY}
				fmt.Printf("%v, ", minAntinode)
				maxAntinode := Coord{X: maxX + diffX, Y: maxY + diffY}
				fmt.Printf("%v\n", maxAntinode)

				// Need to invert the logic: if leftwise, revert
				if cur.Y > next.Y {
					minAntinode.Y = maxY + diffY
					maxAntinode.Y = minY - diffY
				}
				if p.IsNodeValid(minAntinode) {
					p.Antinodes[minAntinode.X][minAntinode.Y] = '#'
				}
				if p.IsNodeValid(maxAntinode) {
					p.Antinodes[maxAntinode.X][maxAntinode.Y] = '#'
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
