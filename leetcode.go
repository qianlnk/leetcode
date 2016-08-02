package leetcode

import (
	"fmt"
	"strings"
)

//problem1
func twoSum(nums []int, target int) []int {
	for i, a := range nums {
		for j, b := range nums {
			if i == j {
				continue
			}
			if (a + b) == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

//problem2
//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
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

//problem3
func lengthOfLongestSubstring(s string) int {
	length := 0
	bs := []byte(s)
	for offset := 0; offset < len(bs); offset++ {
		var tmpbs []byte
		tmplen := 0
		for i := offset; i < len(bs); i++ {
			if strings.Contains(string(tmpbs), string(bs[i])) == false {
				tmplen++
				tmpbs = append(tmpbs, bs[i])
			} else {
				break
			}
		}

		if tmplen > length {
			length = tmplen
		}
	}
	return length
}

//problem4
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var nums []int
	len1 := len(nums1)
	len2 := len(nums2)
	i1 := 0
	i2 := 0

	if len1 == 0 && len2 == 0 {
		return 0
	} else if len1 == 0 {
		nums = nums2
	} else if len1 == 0 {
		nums = nums1
	} else {
		for {
			if nums1[i1] < nums2[i2] {
				nums = append(nums, nums1[i1])
				i1++
			} else {
				nums = append(nums, nums2[i2])
				i2++
			}

			if i2 >= len2 {
				nums = append(nums, nums1[i1:]...)
				break
			} else if i1 >= len1 {
				nums = append(nums, nums2[i2:]...)
				break
			}
		}
	}

	if len(nums)%2 != 0 {
		return float64(nums[len(nums)/2])
	} else {
		return (float64(nums[len(nums)/2-1]) + float64(nums[len(nums)/2])) / 2
	}
}

//problem5
func longestPalindrome(s string) string {
	var res []byte
	bs := []byte(s)
	for offset := 0; offset < len(s); offset++ {
		if len(s)-len(res) < offset {
			break
		}
		var tmpbs []byte
		for i := offset; i < len(bs); i++ {
			tmpbs = append(tmpbs, bs[i])
			if len(tmpbs) <= len(res) {
				continue
			}

			pal := true
			for j := 0; j < len(tmpbs)/2; j++ {
				if tmpbs[j] != tmpbs[len(tmpbs)-j-1] {
					pal = false
					break
				}
			}

			if pal {
				res = tmpbs
			}
		}
	}
	return string(res)
}

//problem6
func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	bs := []byte(s)
	var tmpres [][]byte
	vertical := true
	offset := 0
	rowOffset := numRows
	for offset < len(bs) {
		if numRows == 2 {
			vertical = true
		}
		if vertical {
			for j := 0; j < numRows; j++ {
				if offset == j {
					tr := []byte{bs[offset]}
					tmpres = append(tmpres, tr)
				} else {
					tmpres[j] = append(tmpres[j], bs[offset])
				}
				offset++
				if offset >= len(bs) {
					break
				}
			}
			vertical = false
		} else {
			for j := numRows - 1; j >= 0; j-- {
				tmpres[j] = append(tmpres[j], '*')
			}
			tmpres[rowOffset-2][len(tmpres[rowOffset-2])-1] = bs[offset]
			offset++
			rowOffset--
			if rowOffset == 2 {
				vertical = true
				rowOffset = numRows
			}
		}
	}
	var res []byte
	for i := 0; i < len(tmpres); i++ {
		for j := 0; j < len(tmpres[i]); j++ {
			if tmpres[i][j] != '*' {
				res = append(res, tmpres[i][j])
			}
		}
	}
	return string(res)
}

//problem8
func myAtoi(str string) int {
	bs := []byte(str)
	var res int64
	var step int64 = 1
	var tmpres int64
	for i := len(bs) - 1; i >= 0; i-- {
		fmt.Println(bs[i], tmpres)
		if bs[i] >= '0' && bs[i] <= '9' {
			tmpres += (int64(bs[i]) - int64('0')) * step
			step *= 10
		} else if bs[i] == '+' {
			res = tmpres
			tmpres = 0
			step = 1
		} else if bs[i] == ' ' {
			if tmpres != 0 {
				res = tmpres
			}
			tmpres = 0
			step = 1
		} else if bs[i] == '-' {
			res = tmpres * -1
			tmpres = 0
			step = 1
		} else if bs[i] == '.' {
			tmpres = 0
			step = 1
		} else {
			tmpres = 0
			step = 1
		}
	}

	if tmpres != 0 {
		res = tmpres
	}
	fmt.Println(int64(res))
	if res > 2147483647 {
		return 2147483647
	} else if res < -2147483648 {
		return -2147483648
	}
	return int(res)
}
