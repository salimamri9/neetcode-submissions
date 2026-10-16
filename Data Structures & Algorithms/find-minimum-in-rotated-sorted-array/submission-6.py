class Solution:
    def findMin(self, nums: List[int]) -> int:
        size = len(nums)
        if size == 1:
            return nums[0]
            
        if nums[0] < nums[size-1]:
            return nums[0]

        l = 0 
        r = size - 1

        while (l<r):
            mid = (l+r) // 2
            
            if nums[mid] > nums[r]:
                l = mid + 1
            else:
                r = mid
                
        return nums[l]

            
