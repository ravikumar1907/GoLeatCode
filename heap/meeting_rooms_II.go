package main

import (
	"container/heap"
	"fmt"
	"sort"
)

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

type intvl struct {
	Start int
	End   int
}

type Intervals []*intvl

type MeetRoom1 struct {
	Intervals
}
type MeetRoom2 struct {
	Intervals
}

func (h Intervals) Len() int      { return len(h) }
func (h Intervals) Top() *intvl   { return h[0] }
func (h Intervals) Empty() bool   { return (len(h) == 0) }
func (h Intervals) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *Intervals) Push(i interface{}) {
	(*h) = append(*h, i.(*intvl))
}
func (h *Intervals) Pop() interface{} {
	i := (*h)[len(*h)-1]
	(*h) = (*h)[:len(*h)-1]
	return i
}

func (h MeetRoom1) Less(i, j int) bool {
	return h.Intervals[i].Start < h.Intervals[j].Start
}

func (h MeetRoom2) Less(i, j int) bool {
	return h.Intervals[i].End < h.Intervals[j].End
}

// Two pointer solution.
//253. Meeting Rooms II
/*
	Given an array of meeting time intervals intervals where intervals[i] = [starti, endi],
	return the minimum number of conference rooms required.
*/

func minMeetingRooms(intervals [][]int) int {
	start := make([]int, len(intervals))
	end := make([]int, len(intervals))
	for i, v := range intervals {
		start[i] = v[0]
		end[i] = v[1]
	}
	sort.Ints(start)
	sort.Ints(end)
	s, e := 0, 0
	res, count := 0, 0
	for s < len(intervals) {
		if start[s] < end[e] {
			s += 1
			count += 1
		} else {
			e += 1
			count -= 1
		}
		res = max(res, count)
	}
	return res
}

func minMeetingRooms2(intervals [][]int) int {
	var minHeap MeetRoom2
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	heap.Init(&minHeap)
	for i := 0; i < len(intervals); i++ {
		if !minHeap.Empty() {
			if intervals[i][0] >= minHeap.Top().End {
				heap.Pop(&minHeap)
			}
		}
		heap.Push(&minHeap, &intvl{intervals[i][0], intervals[i][1]})
	}
	return minHeap.Len()
}

//252. Meeting Rooms
/*
	Given an array of meeting time intervals where intervals[i] = [starti, endi],
	determine if a person could attend all meetings.
*/
func canAttendMeetings(intervals [][]int) bool {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	for i := 1; i < len(intervals); i++ {
		s1 := intervals[i-1][0]
		s2 := intervals[i][0]
		e1 := intervals[i-1][1]
		e2 := intervals[i][1]
		if s1 < s2 && e2 < e1 {
			return false
		}
		if s2 < e1 {
			return false
		}
	}
	return true
}

func canAttendMeetings2(intervals [][]int) bool {
	var min_heap MeetRoom1
	heap.Init(&min_heap)
	for i := 0; i < len(intervals); i++ {
		heap.Push(&min_heap, &intvl{intervals[i][0], intervals[i][1]})
	}
	var top *intvl
	if !min_heap.Empty() {
		top = heap.Pop(&min_heap).(*intvl)
	}
	for !min_heap.Empty() {
		next := heap.Pop(&min_heap).(*intvl)
		if top.Start < next.Start && next.End < top.End {
			return false
		}
		if next.Start < top.End {
			return false
		}
		top = next
	}
	return true
}

func main() {
	in := [][]int{{0, 30}, {60, 240}, {90, 120}}
	fmt.Println(canAttendMeetings2(in))
	in = [][]int{{0, 30}, {5, 10}, {15, 20}}
	in = [][]int{{1, 8}, {4, 6}}
	fmt.Println(minMeetingRooms2(in))
}
