package main

import "fmt"

type Node struct {
    value int
    next *Node
}

func createNode(v int) *Node {
    return &Node{value: v, next: nil}
}

func convertLinkedListToSlice(head *Node) []int {
    out := make([]int, 0)
    for p := head; p != nil; p = p.next {
	out = append(out, p.value)
    }
    return out
}

func insertSortLL(inList []int) []int {
    // create a linked list with just one item in it
    var head *Node = createNode(inList[0])
    
    // fot every remaining item
    for _, k := range inList[1:] {
        newNode := createNode(k)
	
	// walk down the linked list
	for prev, cur := (*Node)(nil), head; cur != nil; prev, cur = cur, cur.next {
	    
	    // if this is where we should insert
	    if cur.value > k {
		
		// if not at the start of the list
	        if prev != nil {
		    prev.next = newNode
		} else {
		    // otherwise, we're at the start of the list
		    head = newNode
		}
		newNode.next = cur
		break
	    }
	}
	// the first-smallest bug needs to be fixed
    }

    return convertLinkedListToSlice(head)
}

func main() {
    inList := []int{9, 2, 8, 3, 7, 4, 6, 5}
    fmt.Println(insertSortLL(inList))
}
