func minTaps(n int, ranges []int) int {
    intervals := make([][2]int, 0)
    for i, r := range ranges {
        if r == 0 {
            continue // Skip taps that can't cover any area
        }
        left := i - r
        if left < 0 {
            left = 0
        }
        right := i + r
        intervals = append(intervals, [2]int{left, right})
    }

    // Sort intervals by start position
    sort.Slice(intervals, func(i, j int) bool {
        if intervals[i][0] == intervals[j][0] {
            return intervals[i][1] > intervals[j][1]
        }
        return intervals[i][0] < intervals[j][0]
    })

    taps := 0
    end := 0
    farthest := 0
    i := 0

    for end < n {
        // Find the farthest-reaching interval from the current position
        updated := false
        for i < len(intervals) && intervals[i][0] <= end {
            if intervals[i][1] > farthest {
                farthest = intervals[i][1]
                updated = true
            }
            i++
        }

        if !updated { // Cannot reach further
            return -1
        }

        taps++ // Open a tap
        end = farthest // Extend the covered area

        if end >= n { // Check if the garden is fully covered
            return taps
        }
    }

    return -1 // Should not reach here if the garden can be covered
}