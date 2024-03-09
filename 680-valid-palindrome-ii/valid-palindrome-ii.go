func validPalindrome(s string) bool {
    left, right := 0, len(s)-1
    token := true
    reverted := false
    var revertLeft, revertRight int
    for left < right {
        fmt.Println(left, right, string(s[left]),  string(s[right]))
        if s[left] != s[right] {
            if token && s[left] == s[right-1] {
                revertLeft, revertRight = left+1, right
                right--
                token = false
            } else if token && s[left+1] == s[right] {
                revertLeft, revertRight = left, right-1
                left++
                token = false
            }  else {
                if !token && !reverted {
                    fmt.Println("here", left, right)
                    left, right = revertLeft, revertRight
                    reverted = true
                } else {
                    return false
                }
            }
        } else {
            left++
            right--
        }
    }
    return true
}
