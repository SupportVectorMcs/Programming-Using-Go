/*
 * Author: Ronald He
 * Date: 09/20/2014
 */

/*=============================================================================
 * Cellular automate (CAs) are a simplified model of physics and complex 
 * systems that has been studied for many years. One-dimensional CAs are 
 * defined by the following rules:
 *     The "universe" is an infinitely long line of "cells."
 *     Each "cell" can either be filled or empty.
 *     The state of a cell i at time t depends on the states of cells i - 1, i,
 *         and i + 1 at time t - 1. The rule for how cell i depends on these
 *         cells defines the particular CA you are studying.
 * More details: http://www.ronhecmu.com/Programming-Using-Go/ 
 *===========================================================================*/

package main

import ("fmt" 
        "os"
        "math"
        "strconv"
        )

func main() {

    // make sure the number and type of arguments is correct
	if len(os.Args) != 4 {
        fmt.Println("Error: RULE, WIDTH, STEPS are needed.")
        return
    }

    // make sure rule is valid
    RULE, err := strconv.Atoi(os.Args[1])
    if err != nil {
        fmt.Println("Error: RULE must be an integer.")
        return
    } else if RULE < 0 {
        fmt.Println("Error: RULE must be positive.")
        return
    }

    // convert rule to numbers between 0 and 255
    if 8 == len(os.Args[1]) {
        var temp int
        for i := 0; i < 8; i++ {
            temp = temp + int(math.Pow(2, float64(i))) * (RULE & 1)
            RULE /= 10
        }
        RULE = temp
    } else if len(os.Args[1]) <= 3 { 
    } else {
        fmt.Println("Error: RULE is invalid.")
        return
    }

    // make sure WIDTH is valid
    WIDTH, err := strconv.Atoi(os.Args[2])
    if err != nil {
        fmt.Println("Error: WIDTH must be an integer.")
        return
    } else if WIDTH < 0 {
        fmt.Println("Error: WIDTH must be positive.")
        return
    }

    // make sure STEPS is valid
    STEPS, err := strconv.Atoi(os.Args[3])
    if err != nil {
        fmt.Println("Error: STEPS must be an integer.")
        return
    } else if STEPS < 0 {
        fmt.Println("Error: STEPS must be positive.")
        return
    }

    CA(RULE, WIDTH, STEPS)
}

// celluar automata function takes in RULE, WIDTH, STEPS and prints the pattern
func CA(RULE, WIDTH, STEPS int) {
    // print first row
    var i int
    // 1 to (WIDTH / 2 - 1)
    for i = 0; i < WIDTH / 2; i++ {             
        fmt.Print(" ")
    }
    // WIDTH / 2
    fmt.Print("#")
    // (WIDTH / 2 + 1) to (WIDTH - 1)
    for i = WIDTH / 2 + 1; i < WIDTH - 1; i++ { 
        fmt.Print(" ")
    }
    // WIDTH
    fmt.Println(" ")                            

    // first row
    row1 := make([]int, WIDTH + 2)
    row1[WIDTH / 2 + 1] = 1;

    // print following rows
    for j := 2; j < STEPS + 2; j++ {
        row1 = printNextRow(row1, RULE) 
    }
}

// get the next step instruction from rule
func getRule(RULE, n int) (int){
    var ret int
    ret = RULE & int(math.Pow(2, float64(8 - n)))
    if ret > 0 {
        ret = 1
    }
    return ret
}

// print next row
func printNextRow(row []int, RULE int) ([]int){
    nextRow := make([]int, 1)
    // analyze last row
    for k := 0; k < len(row) - 2; k++ {
        // get the rule index
        var ruleNo int = (1 - row[k]) * 4 + (1 - row[k + 1]) * 2 + 
            (1 - row[k + 2]) + 1
        // get the instruction from rule
        var digit = getRule(RULE, ruleNo)
        nextRow = append(nextRow, digit)

        if 1 == digit {
            fmt.Print("#")
        } else {
            fmt.Print(" ")
        }
    }
    fmt.Println()
    nextRow = append(nextRow, 0)

    // save present states for next step
    row = nextRow
    return row
}