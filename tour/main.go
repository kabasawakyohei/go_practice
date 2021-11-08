package main

import (
    "fmt"
)

func main() {
    fmt.Println(adder(1, 2))
}

func adder(a, b int) int {
    return a + b
}