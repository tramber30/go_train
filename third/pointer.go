package main

import "fmt"

type Vertex struct {
        x int
        Y int
}

func main() {
        v := Vertex{1, 2}
        fmt.Println(v)
        v.x = 5
        fmt.Println(v)
        p := &v
        fmt.Println("test pointer: ", p)
}


