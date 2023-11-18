#include<vector.h>



class Solution {
public:
    int max(vector<int>& arr){
            int max_arr = arr[0];

            for(int i = 0 ; i < arr.size() ; i++){
                if(arr[i] > max_arr){
                    max_arr = arr[i];
                }
                
            }
            return max_arr;
        }


        int sum(vector<int>& arr){
            int total = 0;
            for(int i = 0 ; i< arr.size(); i++){
                total += arr[i];
            }

            return total;
        }

        bool is_feasible(int capacity , vector<int>& weights , int days){
            int d = 1;
            int total = 0;
            
            for(int i = 0 ; i < weights.size() ; i++){
                total += weights[i];
                if(total > capacity){
                    total = weights[i];
                    d++ ;
                    if(d > days){
                        return false;
                    }
                }
            }
            return true;
        }

    int shipWithinDays(vector<int>& weights, int days) {

        

        int left = max(weights);
        int right = sum(weights);
        int mid = 0;
        while(left  < right){
            mid = left + int((right - left)/2);
            if(is_feasible(mid, weights , days)){
                right = mid;
            }
            else{
                left = mid + 1;

            }
            
        }

        return left;
    }
};