type paren struct {
    sign byte
    index int
}

func minRemoveToMakeValid(s string) string {
    i := 0
    stack := make([]paren, 0)
    for i < len(s) {
        if s[i] == ')' || s[i] == '(' {
            if s[i] == ')' && len(stack) > 0 && stack[len(stack)-1].sign == '(' { 
                stack = stack[:len(stack)-1]
            } else {
                stack = append(stack, paren{
                    sign: s[i],
                    index: i,
                })
            }  
        }
        i++
    }
    if len(stack) == 0 {
        return s
    }
    stackIndex := 0
    valid := ""
    fmt.Println(stack)
    for i:=0;i<len(s);i++ {
        if  stackIndex < len(stack) && stack[stackIndex].index == i {
            fmt.Println(stackIndex)
            stackIndex++
            continue
        }
        valid += string(s[i])
    }
    return valid
}
