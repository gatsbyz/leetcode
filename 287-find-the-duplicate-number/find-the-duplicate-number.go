func findDuplicate(nums []int) int {
    for _, num := range nums {
        if nums[(num%len(nums))-1] / len(nums) > 0 {
            return num%len(nums)
        }
        nums[(num%len(nums))-1] += len(nums)
    }

    return -1
}