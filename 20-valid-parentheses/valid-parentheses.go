func isValid(s string) bool {
    queue := []byte{}
    for i:=0; i<len(s); i++ {
        if len(queue) != 0 && (queue[len(queue)-1] == '(' && s[i] == ')' || queue[len(queue)-1] == '[' && s[i] == ']' || queue[len(queue)-1] == '{' && s[i] == '}') {
            queue = queue[:len(queue)-1]
        } else {
            queue = append(queue, s[i])
        }
    }

    return len(queue) == 0
}