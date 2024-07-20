func pageCount(n int32, p int32) int32 {
    if p%2 != 0 {
        p--
    }
    i1 := 0
    i2 := 0
    for i := 0; int32(i) < n; i++ {
        if int32(i) < p {
            i1++
        } else {
            i2++
        }
    }
    if i1 > i2 {
        return int32(i2/2)
    }
    return int32(i1/2)
}
