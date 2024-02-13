func firstUniqChar(s string) int {
    list := [26]int{}
    for i := range s {
        list[s[i]-'a']++
    }
    for i := range s {
        if list[s[i]-'a'] == 1 {
            return i
        }
    }
    return -1
}