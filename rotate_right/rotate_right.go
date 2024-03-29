package rotate_right

//给定一个链表，旋转链表，将链表每个节点向右移动 k 个位置，其中 k 是非负数。
//
//示例 1:
//
//输入: 1->2->3->4->5->NULL, k = 2
//输出: 4->5->1->2->3->NULL
//解释:
//向右旋转 1 步: 5->1->2->3->4->NULL
//向右旋转 2 步: 4->5->1->2->3->NULL
//示例 2:
//
//输入: 0->1->2->NULL, k = 4
//输出: 2->0->1->NULL
//解释:
//向右旋转 1 步: 2->0->1->NULL
//向右旋转 2 步: 1->2->0->NULL
//向右旋转 3 步: 0->1->2->NULL
//向右旋转 4 步: 2->0->1->NULL
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/rotate-list
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) (result *ListNode) {
	if head == nil || k == 0 {
		result = head
		return
	}
	var fast = head

	var count = 0
	for i := 0; i < k; i++ {
		fast = fast.Next
		if fast == nil {
			fast = head
			count = i + 1
			i = -1
			k = k % count
		}
	}
	if fast == nil || fast == head {
		result = head
		return
	}
	var slow = head
	for {
		if fast.Next == nil {
			result = slow.Next
			fast.Next = head
			slow.Next = nil
			return
		}
		fast = fast.Next
		slow = slow.Next
	}
}
