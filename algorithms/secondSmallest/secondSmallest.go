package main

import "fmt"

func secondSmallest(L []int) int {
    if L[0] < L[1] {
	prev, cur := L[1], L[0]
    } else {
	prev, cur := L[0], L[1]
    }
    
    for i := 2; i < len[L]; i++ {
	if L[i] < cur {
	    cur = L[i]
	    prev = cur
	} else if L[i] < prev {
	    prev = L[i]
	}
    }
    return prev
}

func main() {
    L := []int{1, 6, 2, 7, 9, 8}
    fmt.Println(secondSmallest(L))
}
