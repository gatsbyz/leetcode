func rangeBitwiseAnd(left int, right int) int {
    // Shift both numbers to the right until you find something in common in the prefix of both numbers
    count := 0
    for left != right {
        left >>= 1
        right >>= 1
        count += 1
    }
    
    return left << count // Shift back to get the prefix in the original bit positions it appeared in
}