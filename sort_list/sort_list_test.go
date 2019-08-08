package sort_list

import (
	"fmt"
	"testing"
)

func TestSortList(t *testing.T) {
	a := ListNode{Val: 0, Next: &ListNode{Val: 4, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 8, Next: &ListNode{Val: 2, Next: &ListNode{Val: 0}}}}}}}
	c := sortList(&a)
	fmt.Println(c)
}

func TestMerge(t *testing.T) {
	a := ListNode{Val: 0, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2}}}
	b := ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4}}}
	head, last := merge2Lists(&a, &b)
	fmt.Println(head,last)
}
