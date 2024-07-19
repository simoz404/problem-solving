func dayOfProgrammer(year int32) string {
    var s []rune
    i := year
    for i > 0 {
        s = append([]rune{(i%10)+48}, s...)
        i/=10
    }
    if year == 1918 {
        return "26.09.1918"
    }
    if year >= 1919 {
        if year % 4 == 0 && year % 100 != 0 || year % 400 == 0 {
        return "12.09." + string(s)
    }
     return "13.09." + string(s)
    }
     if year % 4 == 0 {
        return "12.09." + string(s)
    }
     return "13.09." + string(s)
    // Write your code here
    
}
