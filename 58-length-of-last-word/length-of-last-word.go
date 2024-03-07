func lengthOfLastWord(s string) int {
    count := 0
    for i:=0; i<len(s); i++ {
        if count > 0 && s[len(s)-i-1] == ' ' {
            break
        } else if s[len(s)-i-1] != ' ' {
            count++
        }
    }
    return count
}