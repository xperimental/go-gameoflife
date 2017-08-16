package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
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

func parseAscii(input string) (cellGrid, error) {
	return cellGrid{}, nil
}

func renderGrid(grid cellGrid) string {
	return ""
}

func calculateNextGeneration(input cellGrid) cellGrid {
	return input
}
