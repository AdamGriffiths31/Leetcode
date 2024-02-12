#include <iostream>
#include <vector>

using namespace std;

int missingNumber(vector<int>& nums) {
    int totalSum = nums.size() * (nums.size() + 1) / 2;
    for (int i = 0; i < nums.size(); i++){
        totalSum -= nums[i];
    }
    return totalSum;
}

int main() {
    vector<int> nums = {9,6,4,2,3,5,7,0,1};
    cout << missingNumber(nums) << endl;
    return 0;
}