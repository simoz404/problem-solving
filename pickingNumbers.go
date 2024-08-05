package main

import "fmt"

func pickingNumbers(a []int32) int32 {
    var s []int32
    var s1[][]int32
    for i := 0; i < len(a); i++ {
        for j := 0; j < len(a); j++ {
            if a[i] - a[j] == 1  {
                s = append(s, a[j])
            } else if a[i] - a[j] == 0  {
                 s = append(s, a[j])
            }
        }
        s1 = append(s1, s)
        s = []int32{}
    }
    l := 0
    for _, v := range s1 {
        if len(v) > l {
            l = len(v)
        }
    }
    return int32(l)
}

func main() {
	fmt.Println(pickingNumbers([]int32{1, 2, 2, 3, 1, 2}))
}
