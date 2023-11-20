class Solution:
    def fourSum(self, nums: [int], target: int) -> [[int]]:
        nums.sort()
        # [-2, -1, 0 , 1 , 2]

        quadruplets = set()
        res = []
        n = len(nums)
        for i in range(0, n-3):
            for j in range(i+1, n-2):
                k = j + 1
                l = n - 1
                while k < l:
                    #condition
                    total = nums[i] + nums[j] + nums[k] + nums[l]
                    if(total == target):
                        quadruplets.add((nums[i] , nums[j] , nums[k] , nums[l] ))
                        k += 1
                        l -= 1
                    elif(total > target):
                        l -= 1
                    else:
                        k += 1
        for quadruplet in quadruplets:
            res.append(quadruplet)
        
        return res
            


        
