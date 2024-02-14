func rearrangeArray(nums []int) []int {
    rearranged := make([]int, len(nums))
    pos := 0
    neg := 1
    for i := range nums{
        if nums[i] > 0 {
            rearranged[pos] = nums[i]
            pos+=2
        } else {
            rearranged[neg] = nums[i]
            neg+=2
        }
    }
    return rearranged
}