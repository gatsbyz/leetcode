func removeInterval(intervals [][]int, toBeRemoved []int) [][]int {
    var output [][]int
    for _, interval := range intervals {
        if interval[0] < toBeRemoved[0] {
            output = append(output, []int{interval[0], int(math.Min(float64(interval[1]), float64(toBeRemoved[0])))})
        }
        
        if interval[1] > toBeRemoved[1] {
            output = append(output, []int{int(math.Max(float64(interval[0]), float64(toBeRemoved[1]))), interval[1]})
        }
    }
    
    return output
}