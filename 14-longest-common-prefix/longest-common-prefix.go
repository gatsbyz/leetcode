
import "fmt"
func longestCommonPrefix(strs []string) string {
    n := 0
    prefix := ""
    for loop := true; loop; {
        var char byte
        for _, str := range strs {
            if n == len(str) {
                return prefix
            }
            if char == 0 {
                char = str[n]
            }
            if char != str[n] {
                return prefix
            }
        }
        prefix += string(char)
        n++
    }
    return prefix
}