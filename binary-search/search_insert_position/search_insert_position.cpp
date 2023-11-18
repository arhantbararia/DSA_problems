#include<vector.h>



class Solution {
public:
    int searchInsert(vector<int>& nums, int target) {
        int left = 0;
        int right = nums.size() - 1;
        int mid = 0;
        if(target < nums[left]){
            return 0;
        }

        if(target > nums[right]){
            return right + 1;
        }

        while(left < right){
            mid = left + int((right-left) / 2 );
            if(nums[mid] > target){
                right = mid;

            }
            else if(nums[mid]  < target){
                left = mid + 1;
            }
            else if(nums[mid] == target){
                return mid;
            }

        }
        return left;

    }
};