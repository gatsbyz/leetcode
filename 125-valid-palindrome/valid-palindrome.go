func isPalindrome(s string) bool {
    left, right := 0, len(s)-1
    for left < right {
        if !isAlphanumeric(s[left]) {
            left++
        } else if !isAlphanumeric(s[right]) {
            right--
        } else if strings.ToLower(string(s[left])) != strings.ToLower(string(s[right])) {
            return false
        } else {
            left++
            right--
        }
    }
    return true
}

func isAlphanumeric(char byte) bool {
    return 'a' <= char && char <= 'z' || '0' <= char && char <= '9' || 'A' <= char && char <= 'Z'
}