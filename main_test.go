package main

import (
	"reflect"
	"testing"
)

func TestParseAscii(t *testing.T) {
	ascii := `........
....*...
...**...
........`

	expectedGrid := cellGrid{
		{false, false, false, false, false, false, false, false},
		{false, false, false, false, true, false, false, false},
		{false, false, false, true, true, false, false, false},
		{false, false, false, false, false, false, false, false},
	}

	grid, err := parseAscii(ascii)

	if err != nil {
		t.Errorf("got error '%s', wanted nil", err)
	}

	if !reflect.DeepEqual(grid, expectedGrid) {
		t.Errorf("got '%v', wanted '%v'", grid, expectedGrid)
	}
}

func TestParseAsciiCharacterError(t *testing.T) {
	ascii := "invalid"
	expectedErr := invalidCharacterError('i')

	_, err := parseAscii(ascii)

	if err != expectedErr {
		t.Errorf("got error '%s', wanted '%s'", err, expectedErr)
	}
}

func TestRenderGrid(t *testing.T) {
	grid := cellGrid{
		{false, false, false, false, false, false, false, false},
		{false, false, false, false, true, false, false, false},
		{false, false, false, true, true, false, false, false},
		{false, false, false, false, false, false, false, false},
	}

	expectedAscii := `........
....*...
...**...
........
`

	ascii := renderGrid(grid)

	if ascii != expectedAscii {
		t.Errorf("got '%s', wanted '%s'", ascii, expectedAscii)
	}
}
