package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	processFile("../inputs/7.input")
	// processFile("../inputs/7.dummy.input")
}

type Operand string

const (
	Add      Operand = "+"
	Multiply Operand = "*"
)

type Equation struct {
	Result int
	Values []int
}

type Node struct {
	val int
	Add *Node
	Mul *Node
}

func (n *Node) InsertVal(val int, result int, found *bool) {
	if n.Add == nil {
		n.Add = &Node{val: n.val + val}
		if n.Add.val == result {
			*found = true
			return
		}
	} else {
		n.Add.InsertVal(val, result, found)
	}

	if n.Mul == nil {
		n.Mul = &Node{val: n.val * val}
		if n.Mul.val == result {
			*found = true
			return
		}
	} else {
		n.Mul.InsertVal(val, result, found)
	}
}

// Solve resolves an equation.
// If the equation is unsolvable, return 0. Otherwise, return its value.
func (e *Equation) Solve() int {
	var n *Node = new(Node)
	n.val = e.Values[0]
	resultFound := false

	for i := 1; i < len(e.Values); i++ {
		n.InsertVal(e.Values[i], e.Result, &resultFound)
		if resultFound {
			return e.Result
		}
	}

	return 0
}

func processFile(fileName string) {
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	sum := 0
	i := 0
	for scanner.Scan() {
		fmt.Println("processing line ", i, "...")
		i++
		line := scanner.Text()
		equation := buildEquation(line)
		sum += equation.Solve()
	}

	fmt.Println(sum)
}

func (e *Equation) Show() {
	fmt.Printf("%d: %v\n", e.Result, e.Values)
}

func buildEquation(line string) *Equation {
	equation := new(Equation)

	formula := strings.Split(line, ":")
	equation.Result, _ = strconv.Atoi(formula[0])
	values := strings.Split(strings.TrimLeft(formula[1], " "), " ")

	equation.Values = make([]int, len(values))
	for i, v := range values {
		c, _ := strconv.Atoi(v)
		equation.Values[i] = c
	}

	return equation
}
