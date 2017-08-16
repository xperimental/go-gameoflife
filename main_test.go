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

func TestGetNextState(t *testing.T) {
	for _, test := range []struct {
		desc  string
		state bool
		alive int
		next  bool
	}{
		{
			desc:  "fewer than two -> dead",
			state: true,
			alive: 1,
			next:  false,
		},
		{
			desc:  "two -> alive",
			state: true,
			alive: 2,
			next:  true,
		},
		{
			desc:  "three -> alive",
			state: true,
			alive: 3,
			next:  true,
		},
		{
			desc:  "more than three -> dead",
			state: true,
			alive: 4,
			next:  false,
		},
		{
			desc:  "dead but three -> alive",
			state: false,
			alive: 3,
			next:  true,
		},
	} {
		test := test
		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()

			next := getNextState(test.state, test.alive)

			if next != test.next {
				t.Errorf("got next '%v', wanted '%v'", next, test.next)
			}
		})
	}
}

func TestCalculateNextGeneration(t *testing.T) {
	for _, test := range []struct {
		ascii         string
		expectedAscii string
	}{
		{
			ascii: `........
....*...
...**...
........`,

			expectedAscii: `........
...**...
...**...
........
`,
		},
		{
			ascii: `........
...**...
...**...
........`,

			expectedAscii: `........
...**...
...**...
........
`,
		},
	} {
		grid, _ := parseAscii(test.ascii)

		next := calculateNextGeneration(grid)

		nextAscii := renderGrid(next)

		if nextAscii != test.expectedAscii {
			t.Errorf("got '%s', wanted '%s'", nextAscii, test.expectedAscii)
		}
	}
}
