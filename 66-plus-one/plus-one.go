
func plusOne(digits []int) []int {
    carry := 0
    for i:=len(digits)-1; i>=0; i-- {
        tempValue := digits[i] + carry
        if i == len(digits)-1 {
            tempValue++
        }
        digits[i] = tempValue % 10
        carry = tempValue / 10
    }
    if carry > 0 {
        digits = append([]int{1}, digits...)
    }
    return digits
}