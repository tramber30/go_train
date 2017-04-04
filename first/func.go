package main

import "fmt"

func add (x int, y int) int {
        return x + y
}

func swap (a, b string) (string, string) {
        return b, a
}

func split(sum int) (x, y int) {
        x = sum * 4 / 9
        y = sum -x
        return
}

func main () {
        a ,b := swap("Salut", "ca va?")
        fmt.Println(add(37,3))
        fmt.Println(a, b)
        fmt.Println(split(30))
}
