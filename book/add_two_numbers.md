# Add Two Numbers
	You are given two linked lists representing two non-negative numbers. The digits are stored in reverse order and each of their nodes contain a single digit. Add the two numbers and return it as a linked list.

	Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
	Output: 7 -> 0 -> 8
	
# 題目大意：
	給你兩個存儲非負數的鏈表，每個鏈表對應的位置都存儲一個簡單的數字，將對應的數字加起來行成一個新的鏈表並返回。

# 思路：
	1、參數檢查，l1或者l2為空的情況
	2、對應為相加，不管是否大於10先存在當前節點，下一個相加的時候判斷上一個節點的結果是否大於10，是的話則進位
	3、注意兩個鏈表長度不一致的情況


#代碼：
```golang
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	res := new(ListNode)

	res.Val = l1.Val + l2.Val

	tmpl1 := l1.Next
	tmpl2 := l2.Next
	tmpres := res
	for tmpl1 != nil || tmpl2 != nil {
		var tmpVal int
		if tmpl1 == nil {
			tmpVal = tmpl2.Val
		} else if tmpl2 == nil {
			tmpVal = tmpl1.Val
		} else {
			tmpVal = tmpl1.Val + tmpl2.Val
		}

		if tmpres.Val >= 10 {
			tmpres.Val %= 10
			tmpres.Next = &ListNode{
				Val: tmpVal + 1,
			}
		} else {
			tmpres.Next = &ListNode{
				Val: tmpVal,
			}
		}
		tmpres = tmpres.Next
		if tmpl1 != nil {
			tmpl1 = tmpl1.Next
		}

		if tmpl2 != nil {
			tmpl2 = tmpl2.Next
		}
	}

	if tmpres.Val >= 10 {
		tmpres.Val %= 10
		tmpres.Next = &ListNode{
			Val: 1,
		}
	}

	return res
}
```

[上一題](https://github.com/qianlnk/leetcode/blob/master/book/two_sum.md "Two Sum")|[目錄](https://github.com/qianlnk/leetcode/blob/master/README.md)|[下一題](https://github.com/qianlnk/leetcode/blob/master/book/Longest_Substring_Without_Repeating_Characters.md "Longest Substring Without Repeating Characters")
:------------: |:----------:| :-----------:
