/*
 * Author: Long He
 * Andrew ID: longh
 * Date: 12/03/2014
 */

package main

import (
    "fmt"
    "os"
    "math/rand"
    "time"
    "math"
)

func main() {
    // make sure the string representing the HP protein sequence is entered
    if len(os.Args) <= 1 {
        fmt.Println("Error: Need the HP protein sequence.")
        return
    }

    // parse the string and save it as a []string
    var seq []string
    for _, i := range os.Args[1] {
        // make sure the representing sequence is right
        if string(i) != "P" && string(i) != "H" {
            fmt.Println("Error: The sequence can only contain H and P.")
            return
        }
        // add H or P into this slice
        seq = append(seq, string(i))
    }

    // optimizeFold takes a protein sequence and returns the lowest-energy 
    // fold and lowest-energy by running the simulated annealing algorithm
    fold, en := optimizeFold(seq)

    // structure to return
    var structure string
    for _, i := range fold {
        structure += i
    }
    fmt.Println("Structure: ", structure)
    fmt.Println("Energy: ", en)

    drawFold(fold, seq)
}

// randomFold takes the length of the protein sequence and returns a random
// fold as a sequence of l, r, and f commands
func randomFold(len int) []string {
    // use the time as the seed for the random number generator
    rand.Seed(time.Now().UnixNano())

    var fold []string
    for len > 0 {
        // l, r and f all have 1 / 3 chance
        temp := rand.Float64() * 3
        if temp < 1 {
            fold = append(fold, "l")
        } else if temp < 2 {
            fold = append(fold, "r")
        } else {
            fold = append(fold, "f")
        }
        len--
    }
    return fold;
}

// matFold takes a fold S and a sequence p, generates the fold onto a 
// sufficiently large 2D matrix and return this matrix
func matFold(S []string, p []string) (bool, [][]string) {
    var mat [][]string = make([][]string, 100)
    for i := 0; i < 100; i++ {
        mat[i] = make([]string, 100)
    }
    // location to start
    x := 50
    y := 50
    mat[x][y] = p[0]
    // initial direction
    direction := "up"
    for i := 0; i < len(S); i++ {
        if direction == "up" {
            if S[i] == "l" {
                x--
                mat[x][y] = "line"
                x--
                // refuse cross
                if mat[x][y] != "" {
                    return true, mat
                }
                // save the point
                mat[x][y] = p[i + 1]
                // change the direction
                direction = "left"
            } else if S[i] == "r" {
                x++
                mat[x][y] = "line"
                x++
                // refuse cross
                if mat[x][y] != "" {
                    return true, mat
                }
                // save the point
                mat[x][y] = p[i + 1]
                // change the direction
                direction = "right"
            } else if S[i] == "f" {
                y--
                mat[x][y] = "line"
                y--
                // refuse cross
                if mat[x][y] != "" {
                    return true, mat
                }
                // save the point
                mat[x][y] = p[i + 1]
            }
        } else if direction == "down" {
            if S[i] == "l" {
                x++
                mat[x][y] = "line"
                x++
                // refuse cross
                if mat[x][y] != "" {
                    return true, mat
                }
                // save the point
                mat[x][y] = p[i + 1]
                // change the direction
                direction = "right"
            } else if S[i] == "r" {
                x--
                mat[x][y] = "line"
                x--
                // refuse cross
                if mat[x][y] != "" {
                    return true, mat
                }
                // save the point
                mat[x][y] = p[i + 1]
                // change the direction
                direction = "left"
            } else if S[i] == "f" {
                y++
                mat[x][y] = "line"
                y++
                // refuse cross
                if mat[x][y] != "" {
                    return true, mat
                }
                // save the point
                mat[x][y] = p[i + 1]
            } 
        } else if direction == "left" {
            if S[i] == "l" {
                y++
                mat[x][y] = "line"
                y++
                // refuse cross
                if mat[x][y] != "" {
                    return true, mat
                }
                // save the point
                mat[x][y] = p[i + 1]
                // change the direction
                direction = "down"
            } else if S[i] == "r" {
                y--
                mat[x][y] = "line"
                y--
                // refuse cross
                if mat[x][y] != "" {
                    return true, mat
                }
                // save the point
                mat[x][y] = p[i + 1]
                // change the direction
                direction = "up"
            } else if S[i] == "f" {
                x--
                mat[x][y] = "line"
                x--
                // refuse cross
                if mat[x][y] != "" {
                    return true, mat
                }
                // save the point
                mat[x][y] = p[i + 1]
            } 
        } else if direction == "right" {
            if S[i] == "l" {
                y--
                mat[x][y] = "line"
                y--
                // refuse cross
                if mat[x][y] != "" {
                    return true, mat
                }
                // save the point
                mat[x][y] = p[i + 1]
                // change the direction
                direction = "up"
            } else if S[i] == "r" {
                y++
                mat[x][y] = "line"
                y++
                // refuse cross
                if mat[x][y] != "" {
                    return true, mat
                }
                // save the point
                mat[x][y] = p[i + 1]
                // change the direction
                direction = "down"
            } else if S[i] == "f" {
                x++
                mat[x][y] = "line"
                x++
                // refuse cross
                if mat[x][y] != "" {
                    return true, mat
                }
                // save the point
                mat[x][y] = p[i + 1]
            } 
        }
    }
    return false, mat 
}

// drawFold takes a 2D matrix and "draws" this fold
func drawFold(fold []string, p []string) {
    /** create a new canvas
    * w width
    * h height
    * creat a canvas that we will be able to draw on
    */
    w := 1000
    h := 1000
    canvas := CreateNewCanvas(w, h)

    // set stroke color as black and line width as 1
    canvas.SetLineWidth(1)
    canvas.SetStrokeColor(MakeColor(0, 0, 0))
    
    // start from the center of this canvas
    x := float64(w / 2)
    y := float64(h / 2)
    canvas.gc.ArcTo(x, y, 5, 5, 0, 2 * math.Pi)
    if p[0] == "H" {
        // draw a node
        canvas.Fill()
    } else {
        // draw a circle
        canvas.Stroke()
    }
    // move to prepare for drawing a line
    canvas.MoveTo(x,y)
    // the original direction
    t := math.Pi / 2
    d := math.Pi / 2
    // draw the whole fold
    for i := 0; i < len(fold); i++{
        if fold[i] == "l" {
            // turn left
            t += d
        }else if fold[i] == "r" {
            // turn right
            t -= d
        }
        x += 50 * math.Cos(t)
        y += -50 * math.Sin(t)
        // draw a line
        canvas.LineTo(x, y)
        canvas.Stroke()
        // draw a point
        if p[i + 1] == "H" {
            // draw a solid node
            canvas.gc.ArcTo(x, y, 5, 5, 0, 2 * math.Pi)
            canvas.Fill()
            canvas.MoveTo(x,y)
        } else {
            // draw a circle
            canvas.gc.ArcTo(x, y, 5, 5, 0, 2 * math.Pi)
            canvas.Stroke()
            canvas.MoveTo(x,y)
        }
    }
    canvas.SaveToPNG("fold.png")
}

// energy takes a fold S and a sequence p and returns the energy accodring to
// the energy function = 10 * x - sum(pi * si)
func energy(S []string, p []string) (bool, int) {
    cross, mat := matFold(S, p)
    if cross {
        return true, 0
    }
    // compute sum(pi * si)
    // compute the score ignoring which residues are adjacent in the walk, 
    // and then subtract 2 for every H in the middle of the sequence, and 1 
    // for every H at the ends of the sequence
    var sum int = 0
    for i := 2; i < len(mat) - 2; i++ {
        for j := 2; j < len(mat) - 2; j++ {
            if mat[i][j] == "H" {
                // adjacent to the ith point in the walk
                if (mat[i - 2][j] == "H" || mat[i - 2][j] == "P") && 
                mat[i - 1][j] != "line" {
                    sum++
                }
                if (mat[i + 2][j] == "H" || mat[i + 2][j] == "P") && 
                mat[i + 1][j] != "line" {
                    sum++
                }
                if (mat[i][j - 2] == "H" || mat[i][j - 2] == "P") && 
                mat[i][j - 1] != "line" {
                    sum++
                }
                if (mat[i][j + 2] == "H" || mat[i][j + 2] == "P") && 
                mat[i][j + 1] != "line" {
                    sum++
                }
                // the number of diagonals amino acids at neighboring lattice 
                // points
                if mat[i - 2][j - 2] == "H" || mat[i - 2][j - 2] == "P" {
                    sum++
                }
                if mat[i + 2][j + 2] == "H" || mat[i + 2][j + 2] == "P" {
                    sum++
                }
                if mat[i - 2][j + 2] == "H" || mat[i - 2][j + 2] == "P" {
                    sum++
                }
                if mat[i + 2][j - 2] == "H" || mat[i + 2][j - 2] == "P" {
                    sum++
                }
            }
        }
    }
    return false, - sum
}

// randomFoldChange takes a fold and randomly changes one of its commands
func randomFoldChange(fold []string) []string {
    // use the time as the seed for the random number generator
    rand.Seed(time.Now().UnixNano())
    temp := rand.Float64()

    // change the index'th commands
    index := rand.Intn(len(fold))

    // each of other two choices has 50% chance 
    if fold[index] == "l" {
        if temp > 0.5 {
            fold[index] = "r"
        } else {
            fold[index] = "f"
        }
    } else if fold[index] == "r" {
        if temp > 0.5 {
            fold[index] = "l"
        } else {
            fold[index] = "f"
        }
    } else {
        if temp > 0.5 {
            fold[index] = "l"
        } else {
            fold[index] = "r"
        }
    }

    return fold
}

// optimizeFold takes a protein sequence and returns the lowest-energy fold
// and lowest-energy by running the simulated annealing algorithm
func optimizeFold(p []string) ([]string, int) {
    // use the time as the seed for the random number generator
    rand.Seed(time.Now().UnixNano())

    MAX := 100000 // max iteration times
    k := float64(5 * len(p)) // parameter k
    T := float64(10 * len(p)) // parameter T
    m := 100000 // if the structure hasn't changed for m iterations, stop
    // count denotes the number of iterations the structure hasn't changed
    count := 0

    // set i = 1 and choose a random sequence of l, r, f
    // make sure there is no cross
    var S []string
    var e int
    var cross bool
    for true {
        S = randomFold(len(p) - 1)
        // compute ei = energy(Si, p)
        cross, e = energy(S, p)
        if cross == false {
            break
        }
    }

    var S_ []string
    var e_ int
    var cross_ bool
    for i := 1; i < MAX && count < m; i++ {
        // change a random letter of the walk Si to a random different letter 
        // to obtain a slightly different structure Si'
        // make sure there is no cross
        for true {
            S_ = randomFoldChange(S)
            // compute e' = energy(Si', p)
            cross_, e_ = energy(S_, p)
            if cross_ == false {
                break
            }
        }

        if e_ == e {
            count++
            continue
        } else if e_ < e { // if e' < e, set Si+1 = Si'
            S = S_
            e = e_
            count = 0
        // if e' > e let q = exp(-(e' - e) / k * T) and with probability q 
        // set Si+1 = S' and with probability (1 - q) set Si+1 = Si
        } else {
            q := math.Exp(-float64(e_ - e) / ( k * T))
            if rand.Float64() < q {
                S = S_
                e = e_
                count = 0
            } else {
                count++
            }
        }
        // let i = i + 1 and if i % 100 = 0 let T = 0.999 * T
        if i % 100 == 0 {
            T = 0.999 * T
        }
    }  
    return S, e
}