func pivotInteger(n int) int {
    sum := 0
    for i:=1;i<=n;i++ {
        sum += i
    }
    backSum := 0
    for i:=n;i>=0;i-- {
        backSum += i
        if sum == backSum {
            return i
        } else if sum < backSum {
            break
        }
        sum -= i
    }
    return -1
}