package _747_Largest_Number_At_Least_Twice_of_Others

/* https://leetcode.com/problems/largest-number-at-least-twice-of-others/description/

In a given integer array nums, there is always exactly one largest element.

Find whether the largest element in the array is at least twice as much as every other number in the array.

If it is, return the index of the largest element, otherwise return -1.

Example 1:

Input: nums = [3, 6, 1, 0]
Output: 1
Explanation: 6 is the largest integer, and for every other number in the array x,
6 is more than twice as big as x.  The index of value 6 is 1, so we return 1.


Example 2:

Input: nums = [1, 2, 3, 4]
Output: -1
Explanation: 4 isn't at least as big as twice the value of 3, so we return -1.


Note:

nums will have a length in the range [1, 50].
Every nums[i] will be an integer in the range [0, 99].
*/

func dominantIndex(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	var max, secondMax, index int
	for i, iv := range nums {
		if iv > max {
			max, secondMax = iv, max
			index = i
		} else if iv > secondMax {
			secondMax = iv
		}
	}
	if secondMax == 0 || float64(max)/float64(secondMax) >= 2 {
		return index
	}

	return -1
}
