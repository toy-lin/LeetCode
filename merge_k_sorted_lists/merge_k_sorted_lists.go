package merge_k_sorted_lists

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

type HeapNode struct {
	N     *ListNode
	Left  *HeapNode
	Right *HeapNode
	LHigh int
	RHigh int
}

func (h *HeapNode) Push(n *ListNode) {
	if h == nil {
		return
	}
	if n.Val < h.N.Val{

	}
}

func (h *HeapNode) Pop() (n *ListNode) {

}

func mergeKLists(lists []*ListNode) *ListNode {
	if lists == nil {
		return nil
	}
	var minNode *ListNode
	var k int
	var sorted *ListNode
	var last *ListNode
	for {
		minNode = nil
		k = -1

		for i, node := range lists {
			if node != nil {
				if minNode == nil {
					minNode = node
					k = i
				} else if minNode.Val > node.Val {
					k = i
					minNode = node
				}
			}
		}
		if k == -1 {
			break
		}
		lists[k] = lists[k].Next
		minNode.Next = nil
		if sorted == nil {
			sorted = minNode
		} else {
			last.Next = minNode
		}
		last = minNode
	}
	return sorted
}
