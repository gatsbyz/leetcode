func maxFrequencyElements(nums []int) int {
    sort.Ints(nums)
    maxFreq := 0
    currentFreq := 0
    currentValue := 0
    totalFrequency := 0
    for i:=0;i<len(nums);i++ {
        if currentValue == nums[i] {
            currentFreq++
            if currentFreq == maxFreq {
                totalFrequency += currentFreq
            } else if currentFreq > maxFreq {
                maxFreq = currentFreq
                totalFrequency = maxFreq
            }
        } else {
            currentValue = nums[i]
            currentFreq = 1
            if maxFreq < currentFreq {
                maxFreq = currentFreq
                totalFrequency = maxFreq
            } else if currentFreq == maxFreq {
                totalFrequency += 1
            }
        }
    }
    return totalFrequency
}