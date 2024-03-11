func customSortString(order string, s string) string {
    orderMap := map[byte]int{} // char -> index
    result := make([]string, 26) // 
    resultString:= ""

    for index, _ := range order {
        orderMap[order[index]] = index
    }

    for index, _ := range s {
        if _, exist := orderMap[s[index]]; exist {
            result[orderMap[s[index]]] += string(s[index])
            // delete(orderMap, s[index])
        } else {
            resultString += string(s[index])
        }
    }

    for _, char := range result {
        if char != "" {
            resultString += string(char)
        }
    }
    return resultString
}

// c 0 b 1 a 2 //  a b c d

// bcafg // abc d