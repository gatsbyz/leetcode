
type job struct {
	startTime int
	endTime   int
	profit    int
}

type jobs struct {
	arr []job
}

func (j *jobs) Len() int {
	return len(j.arr)
}

func (j *jobs) Less(i, k int) bool {
	return j.arr[i].endTime < j.arr[k].endTime
}

func (j *jobs) Swap(i, k int) {
	j.arr[i], j.arr[k] = j.arr[k], j.arr[i]
}

func jobScheduling(startTime []int, endTime []int, profit []int) int {
	jobArr := make([]job, len(startTime))
	for i := range startTime {
		jobArr[i] = job{startTime: startTime[i], endTime: endTime[i], profit: profit[i]}
	}
	sort.Sort(&jobs{arr: jobArr})
	dp := make([]int, len(jobArr))
    dp[0] = jobArr[0].profit
	for i :=1;i<len(dp);i++ {
        idx := findLastLessOrEqual(jobArr[i].startTime, jobArr)
        var prevProfit int
        if idx != -1 {
            prevProfit = dp[idx]
        }
		dp[i] = maxInt(dp[i-1], prevProfit+jobArr[i].profit)
	}
	return dp[len(dp)-1] 
}

func findLastLessOrEqual(target int, jobArr []job) int {
	l := 0
	r := len(jobArr)-1
	res := -1
	for l <= r {
		m := (l+r)/2
		if jobArr[m].endTime <= target {
			res = m
			l = m+1
		}else {
			r = m-1
		}
	}
	return res
}

func maxInt(a, b int) int {
	if a >b {
		return a
	}
	return b
}