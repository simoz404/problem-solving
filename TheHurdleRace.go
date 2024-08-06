func hurdleRace(k int32, height []int32) int32 {
   max := k
   for _, v := range height {
       if v > max {
           max = v
       }
   }
   if max - k <= 0 {
       return 0
   }
   return max - k

}
