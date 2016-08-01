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
