package sort_list

//在 O(n log n) 时间复杂度和常数级空间复杂度下，对链表进行排序。
//
//示例 1:
//
//输入: 4->2->1->3
//输出: 1->2->3->4
//示例 2:
//
//输入: -1->5->3->4->0
//输出: -1->0->3->4->5
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/sort-list
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var n int
	for h := head; h != nil; h = h.Next {
		n++
	}
	var tail = head
	var h1 *ListNode
	var h2 *ListNode
	var headCmp *ListNode
	var eleListLen = 1
	var last *ListNode
	for {
		if eleListLen >= n {
			break
		}
		headCmp = tail
		h1 = tail
		h2 = tail
		for i := 0; i < eleListLen; i++ {
			if h2.Next != nil {
				h2 = h2.Next
			}

			if i != eleListLen-1 {
				headCmp = headCmp.Next
			} else {
				headCmp.Next = nil
			}
		}
		tail = h2.Next
		h2.Next = nil

		headCmp, last = merge2Lists(h1, h2)
		headCmp.Next = tail

		eleListLen *= 2
	}

	return head
}

func merge2Lists(left *ListNode, right *ListNode) (head *ListNode, last *ListNode) {
	if left == nil && right == nil {
		return
	} else if left == nil {
		head = right
	} else if right == nil {
		head = left
	}
	last = head
	if last != nil {
		for {
			if last.Next == nil {
				break
			}
			last = last.Next
		}
		return
	}

	var next *ListNode
	for {
		if left.Val < right.Val {
			next = left
			left = left.Next
		} else {
			next = right
			right = right.Next
		}
		if last == nil {
			last = next
			head = last
		} else {
			last.Next = next
			last = next
		}
		if left == nil {
			last.Next = right
			break
		} else if right == nil {
			last.Next = left
			break
		}
	}

	for {
		if last.Next == nil {
			break
		}
		last = last.Next
	}
	return
}
