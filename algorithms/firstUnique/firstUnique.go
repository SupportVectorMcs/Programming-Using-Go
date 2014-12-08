package main

import (
    "fmt"
)

func firstUnique(L []int) int {
    for _, val1 := range L {
        count := 0
        for _, val2 := range L {
            if val1 == val2 {
	        count++
  	    }
        }
	if count == 1 {
            return val1
        }
    }  
    return 0
}

func main() {
    L1 := []int{2, 3, 4, 5, 2, 4, 5}
    L2 := []int{7, 8, 2, 8, 2, 7, 8}
    L3 := []int{8, 8, 8, 9, 8, 6, 8, 8}
    
    fmt.Println(firstUnique(L1))
    fmt.Println(firstUnique(L2))
    fmt.Println(firstUnique(L3))
}
