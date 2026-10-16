class Solution:
    def findMinIndex(self, nums: List[int]) -> int:
        r = len(nums) - 1
        l = 0 

        while (l<r):
            mid = (l+r) // 2
            
            if nums[mid] > nums[r]:
                l = mid + 1
            else:
                r = mid
                
        return l

    def search(self, nums: List[int], target: int) -> int:
        l = 0
        r = len(nums) - 1
        min = self.findMinIndex(nums)
        
        if target > nums[l] and nums[l] > nums[r]:
            r = min - 1
        elif target == nums[l]:
            return l
        else:
            l = min
        
        while(l<=r):
            mid = (l + r) // 2
            if target == nums[mid]:
                return mid
            if (target > nums[mid]):
                l = mid + 1
            else:
                r = mid - 1 

        return -1
        