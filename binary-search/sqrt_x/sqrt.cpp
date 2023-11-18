class Solution {
public:
    int mySqrt(int x) {
        if(x <= 1){
            return x;
        }
        long left = 0;
        long right = long(x) +1;
        long mid= 0;

        while(left < right ){
            mid = left + int((right - left)/2);
            if( mid > int(x/mid) ){
                right = mid;
            }
            else{
                left = mid + 1;
            }

        }
        return left-1;
    }
};