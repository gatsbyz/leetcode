func moveZeroes(nums []int)  {
    if len(nums) == 1 {
        return
    }
    var pointer int
    for index := range nums {
        if pointer <= index {
            pointer = index+1
            if pointer == len(nums) {
                return
            }
        }
        if nums[index] != 0 {
            continue
        } else {
            for nums[pointer] == 0 {
                pointer++
                if pointer == len(nums) {
                    return
                }
            }
            nums[index] = nums[pointer]
            nums[pointer] = 0
        }
    }
}