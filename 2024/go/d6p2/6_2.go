package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Coordinate struct {
	X int
	Y int
}

type MapSymbol rune

const (
	Empty     MapSymbol = '.'
	VrVisited MapSymbol = '|'
	HrVisited MapSymbol = '-'
	BdVisited MapSymbol = '+'
	Blocked   MapSymbol = '#'
)

func AlreadyVisited(m MapSymbol) bool {
	return m == VrVisited || m == HrVisited || m == BdVisited
}

type Orientation string

const (
	Up      Orientation = "^"
	Right   Orientation = ">"
	Down    Orientation = "v"
	Left    Orientation = "<"
	Invalid Orientation = "?"
)

func (l *LabMap) TurnGuard() {
	l.VisitCurrentPosition()

	switch l.GuardOrientation {
	case Up:
		l.GuardOrientation = Right
	case Right:
		l.GuardOrientation = Down
	case Down:
		l.GuardOrientation = Left
	case Left:
		l.GuardOrientation = Up
	default:
		log.Fatal("invalid direction to turn")
	}
}

func isGuard(s string) bool {
	switch s {
	case string(Up):
		return true
	case string(Right):
		return true
	case string(Down):
		return true
	case string(Left):
		return true
	default:
		return false
	}
}

type LabMap struct {
	Map               [][]rune
	Width             int
	Height            int
	GuardPos          Coordinate
	GuardOrientation  Orientation
	UniqueWalkedCells int
	StepsDone         int
}

func (l *LabMap) InfiniteLoopDetected() bool {
	const possibleDirections = 4
	maxPossibleMovements := possibleDirections * l.Height * l.Width
	return l.StepsDone > maxPossibleMovements
}

func main() {
	// l := processFile("../inputs/6.input")
	l := processFile("../inputs/6.dummy.input")
	c := l.ComputeGuardWalk()
	fmt.Println(c)
}

func processFile(fileName string) *LabMap {
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	l := new(LabMap)

	scanner := bufio.NewScanner(f)

	width := 0
	i := 0
	for scanner.Scan() {
		row := scanner.Text()
		rowWidth := len(row)

		if width == 0 {
			width = rowWidth
		} else {
			if width != rowWidth {
				log.Fatal("invalid width")
			}
		}

		mapRow := make([]rune, rowWidth)
		for j := 0; j < rowWidth; j++ {
			mapRow[j] = rune(row[j])
			if isGuard(string(row[j])) {
				l.GuardPos = Coordinate{X: i, Y: j}
				l.GuardOrientation = Orientation(row[j])
			}
		}

		l.Map = append(l.Map, mapRow)
		i++
	}

	l.Width = width
	l.Height = len(l.Map)

	return l
}

func isOrientation(r rune) bool {
	if r == rune(string(Up)[0]) {
		return true
	}
	if r == rune(string(Down)[0]) {
		return true
	}
	if r == rune(string(Left)[0]) {
		return true
	}
	if r == rune(string(Right)[0]) {
		return true
	}

	return false
}

func GetNextPositionSymbol(currentPosition rune, currentOrientation Orientation) rune {
	if currentPosition == rune(Empty) || isOrientation(currentPosition) {
		if currentOrientation == Up || currentOrientation == Down {
			return rune(VrVisited)
		}

		if currentOrientation == Left || currentOrientation == Right {
			return rune(HrVisited)
		}
	}

	if currentPosition == rune(VrVisited) {
		if currentOrientation == Up || currentOrientation == Down {
			return rune(VrVisited)
		}

		if currentOrientation == Left || currentOrientation == Right {
			return rune(BdVisited)
		}
	}

	if currentPosition == rune(HrVisited) {
		if currentOrientation == Up || currentOrientation == Down {
			return rune(BdVisited)
		}

		if currentOrientation == Left || currentOrientation == Right {
			return rune(HrVisited)
		}
	}

	log.Fatal("invalid current position")
	return rune(Empty)
}

func (l *LabMap) VisitCurrentPosition() {
	l.StepsDone++
	if l.InfiniteLoopDetected() {
		log.Fatal("infinite loop detected")
	}

	currentPositionSymbol := l.Map[l.GuardPos.X][l.GuardPos.Y]
	visitSymbol := GetNextPositionSymbol(currentPositionSymbol, l.GuardOrientation)

	l.Map[l.GuardPos.X][l.GuardPos.Y] = visitSymbol
}

func (l *LabMap) ShowMap() {
	totalVisited := 0

	for i := 0; i < l.Height; i++ {
		visited := 0
		for j := 0; j < l.Width; j++ {
			cell := l.Map[i][j]
			if AlreadyVisited(MapSymbol(cell)) {
				visited++
			}
			fmt.Printf("%s", string(cell))
		}
		totalVisited += visited
		fmt.Printf(" | %d\n", visited)
	}

	for i := 0; i < l.Width; i++ {
		fmt.Printf("-")
	}
	fmt.Printf(" | %d\n", totalVisited)
}

// ComputeGuardWalk walks across the lab and tracks the unique walked cells
// Returns the number of unique walked cells
func (l *LabMap) ComputeGuardWalk() int {
	guardStillInLab := true
	for guardStillInLab {
		guardStillInLab = l.MoveGuard()
		l.ShowMap()
		fmt.Println()
	}

	l.ShowMap()
	return l.UniqueWalkedCells
}

// MoveGuard moves a guard if possible
// Returns whether they leave map boundaries
func (l *LabMap) MoveGuard() bool {
	if l.NextMoveLeaveBounds() {
		l.UniqueWalkedCells++
		l.VisitCurrentPosition()
		return false
	}

	if l.GuardCanMove() {
		l.GuardStep()
	} else {
		l.TurnGuard()
	}

	currentPosition := l.Map[l.GuardPos.X][l.GuardPos.Y]
	switch currentPosition {
	case rune(Empty):
		l.UniqueWalkedCells++
	case rune(Blocked):
		log.Fatal("couldn't walk into a blocked cell!")
	case rune(HrVisited):
	case rune(VrVisited):
	case rune(BdVisited):
		// do nothing, keeping this case for checking purposes
	default:
		log.Fatal("invalid rune on position")
	}

	return true
}

func (l *LabMap) NextMoveLeaveBounds() bool {
	switch l.GuardOrientation {
	case Up:
		if l.GuardPos.X == 0 {
			return true
		}
	case Right:
		if l.GuardPos.Y == (l.Width - 1) {
			return true
		}
	case Down:
		if l.GuardPos.X == (l.Height - 1) {
			return true
		}
	case Left:
		if l.GuardPos.Y == 0 {
			return true
		}
	}

	return false
}

func (l *LabMap) GuardCanMove() bool {
	nextMoveCell := l.PeekNextMove()

	switch nextMoveCell {
	case string(HrVisited):
		return true
	case string(VrVisited):
		return true
	case string(BdVisited):
		return true
	case string(Empty):
		return true
	case string(Blocked):
		return false
	default:
		log.Fatal("cell cleanup not properly done")
	}

	return false
}

// PeekNextMove shows the contents of the guard's next move
// Returns the contents of next cell's move
// This method does not perform boundary checks; validate before
// e.g. with NextMoveLeaveBounds()
func (l *LabMap) PeekNextMove() string {
	switch l.GuardOrientation {
	case Up:
		return string(l.Map[l.GuardPos.X-1][l.GuardPos.Y])
	case Right:
		return string(l.Map[l.GuardPos.X][l.GuardPos.Y+1])
	case Down:
		return string(l.Map[l.GuardPos.X+1][l.GuardPos.Y])
	case Left:
		return string(l.Map[l.GuardPos.X][l.GuardPos.Y-1])
	}

	log.Fatal("invalid orientation")
	return ""
}

func (l *LabMap) GuardStep() {
	l.VisitCurrentPosition()

	switch l.GuardOrientation {
	case Up:
		l.GuardPos.X--
	case Right:
		l.GuardPos.Y++
	case Down:
		l.GuardPos.X++
	case Left:
		l.GuardPos.Y--
	default:
		log.Fatal("invalid orientation")
	}
}
