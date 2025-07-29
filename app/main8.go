package main

import (
    "fmt"
    "local/mymath"
)

func main() {
    a, b := 5, 3
    
    fmt.Printf("%d + %d = %d\n", a, b, mymath.Add(a, b))
    fmt.Printf("%d * %d = %d\n", a, b, mymath.Multiply(a, b))
}