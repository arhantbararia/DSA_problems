class Solution:
    def minDays(self, bloomDay: [int], m: int, k: int) -> int:
        
        def feasible(days) -> bool:
            bonquets=0
            flowers = 0
            for bloom in bloomDay:
                if bloom > days:
                    flowers = 0
                else:
                    bonquets += (flowers + 1) // k  ########## this is a good trick
                    flowers = (flowers + 1) % k    ##############
            return bonquets >= m

        if len(bloomDay) < m * k:
            return -1
        left, right = min(bloomDay), max(bloomDay)
        while left < right:
            mid = left + (right - left) // 2
            if feasible(mid):
                right = mid
            else:
                left = mid + 1
        return left


