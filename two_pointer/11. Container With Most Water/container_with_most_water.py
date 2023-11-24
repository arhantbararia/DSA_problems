class Solution:
    def maxArea(self, height: List[int]) -> int:
        
        water = 0

        left = 0
        right = len(height)-1

        while(left < right ):
            length = right - left
            breadth = min(height[left] , height[right])

            water = max(water , length * breadth)

            if (height[left] < height[right]):
                left += 1
            else:
                right -= 1

        return water 
        
        # for i in range(0 , len(height)):
        #     for j in range( i + 1 , len(height)):
        #         length = j - i
        #         breadth = min(height[i] , height[j])

        #         water = max(water , length* breadth)


