class Solution:
    def minEatingSpeed(self, piles: [int], h: int) -> int:
        def good_limit(speed):
            time = []
            for pile in piles:
                time.append(((pile - 1)//speed + 1))

            if(sum(time) <= h):
                return True
            else:
                return False
            
        
        left = 1
        right = max(piles)

        while(left < right):
            mid = left + (right - left)//2

            if(good_limit(mid)):
                right = mid
            else:
                left = mid+1
        
        return left