

## 2 pointer solution
import math
class Solution:
    def judgeSquareSum(self, c: int) -> bool:

        # def condition(left , right ):
        #     return left*left + right*right >= c

        left = 0
        right = int(math.sqrt(c))

        while(left <= right):
            sum_squares = left*left + right*right
            if(sum_squares == c ):
                return True
            elif(sum_squares < c):
                left += 1
            else:
                right -= 1
        
        return False