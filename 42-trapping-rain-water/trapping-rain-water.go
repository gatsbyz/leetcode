func trap(height []int) int {
	left, right := 0, len(height)-1 // Two pointers
	leftMax, rightMax := 0, 0       // Max heights from both sides
	trapped := 0                    // Total trapped water

	// Iterate until the two pointers meet
	for left < right {
		// Update leftMax and rightMax
		if height[left] > leftMax {
			leftMax = height[left]
		}
		if height[right] > rightMax {
			rightMax = height[right]
		}

		// Calculate trapped water based on the smaller of the two max heights
		if leftMax < rightMax {
			// Water trapped on the current left position is the difference
			// between leftMax and the current height, if any
			trapped += leftMax - height[left]
			left++ // Move left pointer inwards
		} else {
			// Similarly, water trapped on the current right position
			trapped += rightMax - height[right]
			right-- // Move right pointer inwards
		}
	}

	return trapped
}