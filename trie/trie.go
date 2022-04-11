package main

import "fmt"

const AlphabetSize = 26

type trieNode struct {
	children [AlphabetSize]*trieNode
	IsEnd    bool
	IsDir    bool // for in memory file system
}

type WordDictionary struct {
	root *trieNode
}

func Constructor() *WordDictionary {
	return &WordDictionary{root: &trieNode{}}
}

func (T *WordDictionary) AddWord(w string) {
	currentNode := T.root
	wlen := len(w)
	for i := 0; i < wlen; i++ {
		if w[i] == '/' {
			currentNode.IsDir = true
			continue
		}
		idx := w[i] - 'a'
		if currentNode.children[idx] == nil {
			currentNode.children[idx] = &trieNode{}
		}
		currentNode = currentNode.children[idx]
	}
	currentNode.IsEnd = true
}

func (T *WordDictionary) Delete(w string) int {
	currentNode := T.root
	wlen := len(w)
	var nodes []*trieNode
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
			if !n.IsEnd && len(n.children) <= 1 {
				n.children = [AlphabetSize]*trieNode{}
			}
		}
	}
	return 0
}

func (T *WordDictionary) Search(w string) bool {
	var dfs func(k int, root *trieNode) bool
	dfs = func(k int, root *trieNode) bool {
		currentNode := root
		wlen := len(w)
		for i := k; i < wlen; i++ {
			if w[i] == '.' {
				for _, child := range currentNode.children {
					if child != nil && dfs(i+1, child) {
						return true
					}
				}
				return false
			} else {
				idx := w[i] - 'a'
				if currentNode.children[idx] == nil {
					return false
				}
				currentNode = currentNode.children[idx]
			}
		}
		return currentNode.IsEnd
	}
	return dfs(0, T.root)
}

func main() {
	myTrie := Constructor()
	/*fmt.Println(myTrie)
	myTrie.AddWord("ravikumar")
	myTrie.AddWord("ravi")
	fmt.Println(myTrie.Search("ravikumar"))
	fmt.Println(myTrie.Search("ravi"))
	fmt.Println(myTrie.Search("ravik"))*/
	myTrie.AddWord("bad")
	myTrie.AddWord("dad")
	myTrie.AddWord("mad")
	fmt.Println(myTrie.Search("pad"))
	fmt.Println(myTrie.Search("bad"))
	fmt.Println(myTrie.Search(".ad"))
	fmt.Println(myTrie.Search("b.."))
}
