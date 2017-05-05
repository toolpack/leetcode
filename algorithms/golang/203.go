package main

func removeElements(head *ListNode, val int) *ListNode {
	//remove header
	for head != nil {
		if head.Val == val {
			head = head.Next
		} else {
			break
		}
	}
	if head == nil {
		return nil
	}

	cur := head
	for cur.Next != nil {
		if cur.Next.Val == val {
			cur.Next = cur.Next.Next
			continue
		}
		cur = cur.Next
	}
	return head
}
