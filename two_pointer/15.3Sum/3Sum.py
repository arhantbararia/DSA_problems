class Solution:
    def threeSum(self, nums: [int]) -> [[int]]:
        nums.sort()
        triplets = set()

        res = []

        for i in range(0 , len(nums)):
            j = i+ 1
            k = len(nums)-1

            while(j < k ):
                ##condition
                triplet_sum = nums[i] + nums[j] + nums[k]
                if( triplet_sum == 0):
                    triplets.add((nums[i] , nums[j] , nums[k]))
                    j += 1
                    k += 1
                if(triplet_sum < 0):
                    j +=1
                else:
                    k -=1
            
        for triplet in triplets:
            res.append(triplet)
        
        return res
                
