package merge_k_sorted_lists

import (
	"fmt"
	"testing"
	"time"
)

func TestMergeKLists(t *testing.T) {
	a := ListNode{Val: 0, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2}}}
	b := ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4}}}
	c := ListNode{Val: 4, Next: &ListNode{Val: 5, Next: &ListNode{Val: 6}}}
	fmt.Println(time.Now())
	fmt.Println(mergeKLists([]*ListNode{&a, &b, &c}))
	fmt.Println(time.Now())
}
