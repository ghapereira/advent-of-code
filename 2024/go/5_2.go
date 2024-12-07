package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Set struct {
	contents map[int]int
}

func (s *Set) init() {
	s.contents = make(map[int]int)
}

func (s *Set) Add(i int) {
	if s.contents == nil {
		s.init()
	}
	if s.Has(i) {
		return
	}

	s.contents[i] = 0
}

func (s *Set) Has(i int) bool {
	_, ok := s.contents[i]
	return ok
}

type PrecedenceRules struct {
	contents map[int]*Set
}

func (pr *PrecedenceRules) init() {
	pr.contents = make(map[int]*Set)
}

func (pr *PrecedenceRules) Has(i int) bool {
	_, ok := pr.contents[i]
	return ok
}

func (pr *PrecedenceRules) Add(prev, succ int) {
	_, ok := pr.contents[succ]
	if !ok {
		pr.contents[succ] = new(Set)
	}

	pr.contents[succ].Add(prev)
}

func main() {
	middleSum := processFile("inputs/5.input")
	// middleSum := processFile("inputs/5.dummy.input")
	fmt.Println(middleSum)
}

func processFile(fileName string) int {
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
		os.Exit(-1)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	pr := new(PrecedenceRules)
	pr.init()

	middleSum := 0
	firstSection := true
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			firstSection = false
			continue
		}

		switch firstSection {
		case true:
			pr.AddOrderingRules(line)
		case false:
			v, fixed := pr.GetFixedUpdateResult(line)
			if fixed {
				middleSum += v
			}
		}
	}

	return middleSum
}

func (pr *PrecedenceRules) AddOrderingRules(line string) {
	rule := strings.Split(line, "|")
	prev, _ := strconv.Atoi(rule[0])
	succ, _ := strconv.Atoi(rule[1])

	pr.Add(prev, succ)
}

// GetFixedUpdateResult retrieves midpoint results from fixed updates
// Returns the midpoint if update had to be corrected (0 otherwise), and whether the update was fixed.
func (pr *PrecedenceRules) GetFixedUpdateResult(line string) (int, bool) {
	update := strings.Split(line, ",")

	hadToBeFixed := false

	lenUpdate := len(update)
	for i := 0; i < lenUpdate; i++ {
		value, _ := strconv.Atoi(update[i])
		if !pr.Has(value) {
			continue
		}
		for j := i + 1; j < lenUpdate; j++ {
			page, _ := strconv.Atoi(update[j])
			dependencyExists := pr.contents[value].Has(page)
			if dependencyExists {
				hadToBeFixed = true
				temp := update[j]
				update[j] = update[i]
				update[i] = temp

				value, _ = strconv.Atoi(update[i])
			}
		}
	}

	if !hadToBeFixed {
		return 0, false
	}

	oddNumberOfPages := (len(update) % 2) != 0
	if !oddNumberOfPages {
		log.Fatal("should have an odd number of pages!")
		os.Exit(-1)
	}

	midpoint, _ := strconv.Atoi(update[(len(update)-1)/2])
	return midpoint, true
}
