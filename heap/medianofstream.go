package main

import (
	"container/heap"
)

type BaseHeap []int
type MaxHeap struct {
	BaseHeap
}
type MinHeap struct {
	BaseHeap
}

type MinFinder struct {
	smallList *MaxHeap
	largeList *MinHeap
}

func (h BaseHeap) Len() int            { return len(h) }
func (h BaseHeap) Empty() bool         { return len(h) == 0 }
func (h BaseHeap) Top() int            { return h[0] }
func (h BaseHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *BaseHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *BaseHeap) Pop() interface{} {
	n := len(*h)
	x := (*h)[n-1]
	(*h) = (*h)[:n-1]
	return x
}

func (h MaxHeap) Less(i, j int) bool { return h.BaseHeap[i] > h.BaseHeap[j] }
func (h MinHeap) Less(i, j int) bool { return h.BaseHeap[i] < h.BaseHeap[j] }

type MedianFinder struct {
	smallList *MaxHeap
	largeList *MinHeap
}

func new() MedianFinder {
	max := &MaxHeap{}
	min := &MinHeap{}
	heap.Init(max)
	heap.Init(min)
	return MedianFinder{smallList: max, largeList: min}
}

func (s MedianFinder) InsertNum(x int) {
	if s.smallList.Empty() || s.smallList.Top() > x {
		heap.Push(s.smallList, x)
	} else {
		heap.Push(s.largeList, x)
	}
	if s.smallList.Len() > s.largeList.Len()+1 {
		heap.Push(s.largeList, heap.Pop(s.smallList).(int))
	} else if s.smallList.Len() < s.largeList.Len() {
		heap.Push(s.smallList, heap.Pop(s.largeList).(int))
	}

}

func (s MedianFinder) FindMedian() float64 {
	if s.smallList.Len() == s.largeList.Len() {
		return float64(s.smallList.Top()+s.largeList.Top()) / 2.0
	}
	return float64(s.smallList.Top())
}

/*
func main() {

	medianOfAges := new()
	medianOfAges.InsertNum(22)
	medianOfAges.InsertNum(35)
	fmt.Printf("The recommended content will be for ages under: %f\n", medianOfAges.FindMedian())
	medianOfAges.InsertNum(30)
	fmt.Printf("The recommended content will be for ages under: %f\n", medianOfAges.FindMedian())
	medianOfAges.InsertNum(25)
	fmt.Printf("Small Top:%d, Small:%v, Large Top:%d, Large::%v\n", medianOfAges.smallList.Top(), medianOfAges.smallList, medianOfAges.largeList.Top(), medianOfAges.largeList)
	fmt.Printf("The recommended content will be for ages under: %f\n", medianOfAges.FindMedian())
}
*/
