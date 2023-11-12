package main

func main() {
	last := &ListNode{
		Val: 4,
	}

	c := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 0,
		},
	}

	c.Next.Next = last
	last.Next = c

	head := &ListNode{
		Val:  3,
		Next: c,
	}

	detectCycle(head)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	if head.Next == head {
		return head
	}

	slow := head
	fast := head
	for fast != nil {
		slow = slow.Next
		if fast.Next == nil {
			return nil
		}
		fast = fast.Next.Next
		if fast == slow {
			p := head
			for p != slow {
				p = p.Next
				slow = slow.Next
			}
			return p
		}

	}

	return nil
}
