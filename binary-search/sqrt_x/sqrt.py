
#we can easily do this by Binary search or much cooler version of binary search newton's method of approximation



class Solution:
    def mySqrt(self, x: int) -> int:
    
        left , right = 0, x + 1
        
        while(left < right):
            mid = left + (right - left )//2

            if mid * mid > x:
                right = mid
            else:
                left = mid + 1

        return left - 1
    
    ##Newton's version
    def mySqrt(self, x: int) -> int:
    
        g2 = x/2
        g1 = 0
        
        error = 0.0005
        
        while(abs(g2-g1) > error):
            g1 = g2
            g2 = (g1 + x /g2)/2
            
        
        return int(g2)