type ZigzagIterator struct {
    v1 []int // turn even #
    v2 []int // turn odd #
    turn int
    index int
}

// ZigzagIterator(List<int> v1, List<int> v2) initializes the object with the two vectors v1 and v2.
func Constructor(v1, v2 []int) *ZigzagIterator {
    return &ZigzagIterator{
        v1: v1,
        v2: v2,
        index: 0,
        turn: 0,
    }
}

// int next() returns the current element of the iterator and moves the iterator to the next element.
func (zz *ZigzagIterator) next() int {
    var element int
    if (zz.turn % 2  == 0 || len(zz.v2) <= zz.index) && zz.index < len(zz.v1) {
        element = zz.v1[zz.index]
    } else {
        element = zz.v2[zz.index]
    } 

    if len(zz.v1) <= zz.index || len(zz.v2) <= zz.index || zz.turn % 2 == 1 {
        zz.index++
    }
    zz.turn++

    return element
}

// boolean hasNext() returns true if the iterator still has elements, and false otherwise.
func (zz *ZigzagIterator) hasNext() bool {
	return zz.index < len(zz.v1) || zz.index < len(zz.v2)
}

/**
 * Your ZigzagIterator object will be instantiated and called as such:
 * obj := Constructor(param_1, param_2);
 * for obj.hasNext() {
 *	 ans = append(ans, obj.next())
 * }
 */