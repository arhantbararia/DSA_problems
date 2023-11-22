class Solution:
    def numRescueBoats(self, people: [int], limit: int) -> int:
        people.sort()
        boats =  0
        left = 0
        right = len(people)-1
        
        
        while(left <= right):
            #condition
            sum_of_2 = people[left] + people[right]
            if(sum_of_2 <= limit):
                boats += 1
                left +=  1
                right -= 1

            else:
                
                right -= 1
                boats += 1
            
        return boats
            

