package main

import (
	"fmt"
)

type Node struct {
	v    int
	next *Node
}

func Insert(head **Node, n int) {
	if *head == nil {
		*head = &Node{v: n}
	} else {
		tmp := *head
		for tmp.next != nil {
			tmp = tmp.next
		}
		tmp.next = &Node{v: n}
	}
}

func Print(head *Node) {
	tmp := head
	for tmp != nil {
		fmt.Println(tmp.v)
		tmp = tmp.next
	}
}

func PartialReverse1(head *Node, p, q int) {
	var ppos *Node
	var qpos *Node
	var a []int
	var c int
	var i int
	tmp := head
	for tmp != nil {
		c++
		if c == p {
			a = make([]int, q-p+1)
			ppos = tmp

		}
		if a != nil {
			a[i] = tmp.v
			i++
		}
		if c == q {
			qpos = tmp
			break
		}
		tmp = tmp.next
	}
	for ppos != qpos.next {
		ppos.v = a[i-1]
		i--
		ppos = ppos.next
	}
}
func PartialReverse2(head *Node, p, q int) {
	var ppos *Node
	var pprev *Node
	var qpos *Node
	var qnext *Node
	var c int
	tmp := head
	for tmp != nil {
		c++
		if c == p {
			ppos = tmp

		}
		if c == q {
			qpos = tmp
			break
		}
		if ppos == nil {
			pprev = tmp
		}
		tmp = tmp.next
	}
	if pprev != nil {
		pprev.next = nil
	}
	if qpos != nil {
		qnext = qpos.next
		qpos.next = nil
	}
	for ppos != nil {
		tmp = ppos
		ppos = ppos.next
		tmp.next = qnext
		qnext = tmp
	}
	if pprev != nil {
		pprev.next = qnext
	}
}

func Reverse(h *Node) *Node {
	var tmp *Node
	var prev *Node
	for h != nil {
		tmp = h
		h = h.next
		tmp.next = prev
		prev = tmp
	}
	return prev
}

func ReverseRecursive1(a, b *Node) *Node {
	if a == nil {
		return b
	} else {
		tmp := a.next
		a.next = b
		return ReverseRecursive1(tmp, a)
	}
}

func ReverseRecursive2(h *Node) *Node {
	if h == nil || h.next == nil {
		return h
	}
	rest := ReverseRecursive2(h.next)
	h.next.next = h
	h.next = nil
	return rest
}

func MakeCycle(h *Node, pos int) {
	tmp := h
	for pos > 1 {
		pos--
		tmp = tmp.next
	}
	k := tmp
	for tmp != nil && tmp.next != nil {
		tmp = tmp.next
	}
	tmp.next = k
}

func DetectCycle(h *Node) *Node {
	if h == nil {
		return h
	}
	fast := h.next
	slow := h
	for slow != nil && fast != nil {
		if slow == fast {
			return fast
		}
		slow = slow.next
		fast = fast.next
		if fast != nil {
			fast = fast.next
		}
	}
	return nil
}

func FindCycleNode(h *Node) *Node {
	b := DetectCycle(h)
	if b == nil {
		return b
	}
	bnext := b.next
	length := 1
	for bnext != b {
		bnext = bnext.next
		length++
	}
	tmp := h
	for length > 0 {
		tmp = tmp.next
		length--
	}
	bnext = b.next
	for tmp != bnext {
		tmp = tmp.next
		bnext = bnext.next
	}
	return tmp
}

func MakeX(h, node *Node) {
	tmp := h
	for tmp.next != nil {
		tmp = tmp.next
	}
	tmp.next = node
}

func FindNode(h *Node, n int) *Node {
	tmp := h
	for tmp != nil {
		if tmp.v == n {
			return tmp
		}
		tmp = tmp.next
	}
	return nil
}

func FindLength(h *Node) int {
	if h == nil {
		return 0
	}
	tmp := h
	n := 0
	for tmp != nil {
		n++
		tmp = tmp.next
	}
	return n
}

func FindX(x, y *Node) *Node {
	if x == nil || y == nil {
		return nil
	}
	xlen := FindLength(x)
	ylen := FindLength(y)
	d := xlen - ylen
	xtmp := x
	ytmp := y
	if d > 0 {
		for d > 0 {
			xtmp = xtmp.next
			d--
		}
	} else if d < 0 {
		for d < 0 {
			d++
			ytmp = ytmp.next
		}
	}
	for xtmp != nil && ytmp != nil {
		if xtmp == ytmp {
			return xtmp
		}
		xtmp = xtmp.next
		ytmp = ytmp.next
	}
	return nil
}

// rotate from last to pos
func rotate(h *Node, pos int) {
	/*tmp := h
	var prev *Node
	n := 1
	for tmp.next != nil {
		tmp = tmp.next
		n++
	}
	last = tmp
	tmp := h
	count := n - (pos - 1)
	for count > 0 {
		count--
		tmp = tmp.next
	}
	tmp.next = nil
	last.next = tmp
	h = tmp*/
}
func mergeSortedLists(list1, list2 *Node) *Node {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	var head *Node
	var tmp *Node
	headtmp := head
	for list1 != nil && list2 != nil {
		if list1.v < list2.v {
			tmp = list1
			list1 = list1.next
			headtmp = tmp
			headtmp = headtmp.next
		} else if list1.v >= list2.v {
			tmp = list2
			list2 = list2.next
			headtmp = tmp
			headtmp = headtmp.next
		}
		for list1 != nil {
			tmp = list1
			list1 = list1.next
			headtmp = tmp
			headtmp = headtmp.next

		}
		for list2 != nil {
			tmp = list2
			list2 = list2.next
			headtmp = tmp
			headtmp = headtmp.next
		}
	}
	return head
}

func mergeKLists(lists []*Node) *Node {
	n := len(lists)
	if n == 0 {
		return nil
	}
	if n == 1 {
		return lists[0]
	}
	var head *Node
	for i := 0; i < n; i++ {
		head = mergeSortedLists(head, lists[i])
	}

	return head
}

func main() {
	var h *Node
	Insert(&h, 1)
	Insert(&h, 2)
	Insert(&h, 3)
	Insert(&h, 4)
	Insert(&h, 5)
	Insert(&h, 6)
	Insert(&h, 7)
	Insert(&h, 8)
	Insert(&h, 9)
	Insert(&h, 10)
	PartialReverse1(h, 6, 9)
	//Print(h)
	PartialReverse1(h, 6, 9)
	h = Reverse(h)
	var b *Node
	//h = ReverseRecursive1(h, b)
	h = ReverseRecursive2(h)
	//Print(h)
	//	MakeCycle(h, 7)
	b = FindCycleNode(h)
	if b != nil {
		fmt.Println(b.v)
	}
	var t *Node
	//Insert(&t, 1)
	//Insert(&t, 2)
	//Insert(&t, 3)
	Insert(&t, 4)
	Insert(&t, 5)
	b = FindNode(h, 6)
	MakeX(t, b)
	fmt.Println("==============================")
	Print(t)
	x := FindX(h, t)
	if x != nil {
		fmt.Println(x.v)
	}
	var l1 *Node
	var l2 *Node
	var l3 *Node
	Insert(&l1, 1)
	Insert(&l1, 4)
	Insert(&l1, 5)
	Insert(&l2, 1)
	Insert(&l2, 3)
	Insert(&l2, 4)
	Insert(&l3, 2)
	Insert(&l3, 6)
	lists := []*Node{l1, l2, l3}
	newList := mergeKLists(lists)
	Print(newList)

}
