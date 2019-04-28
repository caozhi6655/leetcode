package _1Merge_Two_Sorted_Lists

type ListNode struct {
	Val  int
	Next *ListNode
}

func MergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	result := &ListNode{}
	root := result

	for {
		if l1 == nil || l2 == nil {
			break
		}

		if l1.Val < l2.Val {
			root.Val = l1.Val
			l1 = l1.Next
		} else {
			root.Val = l2.Val
			l2 = l2.Next
		}
		root.Next = &ListNode{}
		root = root.Next
	}

	if l1 != nil {
		root.Val = l1.Val
		root.Next = l1.Next
	}

	if l2 != nil {
		root.Val = l2.Val
		root.Next = l2.Next
	}

	return result
}
