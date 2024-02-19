func removeDuplicates(nums []int) int {
    pointer := 1
    current := nums[0]
    for i:=1;i<len(nums);i++ {
        if current != nums[i] {
            nums[pointer] = nums[i]
            pointer++
            current = nums[i]
        }
    }
    return pointer
}
