package leet

type ListNode struct {
	Val int
	Next *ListNode
}

func InitList(arr []int) *ListNode  {
	count := len(arr)
	if count == 0 {
		return nil
	}

	root := &ListNode{
		Val:arr[0],
		Next:nil,
	}
	pre := root
	for i:= 1; i < count; i++ {
		pre.Next = &ListNode{
			Val:arr[i],
			Next:nil,
		}
		pre = pre.Next
	}

	return root
}

func Reverse(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	var pre, current, next *ListNode = nil, head, nil
	for current != nil {
		next = current.Next
		current.Next = pre
		pre = current
		current = next
	}

	return pre
}