package merge_two_list

//将两个有序链表合并为一个新的有序链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
//
//示例：
//
//输入：1->2->4, 1->3->4
//输出：1->1->2->3->4->4
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/merge-two-sorted-lists
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) (result *ListNode) {
	var tail *ListNode
	for {
		if l1 == nil && l2 == nil {
			return
		}
		if result == nil {
			result = &ListNode{}
			tail = result
		} else {
			tail.Next = &ListNode{}
			tail = tail.Next
		}
		if l1 != nil && l2 != nil {
			if l1.Val > l2.Val {
				tail.Val = l2.Val
				l2 = l2.Next
			} else {
				tail.Val = l1.Val
				l1 = l1.Next
			}
			continue
		}
		if l1 != nil {
			tail.Val = l1.Val
			l1 = l1.Next
			continue
		}
		if l2 != nil {
			tail.Val = l2.Val
			l2 = l2.Next
			continue
		}
	}
}
