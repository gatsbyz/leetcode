func twoSum(nums []int, target int) []int {
    calc := map[int]int{} // key num, value index

    for i, num := range nums {
        find := target - num
        if _, exists := calc[find]; exists {
            return []int{calc[find], i}
        }
        calc[num] = i
    }

    return []int{}

    // start, end := 0, len(nums)-1
    // save := append([]int{}, nums...)
    // sort.Ints(save)
    // for start < end {
    //     if save[start] + save[end] == target {
    //         break
    //     } else if save[start] + save[end] < target {
    //         start++
    //     } else {
    //         end--
    //     }
    // }
    // res := []int{}
    // for i, val := range nums {
    //     if val == save[start] || val == save[end] {
    //         res = append(res, i)
    //     }
    // }
    // return res
}