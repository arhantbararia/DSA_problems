class Solution:
    def threeSumMulti(self, nums: [int], target: int) -> int:
        nums.sort()
        triplets = {}

        

        for i in range(0 , len(nums)):
            j = i+ 1
            k = len(nums)-1

            while(j < k ):
                ##condition
                triplet_sum = nums[i] + nums[j] + nums[k]
                if( triplet_sum == target):
                    if ((i,j,k) not in triplets):
                        triplets = 1
                    else:
                        triplets += 1
                
                        
                    j += 1
                    k += 1
                elif(triplet_sum < target):
                    j +=1
                else:
                    k -=1
            
        res = 0

        for key in triplets:
            res += triplets[key]
        
        
        return  res
  