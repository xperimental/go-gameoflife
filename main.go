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
	return input
}
