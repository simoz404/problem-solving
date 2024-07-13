package main

import "fmt"
//imput example w = hello && decor = * and outer = 2
func main() {
    var w, decor string
    var outer int
    fmt.Scan(&w, &decor, &outer)
    k := outer
    i :=  outer + outer + 1
    j := outer*2 + len(w) + 2
    for i > 0 {
        if i == outer+1 {
            for k > 0 {
                fmt.Print(decor)
                k--
            }
            k = outer
            fmt.Print(" " + w + " ")
            for k > 0 {
                fmt.Print(decor)
                k--
            }
            fmt.Println()
        } else {
            for j > 0 {
                fmt.Print(decor)
                j--
            }
            fmt.Println()
            j = outer*2 + len(w) + 2
        }
        i--
    }
    fmt.Println()// Write answer to stdout
}
