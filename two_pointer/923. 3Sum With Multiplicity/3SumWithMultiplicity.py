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
    
(1, 1, 5) : (0, 1, 9)
(1, 2, 5) : (0, 2, 9)
(1, 2, 5) : (0, 2, 8)
(1, 2, 4) : (0, 2, 7)
(1, 2, 5) : (0, 3, 9)
(1, 2, 5) : (0, 3, 8)
(1, 2, 4) : (0, 3, 7)
(1, 3, 5) : (0, 4, 9)
(1, 3, 5) : (0, 4, 8)
(1, 3, 4) : (0, 4, 7)
(1, 3, 4) : (0, 4, 6)
(1, 3, 3) : (0, 4, 5)
(1, 3, 5) : (0, 5, 9)
(1, 3, 5) : (0, 5, 8)
(1, 3, 4) : (0, 5, 7)
(1, 3, 4) : (0, 5, 6)
(1, 2, 5) : (1, 2, 9)
(1, 2, 5) : (1, 2, 8)
(1, 2, 4) : (1, 2, 7)
(1, 2, 5) : (1, 3, 9)
(1, 2, 5) : (1, 3, 8)
(1, 2, 4) : (1, 3, 7)
(1, 3, 5) : (1, 4, 9)
(1, 3, 5) : (1, 4, 8)
(1, 3, 4) : (1, 4, 7)
(1, 3, 4) : (1, 4, 6)
(1, 3, 3) : (1, 4, 5)
(1, 3, 5) : (1, 5, 9)
(1, 3, 5) : (1, 5, 8)
(1, 3, 4) : (1, 5, 7)
(1, 3, 4) : (1, 5, 6)
(2, 2, 5) : (2, 3, 9)
(2, 2, 5) : (2, 3, 8)
(2, 2, 4) : (2, 3, 7)
(2, 2, 4) : (2, 3, 6)
(2, 2, 3) : (2, 3, 5)
(2, 3, 5) : (2, 4, 9)
(2, 3, 5) : (2, 4, 8)
(2, 3, 4) : (2, 4, 7)
(2, 3, 4) : (2, 4, 6)
(2, 3, 3) : (2, 4, 5)
(2, 3, 5) : (3, 4, 9)
(2, 3, 5) : (3, 4, 8)
(2, 3, 4) : (3, 4, 7)
(2, 3, 4) : (3, 4, 6)
(2, 3, 3) : (3, 4, 5)
(3, 3, 5) : (4, 5, 9)
(3, 3, 5) : (4, 5, 8)
(3, 3, 4) : (4, 5, 7)
(3, 3, 4) : (4, 5, 6)
(3, 4, 5) : (5, 6, 9)
(3, 4, 5) : (5, 6, 8)
(3, 4, 4) : (5, 6, 7)
(4, 4, 5) : (6, 7, 9)
(4, 4, 5) : (6, 7, 8)
(4, 5, 5) : (7, 8, 9)
{(1, 2, 5): 8, (1, 3, 4): 8, (2, 2, 4): 2, (2, 3, 3): 2}


(1, 1, 2) : (0, 1, 5)
(1, 2, 2) : (0, 2, 5)
(1, 2, 2) : (0, 2, 4)
(1, 2, 2) : (0, 2, 3)
(1, 2, 2) : (1, 2, 5)
(1, 2, 2) : (1, 2, 4)
(1, 2, 2) : (1, 2, 3)
(2, 2, 2) : (2, 3, 5)
(2, 2, 2) : (2, 3, 4)
(2, 2, 2) : (3, 4, 5)
{(1, 2, 2): 6}