func numWays(n int, k int) int {
    if n == 1 {
        return k
    }

    a := k
    b := k*k
    for i := 0;  i < n -2; i++ {
        c := (k-1) * (a + b)
        a = b
        b = c
    }

    return b
}