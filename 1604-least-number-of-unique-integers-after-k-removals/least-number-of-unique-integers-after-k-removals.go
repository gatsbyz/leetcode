func findLeastNumOfUniqueInts(arr []int, k int) int {
    recordMap := make(map[int]int)
    records := make([][]int, 0)
    for _, n := range arr {
        recordMap[n]++
    }
    for n, f := range recordMap {
        records = append(records, []int{n ,f})
    }
    slices.SortFunc(records, func(a, b []int) int {
        return cmp.Compare(a[1], b[1])
    })

    count := 0
    for _, record := range records {
        if k > 0 {
            if k > record[1] {
                k -= record[1]
                continue
            }  
            if k < record[1]{
                count++
            }
            k = 0
        } else {
            count++
        }
    }
    return count
}