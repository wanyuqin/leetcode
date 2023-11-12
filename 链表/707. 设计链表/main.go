package main

type MyLinkedList struct {
	Val  int
	Next *LinkList
	Size int
}

type LinkList struct {
	Val  int
	Next *LinkList
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */

//
//["MyLinkedList","addAtHead","deleteAtIndex","addAtHead","addAtHead","addAtHead","addAtHead","addAtHead","addAtTail","get","deleteAtIndex","deleteAtIndex"]
//[[],[2],[1],[2],[7],[3],[2],[5],[5],[5],[6],[4]]

// ["MyLinkedList","addAtHead","deleteAtIndex"]
// [[],[1],[0]]

// ["MyLinkedList","addAtHead","addAtHead","addAtHead","addAtIndex","deleteAtIndex","addAtHead","addAtTail","get","addAtHead","addAtIndex","addAtHead"]
func main() {

	obj := Constructor()
	//
	//obj.AddAtHead(1)
	//obj.AddAtTail(3)
	//obj.AddAtIndex(1, 2)
	//println(obj.Get(1))
	//obj.DeleteAtIndex(1)
	//println(obj.Get(1))

	//fmt.Println(obj)
	// 1 2 7
	obj.AddAtHead(7)
	obj.AddAtHead(2)
	obj.AddAtHead(1)
	obj.AddAtIndex(3, 0)
	obj.DeleteAtIndex(2)
	obj.AddAtHead(6)
	obj.AddAtTail(4)
	println(obj.Get(4))
	obj.AddAtHead(4)
	obj.AddAtIndex(5, 0)
	obj.AddAtHead(6)
}

func Constructor() MyLinkedList {
	return MyLinkedList{
		Size: 0,
	}
}

func (this *MyLinkedList) Get(index int) int {
	res := -1
	if this.Size < index {
		return res
	}
	cur := this.Next
	for index >= 0 {
		res = cur.Val
		cur = cur.Next
		index--

	}
	return res
}

func (this *MyLinkedList) AddAtHead(val int) {
	if this.Size == 0 {
		this.Next = &LinkList{
			Val:  val,
			Next: nil,
		}
		this.Size++
		return
	}

	tmp := this.Next
	add := &LinkList{
		Val:  val,
		Next: tmp,
	}

	this.Next = add
	this.Size++

}

func (this *MyLinkedList) AddAtTail(val int) {
	cur := this.Next
	if this.Size == 0 {
		this.AddAtHead(val)
		return
	}
	for cur.Next != nil {
		cur = cur.Next
	}

	add := &LinkList{
		Val: val,
	}

	cur.Next = add
	this.Size++

}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if this.Size < index {
		return
	}
	if index == 0 {
		this.AddAtHead(val)
		return
	}

	if this.Size == index {
		this.AddAtTail(val)
		return
	}
	cur := this.Next
	for index > 1 {
		cur = cur.Next
		index--
	}
	add := &LinkList{
		Val:  val,
		Next: cur.Next,
	}
	cur.Next = add
	this.Size++

}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if this.Size < index {
		return
	}
	// 删除头节点
	if index == 0 {
		if this.Next.Next != nil {
			// 多个节点
			this.Next = this.Next.Next
			this.Size--
			return
		}

		if this.Next.Next == nil {
			// 只有一个节点
			this.Next = nil
			this.Size--
			return
		}

	}
	cur := this.Next
	for index > 1 {
		cur = cur.Next
		index--
	}
	if cur.Next != nil {
		cur.Next = cur.Next.Next
		this.Size--
		return
	}

	if cur.Next == nil {
		cur = nil
		this.Size--
		return
	}
}
