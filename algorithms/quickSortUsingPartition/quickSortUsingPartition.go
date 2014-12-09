package main

import "fmt"

func partition(inList []int) int {
    pivot := inList[0]
    lastPos := len(inList) - 1

    // swap the first and last items
    inList[0], inList[lastPos] = inList[lastPos], inList[0]

    curIndex := 0
    for i := 0; i < lastPos; i++ {
	if inList[i] < pivot {
	    inList[i], inList[curIndex] = inList[curIndex], inList[i]
	    curIndex++
	}
    }
    inList[curIndex], inList[lastPos] = inList[lastPos], inList[curIndex]
    return curIndex
}

func quickSort(inList []int) {
    if len(inList) > 1 {
	p := partition(inList)
	quickSort(inList[:p])
	quickSort(inList[p + 1:])
    }
}

func main() {
    inList := []int{5, 7, 3, 2, 1, 0, 2, 8}
    quickSort(inList)
    fmt.Println(inList)
}
