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

    // Second pass: Rebuild the string without invalid parentheses
    var builder strings.Builder
    invalidIndexes := make(map[int]bool) // Track the indexes of invalid parentheses
    for _, p := range stack {
        invalidIndexes[p.index] = true
    }

    for i := 0; i < len(s); i++ {
        if !invalidIndexes[i] {
            builder.WriteByte(s[i]) // Append valid characters to the builder
        }
    }

    return builder.String() // Return the valid string
}