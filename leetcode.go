package leetcode

import (
	"fmt"
	"sort"
	"strings"
)

//1
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

//2
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

//3
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

//4
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

//5
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

//6
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

//8
func myAtoi(str string) int {
	const (
		MAX_INT = 2147483647
		MIN_INT = -2147483648
	)
	bs := []byte(strings.Trim(str, " "))
	var res int64
	var sign int64 = 1
	var tmpres int64
	for i := 0; i < len(bs); i++ {
		if bs[i] >= '0' && bs[i] <= '9' {
			tmpres *= 10
			tmpres += (int64(bs[i]) - int64('0'))

			if tmpres > MAX_INT {
				break
			}
		} else if bs[i] == '+' && i == 0 {
			sign = 1
		} else if bs[i] == '-' && i == 0 {
			sign = -1
		} else {
			break
		}
	}
	res = tmpres * sign
	if res > MAX_INT {
		return MAX_INT
	} else if res < MIN_INT {
		return MIN_INT
	}
	return int(res)
}

//9
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	tmp := x
	var n int

	for tmp != 0 {
		n = n*10 + tmp%10
		tmp = tmp / 10
	}

	return x == n
}

//10
func isMatch(s string, p string) bool {
	if p == "" {
		return s == ""
	}
	bp := []byte(p)
	bs := []byte(s)
	bp = append(bp, 0)
	bs = append(bs, 0)
	if bp[1] != '*' {
		if bp[0] == '*' {
			return false
		}
		return ((bp[0] == bs[0]) || (bp[0] == '.' && bs[0] != 0)) && isMatch(string(bs[1:len(bs)-1]), string(bp[1:len(bp)-1]))
	}

	for (bp[0] == bs[0]) || (bp[0] == '.' && bs[0] != 0) {
		if isMatch(string(bs[:len(bs)-1]), string(bp[2:len(bp)-1])) {
			return true
		}
		bs = bs[1:]
	}

	return isMatch(string(bs[:len(bs)-1]), string(bp[2:len(bp)-1]))
}

//11
func maxArea(height []int) int {
	var res int
	for i := 0; i < len(height)-1; i++ {
		for j := i + 1; j < len(height); j++ {
			var area int
			if height[i] > height[j] {
				area = height[j] * (j - i)
			} else {
				area = height[i] * (j - i)
			}

			if area > res {
				res = area
			}
		}
	}
	return res
}

//12
func intToRoman(num int) string {
	rmap := [4][10]string{
		{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"},
		{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"},
		{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"},
		{"", "M", "MM", "MMM"}}

	return rmap[3][num/1000%10] + rmap[2][num/100%10] + rmap[1][num/10%10] + rmap[0][num%10]
}

//13
func romanToInt(s string) int {
	fmt.Println(s)
	rmap := make(map[byte]int)
	rmap['I'] = 1
	rmap['V'] = 5
	rmap['X'] = 10
	rmap['L'] = 50
	rmap['C'] = 100
	rmap['D'] = 500
	rmap['M'] = 1000

	var res int
	bs := []byte(s)
	bs = append(bs, 0)
	if len(bs) <= 0 {
		return 0
	}

	res = rmap[bs[0]]

	for i := 0; bs[i] != 0; i++ {
		fmt.Println(rmap[bs[i]], rmap[bs[i+1]])
		if rmap[bs[i+1]] <= rmap[bs[i]] {
			res += rmap[bs[i+1]]
		} else {
			res = res + rmap[bs[i+1]] - 2*rmap[bs[i]]
		}
	}

	return res
}

//14
func longestCommonPrefix(strs []string) string {
	if len(strs) <= 0 {
		return ""
	} else if len(strs) == 1 {
		return strs[0]
	}

	bs := []byte(strs[0])
	var res string
	var tmpbs []byte
	for i := 0; i < len(bs); i++ {
		tmpbs = append(tmpbs, bs[i])
		for _, str := range strs {
			if len(tmpbs) > len(str) {
				return res
			}

			if string(tmpbs) != string([]byte(str)[0:len(tmpbs)]) {
				return res
			}
		}
		res = string(tmpbs)
	}

	return res

}

//15
func threeSum(nums []int) [][]int {
	var res [][]int
	for i := 0; i < len(nums)-2; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			for k := j + 1; k < len(nums); k++ {
				if nums[i]+nums[j]+nums[k] == 0 {
					var min, mid, max int
					if nums[i] <= nums[j] {
						if nums[k] <= nums[i] {
							min = k
							mid = i
							max = j
						} else if nums[k] >= nums[j] {
							min = i
							mid = j
							max = k
						} else {
							min = i
							mid = k
							max = j
						}
					} else {
						if nums[k] <= nums[j] {
							min = k
							mid = j
							max = i
						} else if nums[k] >= nums[i] {
							min = j
							mid = i
							max = k
						} else {
							min = j
							mid = k
							max = i
						}
					}

					var tmpres []int
					tmpres = append(tmpres, nums[min])
					tmpres = append(tmpres, nums[mid])
					tmpres = append(tmpres, nums[max])
					exist := false
					for _, v := range res {
						if v[0] == tmpres[0] && v[1] == tmpres[1] && v[2] == tmpres[2] {
							exist = true
						}
					}
					if exist == false {
						res = append(res, tmpres)
					}
				}
			}
		}
	}
	return res
}

func threeSumClosest(nums []int, target int) int {
	var res int
	offset := 2147483647
	for i := 0; i < len(nums)-2; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			for k := j + 1; k < len(nums); k++ {
				tmpoffset := nums[i] + nums[j] + nums[k] - target
				if tmpoffset < 0 {
					tmpoffset *= -1
				}
				if tmpoffset < offset {
					res = nums[i] + nums[j] + nums[k]
					offset = tmpoffset
				}
			}
		}
	}

	return res
}

func letterCombinations(digits string) []string {
	if len(digits) <= 0 {
		return nil
	}
	var res []string

	dcmap := [10][]string{
		{},
		{},
		{"a", "b", "c"},
		{"d", "e", "f"},
		{"g", "h", "i"},
		{"j", "k", "l"},
		{"m", "n", "o"},
		{"p", "q", "r", "s"},
		{"t", "u", "v"},
		{"w", "x", "y", "z"},
	}

	bds := []byte(digits)

	if bds[0] < '2' || bds[0] > '9' {
		return nil
	}

	if len(digits) == 1 {
		return dcmap[bds[0]-48]
	}

	offres := letterCombinations(string([]byte(digits)[1:]))

	for i := 0; i < len(dcmap[bds[0]-48]); i++ {
		for j := 0; j < len(offres); j++ {
			res = append(res, dcmap[bds[0]-48][i]+offres[j])
		}
	}

	return res
}

func fourSum(nums []int, target int) [][]int {
	var res [][]int
	for a := 0; a < len(nums)-3; a++ {
		for b := a + 1; b < len(nums)-2; b++ {
			for c := b + 1; c < len(nums)-1; c++ {
				for d := c + 1; d < len(nums); d++ {
					var tmpres sort.IntSlice
					if nums[a]+nums[b]+nums[c]+nums[d] == target {
						tmpres = append(tmpres, nums[a])
						tmpres = append(tmpres, nums[b])
						tmpres = append(tmpres, nums[c])
						tmpres = append(tmpres, nums[d])

						tmpres.Sort()
						exist := false
						for _, v := range res {
							if v[0] == tmpres[0] && v[1] == tmpres[1] && v[2] == tmpres[2] && v[3] == tmpres[3] {
								exist = true
							}
						}
						if exist == false {
							res = append(res, tmpres)
						}
					}
				}
			}
		}
	}
	return res
}

// Definition for singly-linked list.
//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil || n <= 0 {
		return head
	}
	var lstlen int
	tmpNode := head
	for tmpNode != nil {
		lstlen++
		tmpNode = tmpNode.Next
	}

	if lstlen == 1 {
		head.Val = 0
		head.Next = nil
		return nil
	}

	tmpNode = head

	if lstlen <= n {
		freeNode := tmpNode
		head = freeNode.Next
		freeNode.Val = 0
		freeNode.Next = nil
	} else {
		for i := 0; i < lstlen-n-1; i++ {
			tmpNode = tmpNode.Next
		}
		freeNode := tmpNode.Next
		fmt.Println(freeNode.Val)
		tmpNode.Next = freeNode.Next

		freeNode.Val = 0
		freeNode.Next = nil
	}
	return head
}

func isValid(s string) bool {
	brackets := []string{"(", ")", "[", "]", "{", "}"}
}
