func findErrorNums(nums []int) []int {
    dict := map[int]int{}
    for i:=0;i<len(nums);i++ {
        dict[nums[i]]++
    }
    var first, second int
    for i:=1;i<=len(nums);i++ {
        e := dict[i]
        if e == 0 {
            second = i
        } else if e == 2 {
            first = i
        }
    }
    return []int{ first, second}
}