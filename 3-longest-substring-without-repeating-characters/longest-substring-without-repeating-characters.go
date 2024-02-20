func lengthOfLongestSubstring(s string) int {
    stack := []string{}
    maxCount := 0
    countStr := map[string]int{}
    for index:=0;index<len(s);index++ {
        fmt.Println(countStr[string(s[index])])
        if len(stack) != 0 && countStr[string(s[index])] > 0 {
            for i:=0;i<len(stack);i++ {
                if stack[i] == string(s[index]) {
                    stack = stack[i+1:]
                    countStr[string(s[index])]=0
                    break
                }
                countStr[string(s[index])]=0
            }
        }
        stack = append(stack, string(s[index]))
        countStr[string(s[index])]++
        if maxCount < len(stack) {
            maxCount = len(stack)
        }
        fmt.Println(len(stack), stack, maxCount, countStr[string(s[index])])
    }
    return maxCount
}