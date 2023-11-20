/// input is BST therefore sorted array through in order traversal

/**
 * Definition for a binary tree node.
 
 };
 */
#include <vector>


struct TreeNode {
     int val;
     TreeNode *left;
     TreeNode *right;
     TreeNode() : val(0), left(nullptr), right(nullptr) {}
     TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
     TreeNode(int x, TreeNode *left, TreeNode *right) : val(x), left(left), right(right) {}


class Solution {
public:
    vector<int> sortedVector;

    vector<int> makeList(TreeNode* root ){
        if(root != nullptr){
            makeList(root->left);
            sortedVector.push_back(root->val);
            makeList(root->right );

        }
        return nums;

    }

    bool findTarget(TreeNode* root, int k) {
        vector<int> nums;

        nums = makeList(root , nums );
        
        left = 0;
        right = nums.size()-1;

        while(left < right){
            //condition
            total = nums[left] + nums[right];
            if(total == target){
                return true;
            }
            else if(total < target){
                left += 1;
            }
            else{
                right -= 1;
            }

        }
        return false;

        
        
    }
};