/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
    var newNode *ListNode
    for head != nil {
        temp := head.Next
        head.Next = newNode
        newNode = head
        head = temp
    }
    return newNode
}