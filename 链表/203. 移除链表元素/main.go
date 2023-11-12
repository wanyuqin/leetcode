package main

import "fmt"

func main() {
	// 1,2,6,3,4,5,6
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 1,
						Next: &ListNode{
							Val: 1,
							Next: &ListNode{
								Val:  1,
								Next: nil,
							},
						},
					},
				},
			},
		},
	}
	res := removeElements(head, 1)
	fmt.Println(res)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// 1,2,6,3,4,5,6
func removeElements(head *ListNode, val int) *ListNode {
	v := &ListNode{} // 虚拟节点
	v.Next = head
	p := v
	for p != nil && p.Next != nil {
		if p.Next.Val == val {
			p.Next = p.Next.Next
		} else {
			// 节点值不想等的时候 才会进行下一位
			p = p.Next
		}

	}

	return v.Next
}
