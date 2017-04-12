package main

import "fmt"


func main () {
        test := [4]string {
                "tata",
                "toto",
                "tutu",
                "titi",
        }
        fmt.Println(test)
        sl := test[0:1]


        fmt.Println(sl)
//      fmt.Printf("slice type: %T, value: %V\n",sl, sl)
//        fmt.Printf("array type: %T, value: %V\n",test, test)
        fmt.Printf("slice lgth: %d, capacity: %d, content %v\n", len(sl), cap(sl), sl)
        sl = sl[0:3]

        fmt.Printf("slice lgth: %d, capacity: %d, content %v\n", len(sl), cap(sl), sl)       
         sl = test[1:4]
               fmt.Printf("slice lgth: %d, capacity: %d, content %v\n", len(sl), cap(sl), sl)       
}
