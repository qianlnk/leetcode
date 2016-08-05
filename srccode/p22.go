package leetcode

//Given n pairs of parentheses, write a function to generate all combinations of well-formed parentheses.

//For example, given n = 3, a solution set is:

//[
//  "((()))",
//  "(()())",
//  "(())()",
//  "()(())",
//  "()()()"
//]

func generateParenthesis(n int) []string {
	var space [][]int
	for i := 0; i < n; i++ {
		var tmpspace []int
		for j := 0; j <= i+1; j++ {
			if i == n-1 {
				tmpspace = append(tmpspace, j+1)
			} else {
				tmpspace = append(tmpspace, j)
			}
		}
		space = append(space, tmpspace)
	}

	if len(space) == 0 {
		return nil
	}

	k := 0

	for _, v := range space[k] {

	}
}
