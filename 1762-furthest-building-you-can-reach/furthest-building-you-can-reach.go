// IntHeap defines a type that implements heap.Interface for min-heap operations on ints.
type IntHeap []int

// Len returns the number of elements in the heap.
func (h IntHeap) Len() int { return len(h) }

// Less reports whether the element with index i should sort before the element with index j.
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }

// Swap swaps the elements with indexes i and j.
func (h IntHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

// Push adds a new element to the heap.
func (h *IntHeap) Push(x interface{}) {
    *h = append(*h, x.(int))
}

// Pop removes and returns the smallest element from the heap.
func (h *IntHeap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[:n-1]
    return x
}

// furthestBuilding calculates the furthest building you can reach with given bricks and ladders.
func furthestBuilding(heights []int, bricks int, ladders int) int {
    h := &IntHeap{}
    heap.Init(h) // Initialize the heap to keep track of the heights where ladders are used.

    // Iterate through each building to determine if you can move to the next one.
    for i := 0; i < len(heights)-1; i++ {
        diff := heights[i+1] - heights[i] // Calculate the height difference to the next building.

        // If the next building is higher, consider using bricks or a ladder.
        if diff > 0 {
            heap.Push(h, diff) // Pretend to use a ladder by pushing the height difference onto the heap.
            
            // If we've used more ladders than available, replace one ladder usage with bricks.
            if h.Len() > ladders {
                bricks -= heap.Pop(h).(int) // Use bricks for the smallest height difference encountered so far.
            }
            
            // If we don't have enough bricks to cover the next jump, we cannot move further.
            if bricks < 0 {
                return i
            }
        }
    }
    
    // If we can traverse all buildings without running out of resources, return the last building index.
    return len(heights) - 1
}
