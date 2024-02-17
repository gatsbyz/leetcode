type Job struct {
    start, end, profit int
}

// Helper function to perform binary search on the jobs slice.
// Finds the last job that finishes before the specified start time.
func findLastNonOverlappingJob(jobs []Job, startTime int) int {
    lo, hi := 0, len(jobs)-1
    for lo <= hi {
        mid := lo + (hi-lo)/2
        if jobs[mid].end <= startTime {
            if mid == len(jobs)-1 || jobs[mid+1].end > startTime {
                return mid
            }
            lo = mid + 1
        } else {
            hi = mid - 1
        }
    }
    return -1 // Return -1 if no non-overlapping job is found.
}

// Function to calculate the maximum profit from job scheduling.
func jobScheduling(startTime []int, endTime []int, profit []int) int {
    n := len(startTime)
    jobs := make([]Job, n)
    for i := 0; i < n; i++ {
        jobs[i] = Job{startTime[i], endTime[i], profit[i]}
    }

    // Sort jobs based on their end times.
    sort.Slice(jobs, func(i, j int) bool {
        return jobs[i].end < jobs[j].end
    })

    // DP array to store the maximum profit up to each job.
    dp := make([]int, n)
    dp[0] = jobs[0].profit

    for i := 1; i < n; i++ {
        // Profit including the current job.
        includeProfit := jobs[i].profit
        // Find the last job that does not overlap with the current one.
        lastIndex := findLastNonOverlappingJob(jobs, jobs[i].start)
        if lastIndex != -1 {
            includeProfit += dp[lastIndex]
        }
        // Maximum profit up to the current job is either by including the current job
        // or by excluding it (which is the max profit up to the previous job).
        dp[i] = max(dp[i-1], includeProfit)
    }

    // The last element in dp array contains the maximum profit for all jobs.
    return dp[n-1]
}

// Helper function to find the maximum of two integers.
func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}