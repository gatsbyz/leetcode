func findThePrefixCommonArray(A []int, B []int) []int {
	cnts, cnt, ans := make([]bool, len(A)+1), 0, []int{}
	for i, a := range A {
		if cnts[a] {
			cnt++
		} else {
			cnts[a] = true
		}
		if cnts[B[i]] {
			cnt++
		} else {
			cnts[B[i]] = true
		}
		ans = append(ans, cnt)
	}
	return ans
}