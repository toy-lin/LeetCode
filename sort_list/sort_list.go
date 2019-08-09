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
	var h1 *ListNode
	var h2 *ListNode
	var blockLen = 1
	var cmpPre *ListNode
	var cmpPost *ListNode
	var cutAssist *ListNode
	for {
		// 1
		if blockLen >= n {
			break
		}
		// 2
		if cmpPre == nil {
			h1 = head
		} else {
			h1 = cmpPre.Next
		}
		h2 = h1
		cutAssist = h1
		for i := 0; i < blockLen; i++ {
			if h2 != nil {
				h2 = h2.Next

				if i == blockLen-1 {
					cutAssist.Next = nil
				} else {
					cutAssist = cutAssist.Next
				}
			}
		}
		cmpPost = h2
		cutAssist = h2
		for i := 0; i < blockLen; i++ {
			if cmpPost != nil {
				cmpPost = cmpPost.Next
				if i == blockLen-1 {
					cutAssist.Next = nil
				} else {
					cutAssist = cutAssist.Next
				}
			}
		}
		// 4
		h1, h2 = merge2Lists(h1, h2)
		// 5
		if cmpPre == nil {
			head = h1
		} else {
			cmpPre.Next = h1
		}
		cmpPre = h2
		h2.Next = cmpPost

		if cmpPost == nil {
			blockLen *= 2
			cmpPre = nil
		}
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
