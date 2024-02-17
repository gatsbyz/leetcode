// Define a Job struct to hold the start, end time, and profit of each job.
type Job struct {
    start, end, profit int
}

// maxProfit calculates the maximum profit you can achieve given job schedules.
func jobScheduling(startTime, endTime, profit []int) int {
    n := len(startTime)
    jobs := make([]Job, n)
    
    // Populate jobs with start, end, and profit info.
    for i := 0; i < n; i++ {
        jobs[i] = Job{startTime[i], endTime[i], profit[i]}
    }
    
    // Sort jobs based on their end time.
    sort.Slice(jobs, func(i, j int) bool {
        return jobs[i].end < jobs[j].end
    })

    // dp array to store the maximum profit up to each job.
    dp := make([]int, n)
    dp[0] = jobs[0].profit

    // Iterate over jobs to fill dp with maximum profit values.
    for i := 1; i < n; i++ {
        // Initialize includeProfit with the profit of taking the current job.
        includeProfit := jobs[i].profit
        // Find profit from the last non-conflicting job.
        for j := i - 1; j >= 0; j-- {
            if jobs[j].end <= jobs[i].start {
                includeProfit += dp[j]
                break // Found the last non-conflicting job.
            }
        }
        // Maximum profit at i is max of taking the current job or skipping it.
        dp[i] = max(dp[i-1], includeProfit)
    }

    return dp[n-1] // The last element of dp contains the maximum profit.
}

// max returns the maximum of two integers.
func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
