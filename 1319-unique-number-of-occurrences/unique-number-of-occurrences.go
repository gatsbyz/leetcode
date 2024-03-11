func uniqueOccurrences(nums []int) bool {
     record := map[int]int{}
    for i:=0;i<len(nums);i++ {
        record[nums[i]]++
    }
    slice := []int{}
    for _, rec := range record {
        slice = append(slice, rec)
    }
    sort.Ints(slice)

    for i:=1;i<len(slice);i++{
        if slice[i-1] == slice[i] {
            return false
        }
    }

    return true
}