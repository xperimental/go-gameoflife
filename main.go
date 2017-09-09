package main

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/buger/goterm"
	"github.com/spf13/pflag"
)

func main() {
	inputFile := pflag.StringP("input", "i", "", "File to read start grid from.")
	useRandom := pflag.BoolP("random", "r", false, "Use random input.")
	randomRows := pflag.IntP("rows", "y", 50, "Number of rows for random input.")
	randomColumns := pflag.IntP("columns", "x", 200, "Number of columns for random input.")
	delay := pflag.DurationP("delay", "d", 16*time.Millisecond, "Delay between frames.")
	pflag.Parse()

	if len(*inputFile) == 0 && !*useRandom {
		pflag.Usage()
		return
	}

	grid, err := createGrid(*inputFile, *useRandom, *randomRows, *randomColumns)
	if err != nil {
		log.Fatalf("Error creating grid: %s", err)
	}

	goterm.Clear()
	for {
		goterm.MoveCursor(1, 1)

		start := time.Now()
		ascii := renderGrid(grid)
		goterm.Println(ascii)

		grid = calculateNextGeneration(grid)
		loopTime := time.Since(start)
		goterm.Printf("Loop time: %s\n", loopTime)
		goterm.Flush()

		wait := *delay - loopTime
		if wait > 0 {
			time.Sleep(wait)
		}
	}
}

type cellGrid [][]bool

type invalidCharacterError rune

func (e invalidCharacterError) Error() string {
	return fmt.Sprintf("invalid character: %v", rune(e))
}

func createGrid(inputFile string, random bool, rows, cols int) (cellGrid, error) {
	if random {
		return createRandomGrid(rows, cols)
	}

	grid, err := readGrid(inputFile)
	if err != nil {
		return cellGrid{}, fmt.Errorf("error reading grid: %s", err)
	}

	return grid, nil
}

func createRandomGrid(rows, cols int) (cellGrid, error) {
	if rows == 0 {
		return cellGrid{}, errors.New("need at least one row")
	}

	if cols == 0 {
		return cellGrid{}, errors.New("need at least one column")
	}

	rand.Seed(time.Now().Unix())

	grid := cellGrid{}
	for rowIndex := 0; rowIndex < rows; rowIndex++ {
		row := []bool{}
		for colIndex := 0; colIndex < cols; colIndex++ {
			if rand.Float64() > 0.3 {
				row = append(row, true)
			} else {
				row = append(row, false)
			}
		}
		grid = append(grid, row)
	}

	return grid, nil
}

func readGrid(fileName string) (cellGrid, error) {
	if len(fileName) == 0 {
		return cellGrid{}, errors.New("need an input file")
	}

	file, err := os.Open(fileName)
	if err != nil {
		return cellGrid{}, fmt.Errorf("error opening file %s: %s", fileName, err)
	}

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return cellGrid{}, fmt.Errorf("error reading input: %s", err)
	}

	if len(bytes) == 0 {
		return cellGrid{}, errors.New("input is empty")
	}

	grid, err := parseASCII(string(bytes))
	if err != nil {
		return cellGrid{}, fmt.Errorf("error parsing starting grid: %s", err)
	}

	return grid, nil
}

func parseASCII(input string) (cellGrid, error) {
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
			if rowOffset == 0 && colOffset == 0 {
				continue
			}

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
