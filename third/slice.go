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

        sl := test[:4]

        fmt.Println(sl)
        fmt.Printf("slice type: %T, value: %V\n",sl, sl)
        fmt.Printf("array type: %T, value: %V\n",test, test)

}
