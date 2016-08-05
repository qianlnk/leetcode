package leetcode

import (
	"fmt"
	"testing"
)

func TestTowSum(t *testing.T) {
	s := []int{3, 2, 4}
	fmt.Println(twoSum(s, 6))
}

func TestLengthOfLongestSubstring(t *testing.T) {
	s := "pwwkew"
	fmt.Println(lengthOfLongestSubstring(s))
}

func TestFindMedianSortedArrays(t *testing.T) {
	n1 := []int{}
	n2 := []int{2}
	fmt.Println(findMedianSortedArrays(n1, n2))
}

func TestLongestPalindrome(t *testing.T) {
	fmt.Println(longestPalindrome("abacdeaabcbaaewqeasd"))
}

func TestConvert(t *testing.T) {
	fmt.Println(convert("PAYPALISHIRING", 3))
}

func TestMyAtoi(t *testing.T) {
	fmt.Println(myAtoi("9223372036854775809"))
}

//func TestCountIsland(t *testing.T) {
//	IslandMap := [10][]int{
//		{12, 14, 6, 7, 14, 6, 12, 4, 10, 21},
//		{7, 11, 6, 12, 32, 9, 11, 23, 5, 4},
//		{8, 9, 10, 21, 4, 3, 22, 11, 3, 11},
//		{11, 5, 12, 3, 5, 12, 21, 11, 1, 2},
//		{35, 7, 1, 9, 21, 4, 3, 11, 12, 12},
//		{5, 4, 3, 11, 8, 21, 6, 11, 15, 5},
//		{21, 4, 6, 12, 11, 7, 5, 8, 12, 2},
//		{9, 8, 10, 14, 32, 1, 5, 3, 12, 10},
//		{8, 7, 12, 5, 3, 2, 18, 21, 11, 5},
//		{19, 5, 4, 2, 11, 12, 6, 8, 2, 11}}
//	sealevel := 10

//	var visted [][]int
//}
func TestIsMatch(t *testing.T) {
	fmt.Println(isMatch("aava", "*.*a"))
}

func TestRomanToInt(t *testing.T) {
	fmt.Println(romanToInt(intToRoman(2217)))
}

func TestLongestCommonPrefix(t *testing.T) {
	strs := []string{"asd", "asfe", "as", "a"}
	fmt.Println(longestCommonPrefix(strs))
}

func TestThreeSum(t *testing.T) {
	s := []int{-1, 0, 1, 2, -1, -4, 2}
	fmt.Println(threeSum(s))
}

func TestThreeSumClosest(t *testing.T) {
	s := []int{0, 0, 0}
	fmt.Println(threeSumClosest(s, 1))
}

func TestLetterCombinations(t *testing.T) {
	fmt.Println(letterCombinations("2345"))
}

func TestFourSum(t *testing.T) {
	s := []int{-1, 0, 1, 2, -1, -4, 2}
	fmt.Println(fourSum(s, 0))
}

func TestRemoveNthFromEnd(t *testing.T) {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val:  4,
					Next: nil,
				},
			},
		},
	}

	res := removeNthFromEnd(head, -1)
	tmpres := res
	for tmpres != nil {
		fmt.Println(tmpres.Val)
		tmpres = tmpres.Next
	}
}

func TestIsValid(t *testing.T) {
	fmt.Println(isValid("123[{}([(11)])]23{}dsf"))
}
