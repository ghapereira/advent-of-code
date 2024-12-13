package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"

	"golang.org/x/exp/constraints"
)

type Problem struct {
	BlockQuantity   int64
	Blocks          []string
	LastEmptyIndex  int64
	LastFilledIndex int64
}

func (p *Problem) Initialize(memory string) error {
	memorySize := len(memory)
	if (memorySize % 2) == 0 {
		return errors.New("invalid memory size")
	}

	p.BlockQuantity = 0
	for _, val := range memory {
		blockLen, err := strconv.Atoi(string(val))
		if err != nil {
			return errors.New("invalid block rune")
		}
		p.BlockQuantity += int64(blockLen)
	}
	p.Blocks = make([]string, p.BlockQuantity)

	var currentMemory int64 = 0
	currentBlock := 0
	for i := 0; i < memorySize; i += 2 {
		blockLen, err := strconv.Atoi(string(memory[i]))
		if err != nil {
			return errors.New("invalid block rune")
		}

		targetMem := currentMemory + int64(blockLen)
		for ; currentMemory < targetMem; currentMemory++ {
			p.Blocks[currentMemory] = fmt.Sprintf("%d", currentBlock)
		}

		isLastElement := i >= (memorySize - 1)
		if !isLastElement {
			spacesBlockLen, err := strconv.Atoi(string(memory[i+1]))
			if err != nil {
				return errors.New("invalid rune conversion")
			}
			targetMem := currentMemory + int64(spacesBlockLen)
			for ; currentMemory < targetMem; currentMemory++ {
				p.Blocks[currentMemory] = "."
				if p.LastEmptyIndex <= 0 {
					p.LastEmptyIndex = currentMemory
				}
			}
		}
		currentBlock++
	}

	return nil
}

func (p *Problem) Defrag() {
	p.LastFilledIndex = p.BlockQuantity - 1

	for p.LastFilledIndex > p.LastEmptyIndex {
		temp := p.Blocks[p.LastFilledIndex]
		p.Blocks[p.LastFilledIndex] = "."
		p.Blocks[p.LastEmptyIndex] = temp
		p.LastFilledIndex = p.GetPrevFilled()
		p.LastEmptyIndex = p.GetNextEmpty()
	}

	fmt.Println(p.Blocks)
}

func (p *Problem) GetPrevFilled() int64 {
	var i int64

	for i = p.LastFilledIndex; i > p.LastEmptyIndex; i-- {
		if p.Blocks[i] != "." {
			return i
		}
	}

	return i
}

func (p *Problem) GetNextEmpty() int64 {
	var i int64

	for i = p.LastEmptyIndex; i < p.LastFilledIndex; i++ {
		if p.Blocks[i] == "." {
			return i
		}
	}

	return i
}

func (p *Problem) Checksum() int64 {
	var checksum int64 = 0
	for i := 0; p.Blocks[i] != "."; i++ {
		currentBlock, _ := strconv.Atoi(p.Blocks[i])
		checksum += int64(i * currentBlock)
	}

	return checksum
}

func main() {
	// processFile("../inputs/9.input")
	processFile("../inputs/9.dummy.input")
}

func processFile(fileName string) {
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Scan()

	p := new(Problem)
	err = p.Initialize(scanner.Text())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(p.Blocks)

	p.Defrag()

	fmt.Println(p.Checksum())
}

func Abs[T constraints.Integer](x T) T {
	if x < 0 {
		return -x
	}
	return x
}
