func lengthOfLongestSubstring(s string) int {
    stack := [200]bool{}
    maxCount := 0
    length := 0
    pointer := 0
    for index:=0;index<len(s);index++ {
        fmt.Println(s[index])
        if stack[s[index]] {
            for ; stack[s[index]] ; pointer++ {
                stack[s[pointer]] = false
                length--
            }
        }
        stack[s[index]] = true
        length++
        if maxCount < length {
            maxCount = length
        }
    }
    return maxCount
}