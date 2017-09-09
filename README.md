# go-gameoflife

This is a naive implementation of [Conway's Game of Life](https://en.wikipedia.org/wiki/Conway%27s_Game_of_Life) in Golang with output on the terminal (no graphics).

## Usage

```
Usage of ./go-gameoflife:
  -x, --columns int      Number of columns for random input. (default 200)
  -d, --delay duration   Delay between frames. (default 16ms)
  -i, --input string     File to read start grid from.
  -r, --random           Use random input.
  -y, --rows int         Number of rows for random input. (default 50)
```

The program can either load an existing "grid" from an input file using the `--input` option or randomly generate a grid using the `--random`, `--rows` and `--columns` options.

The `--delay` option sets a frame delay. The generation computation time is subtracted from this delay, so as long as the computation is faster than the delay frames should be output at the specified rate.

Examples:

```bash
# Run the glidergun example
go-gameoflife -i glidergun.txt
# Run a random (200x50) grid
go-gameoflife -r
```
