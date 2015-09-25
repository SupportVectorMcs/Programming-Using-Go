// name: Long He
// andrew id: longh
// version 0.1
// date: 29/10/2014 

package main

import (
	"bufio"
	"code.google.com/p/draw2d/draw2d"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
	"strconv"
)

type nCanvas struct {
	gc     *draw2d.ImageGraphicContext
	img    image.Image
	width  int
	height int
}

// nMoveTo moves the current point to position (x, y)
func (c *nCanvas) nMoveTo(x, y float64) {
	c.gc.MoveTo(x, y)
}

// nLineTo draws a line from the current point to (x,y)
// and makes (x, y) the current point using the current 
// line color and current line width
func (c *nCanvas) nLineTo(x, y float64) {
	c.gc.LineTo(x, y)
}

// nSetStrokeColor sets the current line color to col
func (c *nCanvas) nSetStrokeColor(col color.Color) {
	c.gc.SetStrokeColor(col)
}


// nSetFillColor sets the current fill color to col
func (c *nCanvas) nSetFillColor(col color.Color) {
	c.gc.SetFillColor(col)
}

// nSetLineWidth sets the current line width to w
func (c *nCanvas) nSetLineWidth(w float64) {
	c.gc.SetLineWidth(w)
}

// nStroke draws the lines specified by LineTo calls, and 
// clear the pending lines
func (c *nCanvas) nStroke() {
	c.gc.Stroke()
}

// nFillStroke draws the lines specified by LineTo calls 
// but in addition fills the region inside the lines with 
// the current fill color
func (c *nCanvas) nFillStroke() {
	c.gc.FillStroke()
}

// nFill draws the lines specide
func (c *nCanvas) nFill() {
	c.gc.Fill()
}

// nClear fills the entire canvas with the current fill color
func (c *nCanvas) nClear() {
	c.gc.Clear()
}

// nClearRect takes the location (x1, y1) and (x2, y3) and creats 
// a rectangle on the Canvas 
func (c *nCanvas) nClearRect(x1, y1, x2, y2 int) {
	c.gc.ClearRect(x1, y1, x2, y2)
}

// nSaveToPNG takes the filename and save the canvas to PNG
func (c *nCanvas) SaveToPNG(filename string) {
	f, err := os.Create(filename)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	defer f.Close()
	b := bufio.NewWriter(f)
	err = png.Encode(b, c.img)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	err = b.Flush()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Wrote %s OK.\n", filename)
}

// Width returns the Width of the Canvas
func (c *nCanvas) Width() int {
	return c.width
}

// Height returns the height of the Canvas
func (c *nCanvas) Height() int {
	return c.height
}

// nCreateNewCanvas takes the width and height and 
// creates a new white canvas that we will be able to draw on 
func nCreateNewCanvas(w, h int) nCanvas {
	i := image.NewRGBA(image.Rect(0, 0, w, h))
	gc := draw2d.NewGraphicContext(i)

	gc.SetStrokeColor(image.Black)
	gc.SetFillColor(image.White)
	// fill the background
	gc.Clear()
	gc.SetFillColor(image.Black)

	return nCanvas{gc, i, w, h}
}

// nMakeColor takes the parameters for R, G and B and return the color
func nMakeColor(r, g, b uint8) color.Color {
	return &color.RGBA{r, g, b, 255}
}

// create a new type called Board, with the following methods
type Board struct {
	// board save the number of coins in each square
	board [][]int
	size int
}

// NewBoard returns a new Board with values of size*size squares
func NewBoard(size int) *Board {
	var b [][]int
	// initialize the board
	b = make([][]int, size)
	for i := 0; i < size; i++ {
		b[i] = make([]int, size)
	}
	return &Board{b, size}
}

// Topple topples (r, c) until it can't be toppled any more
func (S *Board) Topple(r, c int) {
	for S.board[r][c] >= 4 {
		// move 1 coin from (r, c) to each of the 4 neighbors of 
		// (r, c) (diagonal neighbors don’t count, only north, south, 
		// east, and west)
		if S.Contains (r-1, c) {
			S.board[r-1][c] += 1
			S.board[r][c] -= 1 
		// for the boundary cases, throw the coin out of the checkboard
		} else if !S.Contains (r-1, c) {
			S.board[r][c]--
		}
		if S.Contains (r+1, c) {
			S.board[r+1][c]++
			S.board[r][c]--
		// for the boundary cases, throw the coin out of the checkboard
		} else if !S.Contains (r+1, c) {
			S.board[r][c]--
		}
		if S.Contains (r, c-1) {
			S.board[r][c-1]++
			S.board[r][c]--
		// for the boundary cases, throw the coin out of the checkboard
		} else if !S.Contains (r, c-1) {
			S.board[r][c]--
		}
		if S.Contains (r, c+1) {
			S.board[r][c+1]++
			S.board[r][c]--
		// for the boundary cases, throw the coin out of the checkboard
		} else if !S.Contains (r, c+1) {
			S.board[r][c]--
		}
	}
}

// Contains returns true if (r, c) is within the field
func (S *Board) Contains(r, c int) bool {
	// make sure r and c both are in the range
	if r >= 0 && r < S.size && c >= 0 && c < S.size {
		return true
	}
	return false
}

// Set sets the value of cell (r, c)
func (S *Board) Set(r, c, value int) {
	S.board[r][c] = value
}

// Cell returns the value of the cell (r, c)
func (S *Board) Cell(r, c int) int {
	return S.board[r][c]
}

// IsConverged returns ture if there are no cells with >= 4 coins on them
func (S *Board) IsConverged() bool {
	// for each row
	for i := 0; i < S.size; i++ {
		// for each column
		for j := 0; j < S.size; j++ {
			// return false if one of the square has no less than 4 coins
			if S.board[i][j] >= 4 {
				return false
			}
		}
	}
	return true
}

// NumRows returns the number of rows on the board
func (S *Board) NumRows() int {
	return S.size/2
}

// Numcols returns the number of columns on the board
func (S *Board) NumCols() int {
	return S.size/2
}

// ComputeSteadyStates topples squares until the board has converged 
// to a stable configuration
func (S *Board) ComputeSteadyStates() {
	// compute until it is converged
	for !S.IsConverged() {
		// for each row
		for i := 0; i < S.size; i++ {
			// for each column
			for j := 0; j < S.size; j++ {
				S.Topple(i, j)
			}
		}
	}
}

// DrawBoard draws the board to a PNG
func (S *Board) DrawBoard() {
	// set the name of the output png
	var filename string = "board.png"

	// create a new canvas
	c := nCreateNewCanvas(S.size, S.size)
	// for each row
	for i := 0; i < S.size; i++ {
		// for each column
		for j := 0; j < S.size; j++ {
			// set the colors for squares with different numbers of coins
			if S.board[i][j] == 1 {
				c.nSetFillColor(nMakeColor(85, 85, 85))
			} else if S.board[i][j] == 2 {
				c.nSetFillColor(nMakeColor(170, 170, 170))
			} else if S.board[i][j] == 3 {
				c.nSetFillColor(nMakeColor(255, 255, 255))
			} else if S.board[i][j] == 0 {
				c.nSetFillColor(nMakeColor(0, 0, 0))
			}
			// draw the squares
			c.nClearRect(i, j, i+1, j+1)
		}
	}
	// save the canvas to filename.PNG
	c.SaveToPNG(filename)
}

func main() {
	// make sure the size of checkboard is valid
    size, err := strconv.Atoi(os.Args[1])
    if err != nil {
        fmt.Println("Error: size must be an integer.")
        return
    } else if size < 0 {
        fmt.Println("Error: size must be positive.")
        return
    }

    // make sure the number of coins that are to be placed is valid
    pile, err := strconv.Atoi(os.Args[2])
    if err != nil {
        fmt.Println("Error: pile must be an integer.")
        return
    } else if pile < 0 {
        fmt.Println("Error: pile must be positive.")
        return
    }

    // initialize a new Board
    S := NewBoard(size)

    // set the initial number of coins in the middle of this checkboard
    S.Set(size/2, size/2, pile)

    // ComputeSteadyStates topples squares until the board has converged 
    // to a stable configuration
    S.ComputeSteadyStates()

    // draw the board to a PNG
    S.DrawBoard()
}