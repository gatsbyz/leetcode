func divideArray(nums []int, k int) [][]int {
    sort.Slice(nums, func(i, j int) bool {
        return nums[i] < nums[j]
    })
    array := [][]int{}
    for index, _ := range nums {
        // 0-0,1,2 1-3,4,5
        arrayIndex := index / 3
        subIndex := index % 3
        if subIndex == 0 {
            array = append(array, make([]int, 0))
        }
        array[arrayIndex] = append(array[arrayIndex], nums[index])
        if (subIndex == 2 || index == len(nums)-1) && array[arrayIndex][len(array[arrayIndex])-1]-array[arrayIndex][0] > k {
            return [][]int{}
        }
    }
    return array
}