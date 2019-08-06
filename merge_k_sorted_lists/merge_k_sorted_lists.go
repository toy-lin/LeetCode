package merge_k_sorted_lists

import "container/heap"

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) < 1 {
		return nil
	}

	var h nodeHeap
	for i := 0; i < len(lists); i++ {
		if lists[i] != nil {
			h = append(h, lists[i])
		}
	}
	if len(h) < 1 {
		return nil
	}
	heap.Init(h)
	var last = h[0]
	var sorted = last
	for {
		h[0] = h[0].Next
		if h[0] == nil {
			h = heap.Pop(h).(nodeHeap)
		}
		if len(h) < 1 {
			break
		}
		if len(h) == 1 {
			last.Next = h[0]
			break
		}
		heap.Fix(h, 0)

		last.Next = h[0]
		last = last.Next
	}
	return sorted
}

type nodeHeap []*ListNode

func (h nodeHeap) Len() int {
	return len(h)
}

func (h nodeHeap) Less(i, j int) bool {
	return h[i].Val < h[j].Val
}

func (h nodeHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (nodeHeap) Push(x interface{}) {}
func (h nodeHeap) Pop() interface{} { return h[:len(h)-1] }
