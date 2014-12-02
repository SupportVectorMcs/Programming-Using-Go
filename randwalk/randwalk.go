/*
 * Author: Ronald He
 * Date: 09/16/2014
 */

/*=============================================================================
 * Random Walks
 *
 * Implement a complete program that simulates a random walk in a width x 
 * height field. It is runnable with the following command:
 *     go run randwalk.go WIDTH HEIGHT STEPSIZE NUMBER-OF-STEPS SEED
 * For example: 
 *     go run randwalk.go 10 100.1 0.2 1000 39481
 * should simulate a random walk of 1000 steps in a field of width 10 and 
 * height 100, where each step is of distance 0.2. The final command line
 * parameter is the seed for the random number generator.
 *===========================================================================*/

package main
import (
    "math/rand"
    "math"
    "fmt"
    "os"
    "strconv"
    )

func main() {
    if len(os.Args) != 6 {
        fmt.Println("Error: WIDTH, HEIGHT, STEPSIZE, NUMBER_OF_STEPS and SEED are needed.")
        return
    }

    WIDTH, err := strconv.ParseFloat(os.Args[1], 10)
    if err != nil {
        fmt.Println("Error: WIDTH must be a real number.")
        return
    } else if WIDTH < 0 {
        fmt.Println("Error: WIDTH must be positive.")
        return
    }

    HEIGHT, err := strconv.ParseFloat(os.Args[2], 10)
    if err != nil {
        fmt.Println("Error: HEIGHT must be a real number.")
        return
    } else if HEIGHT < 0 {
        fmt.Println("Error: HEIGHT must be positive.")
        return
    }

    STEPSIZE, err := strconv.ParseFloat(os.Args[3], 10)
    if err != nil {
        fmt.Println("Error: STEPSIZE must be a real number.")
        return
    } else if STEPSIZE < 0 {
        fmt.Println("Error: STEPSIZE must be positive.")
        return
    }

    NUMBER_OF_STEPS, err := strconv.Atoi(os.Args[4])
    if err != nil {
        fmt.Println("Error: NUMBER-OF-STEPS must be an integer.")
        return
    } else if NUMBER_OF_STEPS < 0 {
        fmt.Println("Error: NUMBER-OF-STEPS must be positive.")
        return
    }

    SEED, err := strconv.ParseInt(os.Args[5], 10, 64)
    if err != nil {
        fmt.Println("Error: SEED must be an integer.")
        return
    } else if SEED < 0 {
        fmt.Println("Error: SEED must be positive.")
        return
    }

    randomWalk(WIDTH, HEIGHT, STEPSIZE, NUMBER_OF_STEPS, SEED)
}

// make sure position within [0, n) x [0, n)
func inField(coord, n float64) bool {
    return coord >= 0 && coord < n
}

// generate random movements
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

func randomWalk(WIDTH, HEIGHT, STEPSIZE float64, NUMBER_OF_STEPS int, SEED int64) {
    
    // the seed for the random number generator
    rand.Seed(SEED)

    var x, y = WIDTH / 2, HEIGHT / 2
    fmt.Println(x, y)
    for i := 0; i < NUMBER_OF_STEPS; i++ {
        x, y = randStep(x, y, STEPSIZE, WIDTH, HEIGHT)
        fmt.Println(x,y)
    }
    var distanceA float64 = math.Pow((x - WIDTH / 2), 2)
    var distanceB float64 = math.Pow((y - HEIGHT / 2), 2)
    fmt.Println("Distance =", math.Sqrt(distanceA + distanceB))
}