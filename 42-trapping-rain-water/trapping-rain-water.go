func trap(height []int) int {
    maxLeft, maxRight := 0, 0
    left, right := 0, len(height)-1
    trapped := 0
    for left < right {
        fmt.Println(trapped, left, right)
        if height[left] > maxLeft {
            maxLeft = height[left]
        } 
        if height[right] > maxRight {
            maxRight = height[right]
        }

        if maxLeft < maxRight {
            trapped += maxLeft - height[left]
            left++
        } else {
            trapped += maxRight - height[right]
            right--
        }
    }

    return trapped
}