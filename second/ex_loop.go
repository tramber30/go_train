package main

import (
        "fmt"
        "math"
        "math/rand"
)

func Round(val float64, rank int) float64 {
        temp := val * math.Pow10(rank)
        if  float64(math.Trunc(temp + 0.5)) > float64(math.Trunc(temp)) {
                return float64(math.Trunc(temp + 0.5)) / float64(math.Pow10(rank))
        }
        return float64(math.Trunc(temp)) / float64(math.Pow10(rank))
}


func MySqrt(x, i int) float64 {
        z := 5.0
        for it := 0; it < i; it++ {
                high := math.Pow(z, float64(2)) - float64(x)
                low := 2 * z
                z = z - (high / low)
                fmt.Println(z)
                fmt.Printf("iteration %v, result: %v\n", it, z)
        }
        return z
}

func SelfSqrt(x int) (z float64, seed, it int) {
        seed = rand.Int()
        z = float64(seed)
        var res float64
        for it = 0; res != z ; it++ {
                res = z
                high := math.Pow(z, float64(2)) - float64(x)
                low := 2 * z
                z = z - (high / low)
//                fmt.Printf("iteration %v, result: %v\n", it, z)
        }
        return
}

func main() {
        i := 4076
        fmt.Println("square root of: ", i)
        for y := 0; y < 3; y++ {
               res, seed, it := SelfSqrt(i)
               fmt.Printf("resultat: %v, seed: %v, itération: %v\n", res, seed, it)
        }
        fmt.Println("Math:", math.Sqrt(float64(i)))
}

