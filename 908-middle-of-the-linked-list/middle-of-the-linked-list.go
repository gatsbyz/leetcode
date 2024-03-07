/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNode(head *ListNode) *ListNode {
    pointer1 := head
    pointer2 := head

    for pointer2 != nil {
        pointer2 = pointer2.Next
        if pointer2 == nil {
            break
        }
        pointer2 = pointer2.Next
        pointer1 = pointer1.Next
    }

    return pointer1
}