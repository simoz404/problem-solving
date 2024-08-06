func utopianTree(n int32) int32 {
    k := 1
   for i := 1; i <= int(n); i++ {
       if i % 2 == 0 {
           k++
       } else {
           k*=2
       }
   }
    return int32(k)
}
