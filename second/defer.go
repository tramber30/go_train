package main

import (
        "fmt"
)

func main() {
        f()
        fmt.Println("returned normally from f")
}

func f() {
        defer func() {
                if r := recover(); r != nil {
                        fmt.Println("f: Recovered in f", r)
                }
        }()
        fmt.Println("f: Calling g")
        g(0)
        fmt.Println("f: Returned normally from g")
}

func g(i int) {
        if i > 3 {
                fmt.Println("g: Panicking!")
                panic(fmt.Sprintf("%v", i))
        }
        defer fmt.Println("g: Defer in g", i)
        fmt.Println("g: Printing in g", i)
        g(i + 1)
}
