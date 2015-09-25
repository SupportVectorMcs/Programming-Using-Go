/** Name: Long He
* Andrew ID: longh
* date: 09/27/2014
* version: 1.0
*/

package main

import (
    "fmt"
    "os"
    "strconv"
    "math/rand"
    "math"
)

func main() {
    
    // make sure the number and type of arguments is correct
    if len(os.Args) != 4 {
        fmt.Println("Error: r, d, RULE are needed.")
        return
    }


    // make sure r is valid
    r, err := strconv.ParseFloat(os.Args[1], 10)
    if err != nil {
        fmt.Println("Error: r must be a real number.")
        return
    } else if r < 0 {
        fmt.Println("Error: r must be be positive.")
        return
    }

    // make sure d is valid
    d, err := strconv.ParseFloat(os.Args[2], 10)
    if err != nil {
        fmt.Println("Error: d must be a real number.")
        return
    } else if d < 0 {
        fmt.Println("Error: d must be positive.")
        return
    }

    // make sure RULE is valid
    RULE, err := strconv.Atoi(os.Args[3])
    if err != nil {
        fmt.Println("Error: RULE must be an integer.")
        return
    } else if RULE < 0 {
        fmt.Println("Error: RULE must be positive.")
        return
    }

    // convert RULE to numbers between 0 and 255
    if 8 == len(os.Args[3]) {
        var temp int
        var magnit int = 1
        for RULE > 0 {
            temp += magnit * (RULE & 1)
            RULE /= 10
            magnit *= 2
        }
        RULE = temp
    } else if len(os.Args[3]) <= 3 { 
    } else {
        fmt.Println("Error: RULE is invalid.")
        return
    }

    // input for function PopSize
    var x_0 float64 = 0.1
    var max_t int = 100

    PopSize(r, x_0, max_t)

    // input for function RandomWalk
    var WIDTH float64 = 500
    var HEIGHT float64 = 500
    var STEPSIZE float64 = d
    var NUMBER_OF_STEPS int = 1000
    var SEED int64 = 12345

    randomWalk(WIDTH, HEIGHT, STEPSIZE, NUMBER_OF_STEPS, SEED)
    
    // input for function CA
    var WID int = 500
    var STEPS int = 50

    CA(RULE, WID, STEPS)

    // EXTRA CREDIT
    MyCoolPicture()
}

//==================================================
//                                            PART I
//==================================================

/** Function PopSize takes in r, x(0), Max(t) and 
* connects the points (5 * t, 100 - 100 * x_t)
*/
func PopSize(r, x_0 float64, max_t int) {
    /** create a new canvas
    * w width
    * h height
    * creat a canvas that we will be able to draw on
    * set line width as 1
    * set stroke color as blue
    */
    var w int = 500;
    var h int = 100;
    canvas := CreateNewCanvas(w, h)
    canvas.SetLineWidth(1)
    canvas.SetStrokeColor(MakeColor(0, 0, 255))

    var x_t float64 = x_0;
    for t := 0; t < max_t; {
        canvas.MoveTo(5 * float64(t), 100 - 100 * x_t)
        
        // calculate x(t + 1)
        t++
        x_t = r * x_t * (1.0 - x_t)
        if x_t > 1 {
            x_t = 1
        }
        if x_t < 0 {
            x_t = 0
        }

        canvas.LineTo(5 * float64(t), 100 - 100 * x_t)
        canvas.Stroke()
    }
    canvas.SaveToPNG("PopSize.png")
}

//==================================================
//                                           PART II
//==================================================

/** Function inFiled makes sure position within 
* [0, n) x [0, n)
*/
func inField(coord, n float64) bool {
    return coord >= 0 && coord < n
}

/** Function randStep takes in current position (x, y),
* STEPSIZE, WIDTH and HEIGHT, then returns random movements
*/
func randStep(x, y, STEPSIZE, WIDTH, HEIGHT float64) (float64, float64) {

    var nx, ny float64 = x, y
    for (nx == x && ny == y) || !inField(nx, WIDTH) || !inField(ny, HEIGHT) {
        
        // get the step directions
        var temp = rand.Float64() * 2 * math.Pi
        nx = x + STEPSIZE * math.Cos(temp)
        ny = y + STEPSIZE * math.Sin(temp)
    }
    return nx, ny
}

/** Function randwalk takes in WIDTH, HEIGHT, STEPSIZE, number of steps and SEED,
* then connects the points that random walk from assignment 2 would have visited
*/
func randomWalk(WIDTH, HEIGHT, STEPSIZE float64, NUMBER_OF_STEPS int, SEED int64) {
    
    // the seed for the random number generator
    rand.Seed(SEED)

    /** create a new canvas
    * w width
    * h height
    * creat a canvas that we will be able to draw on
    * set line width as 1
    * set stroke color as black
    */
    var w int = int(WIDTH);
    var h int = int(HEIGHT);
    canvas := CreateNewCanvas(w, h)
    canvas.SetLineWidth(1)
    canvas.SetStrokeColor(MakeColor(0, 0, 0))

    var x, y = WIDTH / 2, HEIGHT / 2

    for i := 0; i < NUMBER_OF_STEPS; {
        canvas.MoveTo(x, y)

        // calculate x(i + 1)
        i++
        x, y = randStep(x, y, STEPSIZE, WIDTH, HEIGHT)

        canvas.LineTo(x, y)
        canvas.Stroke()
    }
    canvas.SaveToPNG("RandomWalk.png")
}

//==================================================
//                                          PART III
//==================================================

/** Function CA takes in RULE, WIDTH and steps, then 
* draws the output of cellular automata where each 
* 5 pixel band is one line of previous output lines 
* and where # is drawn as a 5-by-5 black square and 
* a space as 5-by-5 yellow square.
*/
func CA(RULE, WIDTH, STEPS int) {

    /** create a new canvas
    * w width
    * h height
    * creat a canvas that we will be able to draw on
    */
    var w int = int(WIDTH);
    var h int = 255;
    canvas := CreateNewCanvas(w, h)

    // draw first row
    var i int
    for i = 0; i < WIDTH / 2; {                       // 1 to (WIDTH / 2 - 1)
        canvas.SetFillColor(MakeColor(255, 255, 0))
        canvas.ClearRect(i, 0, i + 5, 5)
        i += 5
    }
    canvas.SetFillColor(MakeColor(0, 0, 0))           // WIDTH / 2
    canvas.ClearRect(i, 0, i + 5, 5)                      
    for i = WIDTH / 2 + 5; i < WIDTH - 1; {           // (WIDTH / 2 + 1) to (WIDTH - 1)
        canvas.SetFillColor(MakeColor(255, 255, 0))
        canvas.ClearRect(i, 0, i + 5, 5)
        i += 5
    }
    canvas.SetFillColor(MakeColor(0, 0, 0))
    canvas.ClearRect(i + 5, 0, i + 10, 10)            // WIDTH

    // first row
    row1 := make([]int, WIDTH / 5 + 2)
    row1[WIDTH / (5 * 2) + 1] = 1;

    // draw following rows
    for j := 2; j < STEPS + 2; j++ {
        row1 = printNextRow(row1, RULE, canvas, j - 1) 
    }

    canvas.SaveToPNG("CA.png")
}

/** Function getRule takes in RULE and n
* then returns the next step instruction from rule
*/
func getRule(RULE, n int) (int){
    var ret int
    ret = RULE & int(math.Pow(2, float64(8 - n)))
    if ret > 0 {
        ret = 1
    }
    return ret
}

/** Function printNextRow takes in row, RULE, canvas
* and rowIndex, then draws the next row according to 
* its previous row
*/
func printNextRow(row []int, RULE int, canvas Canvas, rowIndex int) ([]int){
    nextRow := make([]int, 1)
    // analyze last row
    for k := 0; k < len(row) - 2; k++ {
        // get the rule index
        var ruleNo int = (1 - row[k]) * 4 + (1 - row[k + 1]) * 2 + (1 - row[k + 2]) + 1
        // get the instruction from rule
        var digit = getRule(RULE, ruleNo)
        nextRow = append(nextRow, digit)

        if 1 == digit {
            canvas.SetFillColor(MakeColor(0, 0, 0))
            canvas.ClearRect(k * 5, rowIndex * 5, k * 5 + 5, (rowIndex + 1) * 5)
        } else {
            canvas.SetFillColor(MakeColor(255, 255, 0))
            canvas.ClearRect(k * 5, rowIndex * 5, k * 5 + 5, (rowIndex + 1) * 5)
        }
    }
    nextRow = append(nextRow, 0)

    // save present states for next step
    row = nextRow
    return row
}

//==================================================
//                                      EXTRA CREDIT
//==================================================

/** Function MyCoolPicture draws a cool picture
* using various of colors and lines combinations 
*/
func MyCoolPicture() {

    /** create a new canvas
    * w width
    * h height
    * creat a canvas that we will be able to draw on
    * set the background as black
    * set the line width as 0.05
    */
    var w int = 500
    var h int = 500;
    canvas := CreateNewCanvas(w, h)
    canvas.SetFillColor(MakeColor(0, 0, 0))
    canvas.Clear()
    canvas.SetLineWidth(0.05)

    // draw a curve using lines and colors
    var j uint8 = 0;
    for i := 0; i < 500; {
        canvas.SetStrokeColor(MakeColor(255, 255 - j, j))
        canvas.MoveTo(0, float64(i))
        canvas.LineTo(float64(i), 500)
        canvas.Stroke()
        i++
        j++
    }

    // draw a curve using lines and colors
    j = 0
    for i := 0; i < 500; {
        canvas.SetStrokeColor(MakeColor(255, j, 255 - j))
        canvas.MoveTo(0, float64(i))
        canvas.LineTo(500 - float64(i), 0)
        canvas.Stroke()
        i++
        j++
    }

    // draw a curve using lines and colors
    j = 0
    for i := 0; i < 500; {
        canvas.SetStrokeColor(MakeColor(j, 255, 255 - j))
        canvas.MoveTo(500, float64(i))
        canvas.LineTo(500 - float64(i), 500)
        canvas.Stroke()
        i++
        j++
    }

    // draw a curve using lines and colors
    j = 0
    for i := 0; i < 500; {
        canvas.SetStrokeColor(MakeColor(255 - j, j, 255))
        canvas.MoveTo(float64(i), 0)
        canvas.LineTo(500, float64(i))
        canvas.Stroke()
        i++
        j++
    }

    // draw lines in different colors
    j = 0
    for i := 0; i < 500; {
        canvas.SetStrokeColor(MakeColor(255 - j, j, 255))
        canvas.MoveTo(float64(i), 0)
        canvas.LineTo(float64(i), 500)
        canvas.Stroke()
        i++
        j++
    }

    // draw lines in different colors
    j = 0
    for i := 0; i < 500; {
        canvas.SetStrokeColor(MakeColor(255, j, 255 - j))
        canvas.MoveTo(0, float64(i))
        canvas.LineTo(500, float64(i))
        canvas.Stroke()
        i++
        j++
    }

    // draw a triangle in different colors
    j = 0
    for i := 0; i < 500; {
        canvas.SetStrokeColor(MakeColor(255 - j, j, 255))
        canvas.MoveTo(250, 0)
        canvas.LineTo(float64(i), 500)
        canvas.Stroke()
        i++
        j++
    }

    // draw a triangle in different colors
    j = 0
    for i := 0; i < 500; {
        canvas.SetStrokeColor(MakeColor(j, 255, 255 - j))
        canvas.MoveTo(0, 250)
        canvas.LineTo(500, float64(i))
        canvas.Stroke()
        i++
        j++
    }

    // draw a triangle in different colors
    j = 0
    for i := 0; i < 500; {
        canvas.SetStrokeColor(MakeColor(255, j, 255 - j))
        canvas.MoveTo(250, 500)
        canvas.LineTo(float64(i), 0)
        canvas.Stroke()
        i++
        j++
    }

    // draw a triangle in different colors
    j = 0
    for i := 0; i < 500; {
        canvas.SetStrokeColor(MakeColor(255, 255 - j, j))
        canvas.MoveTo(500, 250)
        canvas.LineTo(0, float64(i))
        canvas.Stroke()
        i++
        j++
    }

    // draw lines in different colors
    j = 0
    for i := 0; i < 500; {
        canvas.SetStrokeColor(MakeColor(j, 255 - j, 255))
        canvas.MoveTo(500, 500 - float64(i))
        canvas.LineTo(0, float64(i))
        canvas.Stroke()
        i++
        j++
    }

    // draw lines in different colors
    j = 0
    for i := 0; i < 500; {
        canvas.SetStrokeColor(MakeColor(255, j, 255 - j))
        canvas.MoveTo(500 - float64(i), 500)
        canvas.LineTo(float64(i), 0)
        canvas.Stroke()
        i++
        j++
    }

    canvas.SaveToPNG("MyCoolPicture.png")
}