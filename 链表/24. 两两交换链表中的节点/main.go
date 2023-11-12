package main

func main() {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{Val: 2, Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val: 4,
				Next: &ListNode{
					Val: 5,
				},
			},
		}},
	}
	swapPairs(head)
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
// 需要注意定义一个虚拟头节点
// 返回的是虚拟头节点
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dummyNode := &ListNode{
		Next: head,
	}
	cur := dummyNode

	for cur.Next != nil && cur.Next.Next != nil {
		tmp := cur.Next            // 节点一
		tmp1 := cur.Next.Next.Next // 节点三
		cur.Next = cur.Next.Next   // 节点二

		cur.Next.Next = tmp // 节点二指向节点一
		tmp.Next = tmp1     // 节点一指向节点三
		cur = cur.Next.Next // 移动cur
	}
	return dummyNode.Next

	//dummy := &ListNode{
	//	Next: head,
	//}
	////head=list[i]
	////pre=list[i-1]
	//pre := dummy
	//for head != nil && head.Next != nil {
	//	pre.Next = head.Next
	//	next := head.Next.Next
	//	head.Next.Next = head
	//	head.Next = next
	//	//pre=list[(i+2)-1]
	//	pre = head
	//	//head=list[(i+2)]
	//	head = next
	//}
	//return dummy.Next
}
