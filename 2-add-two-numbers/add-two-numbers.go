/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    carry := 0
    sum := &ListNode{}
    curr := sum
    for l1 != nil && l2 != nil {
        curr.Next = &ListNode{}
        curr.Next.Val = (l1.Val + l2.Val + carry) % 10
        curr = curr.Next
        carry = (l1.Val + l2.Val + carry) / 10
        l1, l2 = l1.Next, l2.Next
    }
    for l1 != nil || l2 != nil {
            curr.Next = &ListNode{}
        if l1 != nil {
            curr.Next.Val = (l1.Val + carry) % 10
            carry = (l1.Val + carry) / 10
            l1 = l1.Next
        } else {
            curr.Next.Val = (l2.Val + carry) % 10
            carry = (l2.Val + carry) / 10
            l2 = l2.Next
        }
        curr = curr.Next
    }
    if carry > 0 {
        curr.Next = &ListNode{Val: carry }
    }
    return sum.Next
}