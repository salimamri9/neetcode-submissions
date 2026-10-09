class Solution:
    def kIsValid(self, k: int, piles: List[int], h:int) -> int:
        s: int = 0
        for i in range(len(piles)):
            s = s + math.ceil(piles[i] / k ) 
        return h - s
        
    def minEatingSpeed(self, piles: List[int], h: int) -> int:
        r = max(piles)
        l = 1 
        k = 0
        while(r >= l):
            mid = (l + r) // 2
            if(self.kIsValid(mid, piles, h)<0):
                l = mid + 1
            else: 
                r = mid - 1
                k = mid
        return k