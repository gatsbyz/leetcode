func mySqrt(x int) int {
    i := 1
    for i <= x {
        if i*i == x {
            return i
        } else if i*i > x {
            return i-1
        }
        i++
    }
    return 0
}

// 4 -> 2
// 5 -> 2
// 6 -> 2
// 9 -> 3