func findDuplicate(nums []int) int {
    find := map[int]bool{}
    for _, num := range nums {
        if _, ok := find[num]; !ok {
            find[num] = true
        } else {
            return num
        }
    }

    return -1
}