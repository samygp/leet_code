"""
Given an array nums, write a function to move all 0's to the end of it while maintaining the relative order of the non-zero elements.

Example:
Input: [0,1,0,3,12]
Output: [1,3,12,0,0]
You must do this in-place without making a copy of the array.
Minimize the total number of operations.
"""

def move_zeros(nums):
    i, j = 0, len(nums)-1
    while i < j:
        while nums[i] != 0:
            i += 1
        while i < j and nums[j] == 0:
            j -= 1
        if i < j:
            tmp = nums[i]
            nums[i] = nums[j]
            nums[j] = tmp
    print(nums)

nums = [0, 0, 0, 2, 0, 1, 3, 4, 0, 0]
move_zeros(nums)
nums = [0,1,0,3,12]
move_zeros(nums)
