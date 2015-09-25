/**
* Author: Long He
* Andrew ID: longh
* Date: 10/12/2014
* Version: 2.3.1
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*===============================================================
 * Functions to manipulate a "field" of cells --- the main data
 * that must be managed by this program.
 *==============================================================*/

// The data stored in a single cell of a field
type Cell struct {
	kind  string
	score float64
	// for the extra credit:
	// use C or D to denote the previous kind
	preKind string
}

// createField should create a new field of the ysize rows and xsize columns,
// so that field[r][c] gives the Cell at position (r,c).
func createField(rsize, csize int) [][]Cell {
	f := make([][]Cell, rsize)
	for i := range f {
		f[i] = make([]Cell, csize)
	}
	return f
}

// inField returns true iff (row,col) is a valid cell in the field
func inField(field [][]Cell, row, col int) bool {
	return row >= 0 && row < len(field) && col >= 0 && col < len(field[0])
}

// readFieldFromFile should open the given file and read the initial
// values for the field. The first line of the file will contain
// two space-separated integers saying how many rows and columns
// the field should have:
//    10 15
// each subsequent line will consist of a string of Cs and Ds, which
// are the initial strategies for the cells:
//    CCCCCCDDDCCCCCC
//
// If there is ever an error reading, this function should cause the
// program to quit immediately.
func readFieldFromFile(filename string) [][]Cell {
	// open the file and make sure all went well
	file, err := os.Open("filename")
	if err != nil {
		fmt.Println("Error: something went wrong opening the file.")
		fmt.Println("Probably you gave the wrong filename.")
		os.Exit(3)
	}

	// creat the variable to hold the lines
	var lines []string = make([]string, 0)

	// read data from the file
	// creat a scanner using the bufio.newScanner function
	scanner := bufio.NewScanner(file)

	// loop through the lines of a file using a for loop
	for scanner.Scan() {

		// get the current line as a string using scanner.Text()
		// append it to the lines slice
		lines = append(lines, scanner.Text())
	}

	// check to see if there was an error duing the file reading
	if scanner.Err() != nil {
		fmt.Println("Error: there was a problem reading the file")
		// exit if there is an error
		os.Exit(3)
	}

	// parse the first line to get sizes of rows and columns
	var rsize, csize int
	fmt.Sscanf(lines[0], "%v %v", &rsize, &csize)


	// create a new field of the rsize rows and csize columns,
	var cell [][]Cell = createField(rsize, csize)

	// read following lines and store initial values
	// for each row of this field
	for i := 0; i < rsize; i++ {
		// for each column of this field
		// split each line into single characters
		var temp []string = strings.Split(lines[i + 1], "")
		for j := 0; j < csize; j++ {
			// store types in cell[i][j]
			cell[i][j].kind = temp[j]
		}
	}

	// close the file and retrun the lines
	file.Close()
	return cell
}

// drawField should draw a representation of the field on a canvas and save the
// canvas to a PNG file with a name given by the parameter filename.  Each cell
// in the field should be a 5-by-5 square, and cells of the "D" kind should be
// drawn red and cells of the "C" kind should be drawn blue.
func drawField(field [][]Cell, filename string) {
    /** create a new canvas
    * w width
    * h height
    * creat a canvas that we will be able to draw on
    */
    var w int = len(field[0]) * 5;
    var h int = len(field) * 5;
    canvas := CreateNewCanvas(w, h)

    // for each row of this field
    for i := 0; i < len(field); i++ {
    	// for each column of this field
    	for  j := 0; j <len(field[0]); j++ {
    		// for D-prisoner
    		if field[i][j].kind == "D" {
    			if field[i][j].preKind == "D" {
    				// set fill color as red for case: D following a D
    				canvas.SetFillColor(MakeColor(255, 0, 0))
    			} else {
    				// set fill color as yellow for case: D following a C
    				canvas.SetFillColor(MakeColor(255, 255, 0))
    			}	
    			// draw a rectangle in right position
    			canvas.ClearRect(i * 5, j * 5, i * 5 + 5, j * 5 + 5)
    		// for C-prisoner
    		} else {
    			if field[i][j].preKind == "C" {
    				// set fill color as blue for case: C following a C
    				canvas.SetFillColor(MakeColor(0, 0, 255))
    			} else {
    				// set fill color as green for case: C following a D
    				canvas.SetFillColor(MakeColor(0, 255, 0))
    			}	
    			// draw a rectangle in right position
    			canvas.ClearRect(i * 5, j * 5, i * 5 + 5, j * 5 + 5)
    		}
    	}
    }
    // save this canvas to PNG
    canvas.SaveToPNG(filename)
}

/*===============================================================
 * Functions to simulate the spatial games
 *==============================================================*/

// play a game between a cell of type "me" and a cell of type "them" (both me
// and them should be either "C" or "D"). This returns the reward that "me"
// gets when playing against them.
func gameBetween(me, them string, b float64) float64 {
	if me == "C" && them == "C" {
		return 1
	} else if me == "C" && them == "D" {
		return 0
	} else if me == "D" && them == "C" {
		return b
	} else if me == "D" && them == "D" {
		return 0
	} else {
		fmt.Println("type ==", me, them)
		panic("This shouldn't happen")
	}
}

// updateScores goes through every cell, and plays the Prisoner's dilema game
// with each of it's in-field nieghbors (including itself). It updates the
// score of each cell to be the sum of that cell's winnings from the game.
func updateScores(field [][]Cell, b float64) {
    // for each row of this field
    for i := 0; i < len(field); i++ {
    	// for each column of this field
    	for j := 0; j < len(field[0]); j++ {
    		// use calculator function to calculate the score of filed[i][j]
    		field[i][j].score = calculator(field, i, j, b)
    	}
    }
}

// calculator goes plays the Prisoner's dilema game with each of it's in-field 
// neighbors (including itself). It takes field and indexes i, j and returns 
// the sum of that its winnings from the game.
func calculator(field [][]Cell, i, j int, b float64) float64 {
	// initialize the score
	var score float64 = 0;
	// for each row of its neighbors
	for p := i - 1; p < i + 2; p++ {
		// for each column of its neighbors
		for q := j - 1; q < j + 2; q++ {
			// make sure field[p][q] is valid
			if inField(field, p, q) {
				// accumulate the score using gameBetween function
				score += gameBetween(field[i][j].kind, field[p][q].kind, b)
			}
		}
	}
	// return the sum of that its winnings from the game
	return score
}

// updateStrategies create a new field by going through every cell (r,c), and
// looking at each of the cells in its neighborhood (including itself) and the
// setting the kind of cell (r,c) in the new field to be the kind of the
// neighbor with the largest score
func updateStrategies(field [][]Cell) [][]Cell {
	// create a new field of the same size to store the updated kinds
	newField := createField(len(field), len(field[0]))	
	// for each row of field
	for i := 0; i < len(field); i++ {
    	// for each column of field
    	for j := 0; j < len(field[0]); j++ {
    		// use updater function to calculate update the kind of 
    		// field[i][j] and save it as the newField[i][j].kind
    		// save the previous kind
    		newField[i][j].preKind = field[i][j].kind
    		newField[i][j].kind = updater(field, i, j)
    	}
    }
    // return the updated cell	
	return newField
}

// updater takes field and indexes i, j and returns the updated kind.
// It makes sure field[d][k] is in-field, and compare its kind 
// with kind of field[i][j] to decide which cell should be returned.
func updater(field [][]Cell, i, j int) string {
	// use max to denote the highest score among its neighbors
	var max float64 = field[i][j].score
	// use updateKind to denote the kind should be stored in the newField
	var updateKind string = field[i][j].kind
	// for each row of its neighbors 
	for p := i - 1; p < i + 2; p++ {
		// for each column of its neighbors
    	for q := j - 1; q < j + 2; q++ {
    		// if field[p][q] is valid and its score is greater, update 
    		// the max score and store field[p][q].kind
    		if inField(field, p, q) && (field[p][q].score > max) {
    			max = field[p][q].score
    			updateKind = field[p][q].kind
    		}
    	}
    }
    // return the updated kind
    return updateKind
}

// evolve takes an intial field and evolves it for nsteps according to the game
// rule. At each step, it should call "updateScores()" and the updateStrategies
func evolve(field [][]Cell, nsteps int, b float64) [][]Cell {
	for i := 0; i < nsteps; i++ {
		updateScores(field, b)
		field = updateStrategies(field)
	}
	return field
}

// Implements a Spatial Games version of prisoner's dilemma. The command-line
// usage is:
//     ./spatial field_file b nsteps
// where 'field_file' is the file continaing the initial arrangment of cells, b
// is the reward for defecting against a cooperator, and nsteps is the number
// of rounds to update stategies.
//
func main() {
	// parse the command line
	if len(os.Args) != 4 {
		fmt.Println("Error: should spatial field_file b nsteps")
		return
	}

	fieldFile := os.Args[1]

	b, err := strconv.ParseFloat(os.Args[2], 64)
	if err != nil || b <= 0 {
		fmt.Println("Error: bad b parameter.")
		return
	}

	nsteps, err := strconv.Atoi(os.Args[3])
	if err != nil || nsteps < 0 {
		fmt.Println("Error: bad number of steps.")
		return
	}

    // read the field
	field := readFieldFromFile(fieldFile)
    fmt.Println("Field dimensions are:", len(field), "by", len(field[0]))

    // evolve the field for nsteps and write it as a PNG
	field = evolve(field, nsteps, b)

	drawField(field, "Prisoners.png")
}
