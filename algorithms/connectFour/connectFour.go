package main

import (
    "fmt"
)

func connectFour(board [][]bool) bool {
    for i := 0; i < len(board); i++ {
	for j := 0; j + 3 < len(board[0]); j++ {
	    if board[i][j] && board[i][j+1] && board[i][j+2] && board[i][j+3] {
		return true
	    }
	}
    }

    for i := 0; i < len(board[0]); i++ {
        for j := 0; j + 3 < len(board); j++ {
            if board[j][i] && board[j+1][i] && board[j+2][i] && board[j+3][i] {
                return true
            }
        }
    }

    for i := 0; i + 3 < len(board[0]); i++ {
	if board[i][i] && board[i+1][i+1] && board[i+2][i+2] && board[i+3][i+3] {
	    return true
	}
    }
    return false
}

func main() {
    board := [][]bool{
	{false, false, false, false, false, false, false, false, false, false, false, false},
	{false, true,  false, false, false, false, false, false, false, false, false, false},
        {false, false, false, false, false, false, false, false, false, false, false, false},
        {false, false, true,  false, false, false, false, false, false, false, false, false},

        {false, false, false, true,  false, false, false, false, false, false, false, false},
        {false, false, false, true,  false, false, false, false, false, false, false, false},
        {false, false, false, true,  false, false, false, false, false, false, false, false},
        {false, false, false, false, false, false, false, false, false, false, false, false},

        {false, false, false, false, false, false, false, false, true,  false, false, false},
        {false, false, false, false, false, false, false, false, false, true,  false, false},
        {false, false, false, false, false, false, false, false, false, false, true,  false},
        {false, false, false, false, false, false, false, false, false, false, false, true}}
    fmt.Println(connectFour(board))
}
