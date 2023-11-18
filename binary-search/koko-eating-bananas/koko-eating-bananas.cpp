class Solution {
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
public:

    
    bool good_eating_rate(int speed, vector<int>& piles, int h){
        vector<int> time;
        for(int i = 0 ; i < piles.size(); i++){
            int time_taken = ((piles[i]-1)/speed) +1;
            time.push_back(time_taken);
        }

        if(sum(time) <= h){
            return true;
        }
        return false;
    }
    int minEatingSpeed(vector<int>& piles, int h) {

        int left = 1;
        int right = max(piles);

        int mid = 0;

        while(left < right){
            mid = left + int((right-left)/2) ;
            if(good_eating_rate(mid , piles , h)){
                right = mid;
            }
            else{
                left = mid+1;
            }
        }
        return left;
    }
};