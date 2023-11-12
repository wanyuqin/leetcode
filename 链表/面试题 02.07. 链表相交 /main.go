package main

func main() {
	//common := &ListNode{
	//	Val: 8,
	//	Next: &ListNode{
	//		Val: 4,
	//		Next: &ListNode{
	//			Val: 5,
	//		},
	//	},
	//}
	//headA := &ListNode{
	//	Val: 4,
	//	Next: &ListNode{
	//		Val:  1,
	//		Next: common,
	//	},
	//}
	//
	//headB := &ListNode{
	//	Val: 5,
	//	Next: &ListNode{
	//		Val: 0,
	//		Next: &ListNode{
	//			Val:  1,
	//			Next: common,
	//		},
	//	},
	//}

	headA := &ListNode{
		Val: 3,
	}
	headB := &ListNode{
		Val:  2,
		Next: headA,
	}
	println(getIntersectionNode(headA, headB).Val)

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

// 先计算出两个链表的长度
// 然后现将长的链表指针移动到差值长度的节点
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	lenA := 0
	lenB := 0
	dummyHeadA := &ListNode{
		Next: headA,
	}

	dummyHeadB := &ListNode{
		Next: headB,
	}
	curA := dummyHeadA
	curB := dummyHeadB

	for curA.Next != nil {
		lenA++
		curA = curA.Next
	}
	for curB.Next != nil {
		lenB++
		curB = curB.Next
	}

	if lenA > lenB {
		tmpA := dummyHeadA.Next
		skipA := lenA - lenB
		for skipA > 0 {
			skipA--
			tmpA = tmpA.Next
		}

		tmpB := dummyHeadB.Next
		for tmpB != nil {
			if tmpB == tmpA {
				return tmpA
			}
			tmpA = tmpA.Next
			tmpB = tmpB.Next

		}
	}

	if lenB >= lenA {
		tmpB := dummyHeadB.Next
		tmpA := dummyHeadA.Next
		skipB := lenB - lenA
		if skipB == 0 && tmpA == tmpB {
			return tmpA
		}
		for skipB > 0 {
			skipB--
			tmpB = tmpB.Next
		}

		for tmpA != nil {
			if tmpA == tmpB {
				return tmpA
			}
			tmpA = tmpA.Next
			tmpB = tmpB.Next

		}
	}
	return nil

}
