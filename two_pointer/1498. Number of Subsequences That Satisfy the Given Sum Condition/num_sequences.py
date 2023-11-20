class Solution:
    def numSubseq(self, nums: [int], target: int) -> int:
        MOD = 10**9 + 7
        result = 0

        two_multiples = []

        two_multiples[0] = 2**0

        for i in range(1, n-1):
            res = 2*two_multiples[i-1] % MOD
            two_multiples.append(res)
            
        print(two_multiples)

        nums.sort()
        n = len(n)
        left = 0
        right = n-1

        while(left <= right):
            #condition
            total = nums[left] + nums[right]
            if (total <= target):
                result = (result + two_multiples[right-left])%MOD
                left += 1
            else:
                right -= 1
        return result 
                
