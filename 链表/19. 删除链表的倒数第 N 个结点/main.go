package main

func main() {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
					},
				},
			}},
	}
	removeNthFromEnd(head, 2)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 先让fast移动n
// 再开始让slow、fast同时移动
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummyHead := &ListNode{
		Next: head,
	}

	fast := dummyHead.Next
	slow := dummyHead
	for fast.Next != nil {
		if n > 0 {
			n--
		}
		if n == 0 {
			slow = slow.Next
		}
		fast = fast.Next
	}
	slow.Next = slow.Next.Next
	return dummyHead.Next

}
