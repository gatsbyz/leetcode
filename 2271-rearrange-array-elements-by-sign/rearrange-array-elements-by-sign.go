func rearrangeArray(nums []int) []int {
    rearranged := make([]int, len(nums))
    pos := 0
    neg := 1
    for _, num := range nums{
        if num > 0 {
            rearranged[pos] = num
            pos+=2
        } else {
            rearranged[neg] = num
            neg+=2
        }
    }
    return rearranged
}