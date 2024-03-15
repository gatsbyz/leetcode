func reverseWords(s string) string {
    splits := strings.Split(s, " ")
    res := ""

    for index, word := range splits {
        if index != 0 {
            res += " "
        }
        res += reverse(string(word))
    }

    return res
}

func reverse(s string) string {
    res := ""
    for i:=0;i<len(s);i++ {
        res = string(s[i]) + res
    }

    return res
}