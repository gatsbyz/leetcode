func largestPerimeter(nums []int) int64 {
	// Hint fromt the Leetcode's solution.
	var sum, ans int64
	ans = -1
	sum = 0
	sort.Ints(nums)
	for _, v := range nums {
		if int64(v) < sum {
			ans = int64(v) + sum
		}
		sum += int64(v)
	}
	return ans
}