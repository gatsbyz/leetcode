type paren struct {
    sign  byte // Character of the parenthesis, either '(' or ')'
    index int  // Position of the parenthesis in the string
}

func minRemoveToMakeValid(s string) string {
    stack := make([]paren, 0) // Stack to keep track of parentheses

    // First pass: Identify invalid parentheses
    for i := 0; i < len(s); i++ {
        switch s[i] {
        case '(':
            // Push opening parenthesis onto the stack
            stack = append(stack, paren{sign: s[i], index: i})
        case ')':
            // If there's a matching opening parenthesis, pop it off the stack
            if len(stack) > 0 && stack[len(stack)-1].sign == '(' {
                stack = stack[:len(stack)-1]
            } else {
                // If there's no matching opening, mark this closing parenthesis as invalid
                stack = append(stack, paren{sign: s[i], index: i})
            }
        }
    }

    // If the stack is empty, all parentheses are valid
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