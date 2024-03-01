func maximumOddBinaryNumber(s string) string {
    countOne :=0
    result := ""
    for i:=0; i<len(s); i++ {
        if s[i] == '1' {
            countOne++
        }
    }
    for i:=0; i<countOne-1; i++ {
        result += "1"
    }
    for i:=0; i<len(s)-countOne; i++ {
        result += "0"
    }
    return result + "1"
}