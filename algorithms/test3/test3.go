package main

import (
    "fmt"
)

func main() {
    var a int
    setA(&a)
    fmt.Println(a)
}

func setA(a *int) {
    *a = 42
}
