package main

import (
    "fmt"
)

func merge(L1, L2 []int) []int {
    ret := make([]int,0)
    i, j := 0, 0
    for i != len(L1) || j != len(L2) {
        if i == len(L1) {
	    ret = append(ret, L2[j])
 	    j++
	    continue
	} else if j == len(L2) {
            ret = append(ret, L1[i])
	    i++
	    continue
	}
        if L1[i] < L2[j] {
 	    ret = append(ret, L1[i])
	    i++
        } else {
            ret = append(ret, L2[j])
	    j++
        }
    }
    return ret
}

func main() {
    L1 := []int {1, 3, 5, 7, 9}
    L2 := []int {2, 4, 6, 8, 10}
    fmt.Println(merge(L1, L2))
}
