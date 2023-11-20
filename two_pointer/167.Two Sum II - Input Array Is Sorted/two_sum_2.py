class Solution:
    def twoSum(self, numbers: [int], target: int) -> [int]:
        def condition(left , right):
            #some condition with good flag handle
            pass
        

        left = 0
        right = len(numbers)-1

        while(left < right):

            flag = condition(numbers[left] , numbers[right])
            if(flag == 0):
                return [left+1  , right+1]
            elif(flag == 1 ):
                right -= 1
            else:
                left += 1
        
        return [-1, -1]
            

