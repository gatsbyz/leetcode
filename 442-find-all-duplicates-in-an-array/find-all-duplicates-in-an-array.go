func findDuplicates(nums []int) []int {
    dupes := []int{}
    for _, num := range nums {
        if (nums[(num-1)%len(nums)]-1) / len(nums) == 1 {
            dupes = append(dupes, ((num-1) % len(nums)) + 1)
            nums[(num-1)%len(nums)] += len(nums)
        } else if (nums[(num-1)%len(nums)]-1) / len(nums) == 0 {
            nums[(num-1)%len(nums)] += len(nums)
        }
    }

    return dupes
}