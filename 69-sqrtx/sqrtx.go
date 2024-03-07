
import "fmt"
func mySqrt(x int) int {
    l, r := 0, x+1

    for l<r {
        mid := l + (r - l) / 2
        if mid*mid == x {
            return mid
        } else if mid*mid < x {
            l = mid+1
        } else {
            r = mid
        }
    }
    return l-1

    // i := 1
    // for i <= x {
    //     if i*i == x {
    //         return i
    //     } else if i*i > x {
    //         return i-1
    //     }
    //     i++
    // }
    // return 0
}

// 4 -> 2
// 5 -> 2
// 6 -> 2
// 9 -> 3