func smallestDivisor(nums []int, threshold int) int {
    left := 0
    right := 1000000

    for left <= right {
        mid :=  left + (right-left)/2
        d := 0.0

        for _, n := range nums {
            d += math.Ceil(float64(n)/float64(mid))
        }

        if d <= float64(threshold) {
            right = mid -1
        } else {
            left = mid + 1
        }
    }

    return left
}
// 1,2,5,9 - 6
// 1 1+2+5+9 = 17
// 2 1+1+3+5 = 10
// 3 1+1+2+3 = 7
// 4 1+1+2+3 = 7
// 5 1+1+1+2 = 5

