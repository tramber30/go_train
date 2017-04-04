package main

import (
        "fmt"
        "math"
)

func Round(val float64, rank int) float64 {
        temp := val * math.Pow10(rank)
        if  float64(math.Trunc(temp + 0.5)) > float64(math.Trunc(temp)) {
                return float64(math.Trunc(temp + 0.5)) / float64(math.Pow10(rank))
        }
        return float64(math.Trunc(temp)) / float64(math.Pow10(rank))
}


func main() {
        const val = 0.123456006
        fmt.Println("const: ", val)
        fmt.Println("round3:", Round(val, 3))
        fmt.Println("round4:", Round(val, 4))
        fmt.Println("round5:", Round(val, 5))
        fmt.Println("round6:", Round(val, 6))
        fmt.Println("round7:", Round(val, 7))
        fmt.Println("round8:", Round(val, 8))
        fmt.Println("round11:", Round(val, 11))
}
