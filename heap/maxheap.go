package main

import (
	"container/heap"
	"fmt"
)

type IntHeap []int
type Max_Heap struct{ IntHeap }
type Min_Heap struct{ IntHeap }

func (h IntHeap) Len() int            { return len(h) }
func (h IntHeap) Empty() bool         { return len(h) == 0 }
func (h IntHeap) Top() int            { return h[0] }
func (h IntHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *IntHeap) Pop() interface{} {
	x := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return x
}

func (h Max_Heap) Less(i, j int) bool { return h.IntHeap[i] > h.IntHeap[j] }
func (h Min_Heap) Less(i, j int) bool { return h.IntHeap[i] < h.IntHeap[j] }

type MedianFinder1 struct {
	smallList *Max_Heap
	largeList *Min_Heap
}

func New() *MedianFinder1 {
	max := &Max_Heap{}
	min := &Min_Heap{}
	heap.Init(max)
	heap.Init(min)
	return &MedianFinder1{smallList: max, largeList: min}
}

func (this *MedianFinder1) InsertNum(x int) {
	if this.smallList.Empty() || this.smallList.Top() > x {
		heap.Push(this.smallList, x)
	} else {
		heap.Push(this.largeList, x)
	}
	if this.smallList.Len() > this.largeList.Len()+1 {
		heap.Push(this.largeList, heap.Pop(this.smallList).(int))
	} else if this.smallList.Len() < this.largeList.Len() {
		heap.Push(this.smallList, heap.Pop(this.largeList).(int))
	}
}

func (this *MedianFinder1) FindMedian() float64 {
	if this.smallList.Len() == this.largeList.Len() {
		return float64(this.smallList.Top()+this.largeList.Top()) / 2.0
	}
	return float64(this.smallList.Top())
}

func main() {
	/*maxHeap := &Max_Heap{}
	heap.Init(maxHeap)
	heap.Push(maxHeap, 5)
	heap.Push(maxHeap, 8)
	heap.Push(maxHeap, 7)
	heap.Push(maxHeap, 9)
	fmt.Printf("Top:%d, Heap:%v\n", maxHeap.Top(), maxHeap)
	n := maxHeap.Len()
	for n > 0 {
		fmt.Print(heap.Pop(maxHeap).(int))
		fmt.Print(" ")
		n--
	}*/
	medianOfAges := new()
	medianOfAges.InsertNum(22)
	medianOfAges.InsertNum(35)
	fmt.Printf("The recommended content will be for ages under: %f\n", medianOfAges.FindMedian())
	medianOfAges.InsertNum(30)
	fmt.Printf("The recommended content will be for ages under: %f\n", medianOfAges.FindMedian())
	medianOfAges.InsertNum(25)
	//fmt.Printf("Small Top:%d, Small:%v, Large Top:%d, Large::%v\n", medianOfAges.smallList.Top(), medianOfAges.smallList, medianOfAges.largeList.Top(), medianOfAges.largeList)
	fmt.Printf("The recommended content will be for ages under: %f\n", medianOfAges.FindMedian())
	mf := New()
	mf.InsertNum(22)
	mf.InsertNum(35)
	fmt.Printf("The recommended content will be for ages under: %f\n", mf.FindMedian())
	mf.InsertNum(30)
	fmt.Printf("The recommended content will be for ages under: %f\n", mf.FindMedian())
	mf.InsertNum(25)
	//fmt.Printf("Small Top:%d, Small:%v, Large Top:%d, Large::%v\n", medianOfAges.smallList.Top(), medianOfAges.smallList, medianOfAges.largeList.Top(), medianOfAges.largeList)
	fmt.Printf("The recommended content will be for ages under: %f\n", mf.FindMedian())
}
