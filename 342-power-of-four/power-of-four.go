func isPowerOfFour(n int) bool {
    for n > 0 && n % 4 == 0 {
        n = n/4
    }

    if n == 1 {
        return true
    }

    return false
}