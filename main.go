package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	bytes, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalf("Error reading input: %s", err)
	}

	if len(bytes) == 0 {
		log.Fatalf("Need some input.")
	}

	grid, err := parseAscii(string(bytes))
	if err != nil {
		log.Fatalf("Error parsing starting grid: %s", err)
	}

	for {
		start := time.Now()
		ascii := renderGrid(grid)
		fmt.Println(ascii)

		grid = calculateNextGeneration(grid)
		loopTime := time.Since(start)
		fmt.Printf("Loop time: %s", loopTime)
	}
}

type cellGrid [][]bool

type invalidCharacterError rune

func (e invalidCharacterError) Error() string {
	return fmt.Sprintf("invalid character: %v", rune(e))
}

func parseAscii(input string) (cellGrid, error) {
	result := cellGrid{}
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		gridLine := []bool{}
		for _, rune := range line {
			switch rune {
			case '*':
				gridLine = append(gridLine, true)
			case '.':
				gridLine = append(gridLine, false)
			default:
				return result, invalidCharacterError(rune)
			}
		}
		result = append(result, gridLine)
	}
	return result, nil
}

func renderGrid(grid cellGrid) string {
	out := &bytes.Buffer{}
	for _, row := range grid {
		for _, cell := range row {
			if cell {
				out.WriteRune('*')
			} else {
				out.WriteRune('.')
			}
		}
		out.WriteRune('\n')
	}
	return out.String()
}

func calculateNextGeneration(input cellGrid) cellGrid {
	next := cellGrid{}
	for rowIndex, row := range input {
		nextRow := []bool{}
		for colIndex, cell := range row {
			aliveNeighbors := calculateAliveNeighbors(input, rowIndex, colIndex)
			nextState := getNextState(cell, aliveNeighbors)
			nextRow = append(nextRow, nextState)
		}
		next = append(next, nextRow)
	}
	return next
}

func calculateAliveNeighbors(grid cellGrid, rowIndex, colIndex int) int {
	alive := 0
	for rowOffset := -1; rowOffset < 2; rowOffset++ {
		for colOffset := -1; colOffset < 2; colOffset++ {
			row := rowIndex + rowOffset
			col := colIndex + colOffset

			if row < 0 {
				continue
			}

			if row >= len(grid) {
				continue
			}

			if col < 0 {
				continue
			}

			if col >= len(grid[row]) {
				continue
			}

			if !grid[row][col] {
				continue
			}

			alive++
		}
	}
	return alive
}

func getNextState(state bool, aliveNeighbors int) bool {
	if state && aliveNeighbors == 2 {
		return true
	}

	if aliveNeighbors == 3 {
		return true
	}

	return false
}
