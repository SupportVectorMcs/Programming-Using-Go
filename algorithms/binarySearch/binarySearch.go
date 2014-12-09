package main

import "fmt"

func binarySearch(inList []int, k int) (int, bool) {
    left, right := 0, len(inList) - 1
    for left <= right {
	mid := left + (right - left) / 2
	if inList[mid] == k {
	   return mid, true
	} else if inList[mid] > k {
	    right = mid - 1
	} else if inList[mid] < k {
	    left = mid + 1
	}
    }
    return 0, false
}

func main() {
    inList := []int{1, 5, 6, 7, 8, 12, 14, 15, 18, 21, 33}
    fmt.Println(binarySearch(inList, 8))
}
