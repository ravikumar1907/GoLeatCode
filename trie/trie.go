package main

import (
	"fmt"
)

const AlphabetSize = 26

type Node struct {
	children [AlphabetSize]*Node
	IsEnd    bool
}

type Trie struct {
	root *Node
}

func InitTrie() *Trie {
	return &Trie{root: &Node{}}
}

func (T *Trie) Insert(w string) {
	currentNode := T.root
	wlen := len(w)
	for i := 0; i < wlen; i++ {
		idx := w[i] - 'a'
		if currentNode.children[idx] == nil {
			currentNode.children[idx] = &Node{}
		}
		currentNode = currentNode.children[idx]
	}
	currentNode.IsEnd = true
}

func (T *Trie) Delete(w string) int {
	currentNode := T.root
	wlen := len(w)
	var nodes []*Node
	for i := 0; i < wlen; i++ {
		idx := w[i] - 'a'
		if currentNode.children[idx] == nil {
			return -1
		}
		
		nodes = append(nodes, currentNode)
		currentNode = currentNode.children[idx]
	}
	currentNode.IsEnd = false
	if len(currentNode.children) <= 0 {
		for _, n := range nodes {
			if !n.IsEnd &&  len(n.children) <= 1 {
				n.children[] = nil
			}
	}
}
	return 0
}

func (T *Trie) Search(w string) bool {
	currentNode := T.root
	wlen := len(w)
	for i := 0; i < wlen; i++ {
		idx := w[i] - 'a'
		if currentNode.children[idx] == nil {
			return false
		}
		currentNode = currentNode.children[idx]
	}
	if currentNode.IsEnd == true {
		return true
	}
	return false
}

func main() {
	myTrie := InitTrie()
	fmt.Println(myTrie)
	myTrie.Insert("ravikumar")
	myTrie.Insert("ravi")
	fmt.Println(myTrie.Search("ravikumar"))
	fmt.Println(myTrie.Search("ravi"))
	fmt.Println(myTrie.Search("ravik"))
}
