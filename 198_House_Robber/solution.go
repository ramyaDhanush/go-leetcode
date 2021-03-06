package _198_House_Robber

/*https://leetcode.com/problems/house-robber/

You are a professional robber planning to rob houses along a street. Each house has a certain amount of money stashed, the only constraint stopping you from robbing each of them is that adjacent houses have security system connected and it will automatically contact the police if two adjacent houses were broken into on the same night.

Given a list of non-negative integers representing the amount of money of each house, determine the maximum amount of money you can rob tonight without alerting the police.

Example 1:

Input: [1,2,3,1]
Output: 4
Explanation: Rob house 1 (money = 1) and then rob house 3 (money = 3).
             Total amount you can rob = 1 + 3 = 4.
Example 2:

Input: [2,7,9,3,1]
Output: 12
Explanation: Rob house 1 (money = 2), rob house 3 (money = 9) and rob house 5 (money = 1).
             Total amount you can rob = 2 + 9 + 1 = 12.
*/

/* Description
This task can be solve recursively - as brute force and with memoization.
Rob(nums) = max(nums[0]+Rob(nums[2:]), Rob(nums[1:]))
*/
func rob(nums []int) int {
    l := len(nums)
    if l == 0 {
        return 0
    }
    if l == 1 {
        return nums[0]
    }
    if l <= 2 {
        return max(nums[0], nums[1])
    }
    memo := make([]int,l) // index - is the beginning index range.

    memo[l-1] = nums[l-1]
    memo[l-2] = max(nums[l-2], nums[l-1])

    for i := l-3; i >= 0; i-- {
        memo[i] = max(nums[i]+memo[i+2], memo[i+1])
    }
    return max(memo[0], memo[1])
}

func max(x int, y int) int {
    if x > y {
        return x
    }
    return y
}
