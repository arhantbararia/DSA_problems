class Solution:
    def searchInsert(self, nums: [int], target: int) -> int:
        left = 0
        right = len(nums)-1
        
        if target < nums[left]:
            return 0
        
        if target > nums[right]:
            return right + 1
        
        while(left < right):
            mid = (left + right)//2
            if nums[mid] > target:
                right = mid
            elif nums[mid] < target:
                left = mid + 1
            else:
                return mid

        return left
            