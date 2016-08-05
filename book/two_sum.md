# Two Sum  QuestionEditorial Solution  My Submissions
	Total Accepted: 268340
	Total Submissions: 1060754
	Difficulty: Easy
	Given an array of integers, return indices of the two numbers such that they add up to a specific target.

	You may assume that each input would have exactly one solution.

	Example:
	Given nums = [2, 7, 11, 15], target = 9,

	Because nums[0] + nums[1] = 2 + 7 = 9,
	return [0, 1].
	
# 題目大意：
	給定一個整型數組和一個數字，返回數組中兩個相加等於給定數字的兩個數的下標。
	
# 思路1
	數組第一個數字和第二、三、四...相加如果等於target就return
	數組第二個數字和第三、四、五...相加
	...
	
# 代碼
```golang
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
```

# 思路2