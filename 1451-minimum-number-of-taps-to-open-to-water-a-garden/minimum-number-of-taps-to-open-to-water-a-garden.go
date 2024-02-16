func minTaps(n int, ranges []int) int {
  // Initialize a slice to hold the start and end points of each tap's range.
    intervals := make([][2]int, 0)
    for i, r := range ranges {
        // Skip taps with a range of 0, as they cannot cover any area.
        if r == 0 {
            continue
        }
        // Calculate the left and right bounds of the tap's coverage.
        // Ensure the left bound is not negative.
        left := i - r
        if left < 0 {
            left = 0
        }
        // The right bound simply adds the range to the tap's position.
        right := i + r
        // Add the interval to our list of intervals.
        intervals = append(intervals, [2]int{left, right})
    }

    // Sort the intervals based on their start position.
    // If two intervals start at the same position, the longer interval comes first.
    // This ensures that we always choose the interval that covers the most ground when starting at the same point.
    sort.Slice(intervals, func(i, j int) bool {
        if intervals[i][0] == intervals[j][0] {
            return intervals[i][1] > intervals[j][1]
        }
        return intervals[i][0] < intervals[j][0]
    })

    // Initialize variables to track the current end of the covered area,
    // the farthest point we can cover with the next tap, and the index of the interval.
    taps := 0
    end := 0
    farthest := 0
    i := 0

    // Continue until we've covered the entire garden.
    for end < n {
        updated := false // Flag to check if we can extend the coverage.
        // Iterate over intervals starting from the current position.
        for i < len(intervals) && intervals[i][0] <= end {
            // Update the farthest reach if the current interval extends it.
            if intervals[i][1] > farthest {
                farthest = intervals[i][1]
                updated = true
            }
            i++
        }

        // If we couldn't extend the coverage, the garden cannot be fully watered.
        if !updated {
            return -1
        }

        // Count the tap we just used to extend our coverage.
        taps++
        // Update the end to the farthest point covered by the tap.
        end = farthest

        // If we've covered the entire garden, return the number of taps used.
        if end >= n {
            return taps
        }
    }

    // If we exit the loop without covering the garden, return -1.
    // This line should never be reached due to the logic in the loop.
    return -1
}