package main

import (
	"container/heap"
	"sort"
)

/*
1094. Car Pooling
There is a car with capacity empty seats. The vehicle only drives east
 (i.e., it cannot turn around and drive west).
You are given the integer capacity and an array trips
where trips[i] = [numPassengersi, fromi, toi] indicates that the ith trip has numPassengersi
passengers and the locations to pick them up and drop them off are fromi and toi respectively.
The locations are given as the number of kilometers due east from the car's initial location.
Return true if it is possible to pick up and drop off all passengers for all the given trips,
or false otherwise.
*/

type ivl struct {
	end     int
	numPass int
}
type CarPoolMinHeap []*ivl

func (h CarPoolMinHeap) Len() int      { return len(h) }
func (h CarPoolMinHeap) Empty() bool   { return (len(h) == 0) }
func (h CarPoolMinHeap) Top() *ivl     { return h[0] }
func (h CarPoolMinHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *CarPoolMinHeap) Push(i interface{}) {
	(*h) = append((*h), i.(*ivl))
}
func (h *CarPoolMinHeap) Pop() interface{} {
	i := (*h)[len(*h)-1]
	(*h) = (*h)[:len(*h)-1]
	return i
}

func (h CarPoolMinHeap) Less(i, j int) bool {
	return h[i].end < h[j].end
}

// O(nlongn)
func carPooling(trips [][]int, capacity int) bool {
	sort.Slice(trips, func(i, j int) bool {
		return trips[i][1] < trips[j][1]
	})
	var minHeap CarPoolMinHeap
	curPass := 0
	heap.Init(&minHeap)
	for _, t := range trips {
		numPass, start, end := t[0], t[1], t[2]
		for !minHeap.Empty() && minHeap.Top().end <= start {
			curPass -= minHeap.Top().numPass
			heap.Pop(&minHeap)
		}
		curPass += numPass
		if curPass > capacity {
			return false
		}
		heap.Push(&minHeap, &ivl{end, numPass})
	}
	return true
}

/*
1 <= trips.length <= 1000
trips[i].length == 3
1 <= numPassengersi <= 100
0 <= fromi < toi <= 1000
1 <= capacity <= 105
*/
// Considering above constraints 0 <= fromi < toi <= 1000
// we can have below brute force method.
// O(n)
func carPooling2(trips [][]int, capacity int) bool {
	passChange := make([]int, 1001)
	for _, v := range trips {
		numPass, start, end := v[0], v[1], v[2]
		passChange[start] += numPass
		passChange[end] -= numPass
	}
	curPass := 0
	for i := 0; i < 1001; i++ {
		curPass += passChange[i]
		if curPass > capacity {
			return false
		}
	}
	return true
}
