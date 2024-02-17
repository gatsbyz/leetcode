// Define a job struct to hold the start time, end time, and profit of a job.
type job struct {
	startTime int
	endTime   int
	profit    int
}

// // Define a jobs type to implement the sort.Interface for a slice of job.
// type jobs []job

// // Len returns the number of elements in the collection.
// func (j jobs) Len() int {
// 	return len(j)
// }

// // Less reports whether the element with index i should sort before the element with index k.
// func (j jobs) Less(i, k int) bool {
// 	return j[i].endTime < j[k].endTime
// }

// // Swap swaps the elements with indexes i and k.
// func (j jobs) Swap(i, k int) {
// 	j[i], j[k] = j[k], j[i]
// }

// jobScheduling finds the maximum profit you can achieve by scheduling non-overlapping jobs.
func jobScheduling(startTime []int, endTime []int, profit []int) int {
	// Initialize jobArr with jobs using input startTime, endTime, and profit.
	jobArr := make([]job, len(startTime))
	for i := range startTime {
		jobArr[i] = job{startTime: startTime[i], endTime: endTime[i], profit: profit[i]}
	}

	// Sort jobs based on their end time to process them in chronological order.
	// sort.Sort(jobs(jobArr))

    slices.SortFunc(jobArr, func(a, b job) int {
        return cmp.Compare(a.endTime, b.endTime)
    })

	// Initialize a DP array to store the maximum profit achievable up to each job.
	dp := make([]int, len(jobArr))
    dp[0] = jobArr[0].profit  // Base case: The profit from the first job.

	// Iterate through each job to calculate maximum profit up to that job.
	for i := 1; i < len(dp); i++ {
        // Find the index of the last job that finishes before the current job starts.
        idx := findLastLessOrEqual(jobArr[i].startTime, jobArr)
        var prevProfit int
        if idx != -1 {
            prevProfit = dp[idx] // Profit from taking the last non-overlapping job.
        }
		// dp[i] is the max of taking the current job + profit till last non-overlapping job OR skipping the current job.
		dp[i] = maxInt(dp[i-1], prevProfit+jobArr[i].profit)
	}
	return dp[len(dp)-1]  // The last element in dp array contains the maximum profit.
}

// findLastLessOrEqual performs a binary search to find the index of the last job that ends before or exactly at 'target' start time.
func findLastLessOrEqual(target int, jobArr []job) int {
	l, r := 0, len(jobArr)-1
	res := -1
	for l <= r {
		m := (l + r) / 2
		if jobArr[m].endTime <= target {
			res = m // Update result and search in the right half for a possibly better candidate.
			l = m + 1
		} else {
			r = m - 1 // Search in the left half.
		}
	}
	return res
}

// maxInt returns the maximum of two integers.
func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
