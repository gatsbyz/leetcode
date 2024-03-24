/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
    if head.Next == nil {
        return true
    }
    first, second := head, head
    stack := []*ListNode{}
    for second != nil {
        stack = append(stack, first)
        first = first.Next
        second = second.Next
        if second != nil {
            second = second.Next
        } else {
            stack = stack[:len(stack)-1]
            break
        }
    }
    for first != nil {
        if stack[len(stack)-1].Val != first.Val {
            return false
        }
        stack = stack[:len(stack)-1]
        first = first.Next
    }
    if len(stack) > 0 {
        return false
    }
    return true
}