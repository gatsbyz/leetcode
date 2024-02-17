// furthestBuilding calculates the furthest building you can reach given a list of building heights,
// a certain amount of bricks, and a certain number of ladders.
func furthestBuilding(heights []int, bricks int, ladders int) int {
	q := hp{} // Create a new heap of type hp, which will store the height differences (jumps needed).
	n := len(heights) // Get the total number of buildings.
	
    // Iterate through each building, except the last one, since you can't jump from the last building.
	for i, a := range heights[:n-1] {
		b := heights[i+1] // Get the height of the next building.
		d := b - a // Calculate the difference in height (jump needed) to reach the next building.
		
        // If a jump is needed (the next building is higher),
		if d > 0 {
			heap.Push(&q, d) // Push the height difference onto the heap.
			
            // If the heap size exceeds the number of ladders available,
			if q.Len() > ladders {
				bricks -= heap.Pop(&q).(int) // Use bricks for the smallest jump (remove from heap)...
				
                // If bricks are depleted (negative count), you can't reach the next building.
				if bricks < 0 {
					return i // Return the index of the last building you can reach.
				}
			}
		}
	}
	return n - 1 // If you haven't returned by now, you can reach the last building.
}

// hp defines a custom heap type that embeds sort.IntSlice to leverage its methods,
// while also implementing heap's Push and Pop methods for int values.
type hp struct{ sort.IntSlice }

// Push adds an element (height difference that requires a jump) to the heap.
func (h *hp) Push(v any) { h.IntSlice = append(h.IntSlice, v.(int)) }

// Pop removes and returns the smallest element from the heap (the smallest height difference).
func (h *hp) Pop() any {
	a := h.IntSlice
	v := a[len(a)-1] // Get the last element (smallest, due to min-heap property).
	h.IntSlice = a[:len(a)-1] // Remove the last element from the slice.
	return v // Return the removed element.
}
