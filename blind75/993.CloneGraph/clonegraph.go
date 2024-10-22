package main

// Definition for a Node.
type Node struct {
	Val       int
	Neighbors []*Node

	// Below two for linked list problem with random pointer
	Next      *Node
	Random    *Node
}

func cloneGraphBFS(node *Node) *Node {
	if node == nil {
		return nil
	}
	queue := []*Node{node}
	clonedNodes := make(map[*Node]*Node)
	clonedNodes[node] = &Node{Val: node.Val}
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		for _, neighbor := range current.Neighbors {
			if _, exists := clonedNodes[neighbor]; !exists {
				queue = append(queue, neighbor)
				clonedNodes[neighbor] = &Node{Val: neighbor.Val}
			}
			clonedNodes[current].Neighbors = append(clonedNodes[current].Neighbors, clonedNodes[neighbor])
		}

	}
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return head
	}
	current := head
	for current != nil {
		clone := &Node{Val: current.Val}
		clone.Next = current.Next
		current.Next = clone
		current = clone.Next
	}
	current = head
	for current != nil {
		if current.Random != nil {
			current.Next.Random = current.Random.Next
		}
		current = current.Next.Next
	}
	current = head
	cloneHead := head.Next
	cloneCurrent := cloneHead
	for current != nil {
		current.Next = current.Next.Next
		if cloneCurrent.Next != nil {
			cloneCurrent.Next = cloneCurrent.Next.Next
		}
		current = current.Next
		cloneCurrent.Next = cloneCurrent.Next
	}
	return cloneHead
}
