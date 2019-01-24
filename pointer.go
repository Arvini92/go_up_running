package main

import "fmt"



func swap(x *int, y *int) {
*x, *y = *y, *x
}




func main() {
    x, y := 2, 3

fmt.Println("x =", x, ", y =", y)
swap(&x, &y)
fmt.Println("x =", x, ", y =", y)


    var p *int
    if p != nil {
        fmt.Println("Value of p:", *p)
    } else {
        fmt.Println("p is nil")
    }

    var v int  = 42
    p = &v
    if p != nil {
        fmt.Println("Value of p:", *p)
    } else {
        fmt.Println("p is nil")
    }

    var value1 float64 =42.13
    pointer1 := &value1
    fmt.Println("Value 1:", *pointer1)

    *pointer1 = *pointer1 / 31
    fmt.Println("Value 1:", *pointer1)
    fmt.Println("Value 1:", value1)
}