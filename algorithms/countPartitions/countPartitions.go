package main

import (
    "fmt"
)

func countPartitions(L []int, k int) int {
    ret := 0
    m := make(map[int]int)
    for i := 0; i < len(L); i++ {
	if _, ok := m[L[i]]; ok {
	    ret++
        } else {
	    m[k - L[i]] = i
	}
	if L[i] == k / 2 {
	    ret++
	}
    }
    return ret
}

func main() {
    L := []int {2, 3, 4, 5, 7, 8}
    k := 6 
    fmt.Println(countPartitions(L, k))
}
