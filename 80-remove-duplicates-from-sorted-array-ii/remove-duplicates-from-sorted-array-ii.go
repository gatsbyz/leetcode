func removeDuplicates(nums []int) int {
    if len(nums) < 2 {
        return len(nums)
    }
    pointer := 2
    for i:=2; i<len(nums) ; i++{
        if nums[i] != nums[pointer-2] {
            nums[pointer] = nums[i]
            pointer+=1
        }
    }
    return pointer
}