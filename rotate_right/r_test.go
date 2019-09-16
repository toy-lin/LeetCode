package rotate_right

import "testing"

func TestRotateRight(t *testing.T) {
	rotateRight(&ListNode{
		Val: 0,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val:  2,
				Next: nil,
			},
		},
	}, 4)
}
