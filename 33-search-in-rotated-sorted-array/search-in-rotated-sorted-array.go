func search(nums []int, target int) int {
      left := 0
  right := len(nums)-1

  n := -1
  for left <= right {
    mid_index := (left + right)/2
    mid := nums[mid_index]

    if mid == target {
      n = mid_index
      break
    }

// 6,7,1,2,4 
// target = 7


    if target < mid {
      if target < nums[left] {
        if mid < nums[right] {
         // look left
           right = mid_index - 1
        } else {
          // look right
          left = mid_index + 1
        }
      } else {
         // look left
           right = mid_index - 1
      }
    } else {
        if target > nums[right] {
            if nums[left] < mid {
              left = mid_index + 1
            } else {
                // look left
                right = mid_index - 1
            }
      } else {
        // look right
          left = mid_index + 1
      }
    }
  }

	return n

}