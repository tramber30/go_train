package main

import "fmt"

func main() {
        var sum int
        for i:=0; i < 10; i++ {
                sum += i
                fmt.Println(i)
        }
        fmt.Println(sum)
}
