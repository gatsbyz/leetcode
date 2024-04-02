func isIsomorphic(s string, t string) bool {
    return isIsomorphicHelper(s, t) && isIsomorphicHelper(t, s)
}
func isIsomorphicHelper(s string, t string) bool {
    index := 0
    ref := map[byte]byte{}
    for index < len(s) {
        if s[index] == t[index] {
            if _, exists := ref[s[index]]; !exists {
                ref[s[index]] = t[index]
                index++
                continue
            } else {
            if ref[s[index]] == t[index] {
                index++
                continue
            }
            return false
        }
        }
        if _, exists := ref[s[index]]; !exists {
            ref[s[index]] = t[index]
        } else {
            if ref[s[index]] == t[index] {
                index++
                continue
            }
            return false
        }
        index++
    }
    return true
}