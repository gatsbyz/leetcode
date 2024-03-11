func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }
    counter := map[byte]int{}
    for i:=0;i<len(s);i++ {
        counter[s[i]]++
    }
    for i:=0;i<len(t);i++ {
        counter[t[i]]--
    }

    for _, count := range counter {
        if count != 0 {
            return false
        }
    }

    return true
}