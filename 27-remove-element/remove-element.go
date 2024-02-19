func removeElement(nums []int, val int) int {
    if len(nums) == 0 {
        return 0
    }
    startIndex :=0
    endIndex := len(nums) - 1
    for startIndex < endIndex {
        if nums[startIndex] != val {
            startIndex++
        } else {
            swap(nums, startIndex, endIndex)
            endIndex--
        }
    }
    fmt.Println(nums, startIndex)
    if nums[startIndex] == val {
        startIndex -= 1
    }
    return startIndex + 1
}

func swap(nums []int, i, j int) {
    temp := nums[i]
    nums[i] = nums[j]
    nums[j] = temp
}