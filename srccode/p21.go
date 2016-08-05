package leetcode

//Merge Two Sorted Lists
//Merge two sorted linked lists and return it as a new list. The new list should be made by splicing together the nodes of the first two lists

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	res := new(ListNode)
	tmpres := res
	for l1 != nil || l2 != nil {
		if l1 == nil {
			tmpres.Next = l2
			break
		}

		if l2 == nil {
			tmpres.Next = l1
			break
		}

		if l1.Val <= l2.Val {
			tmpres.Next = l1
			l1 = l1.Next
		} else {
			tmpres.Next = l2
			l2 = l2.Next
		}

		tmpres = tmpres.Next
	}
	return res.Next
}
