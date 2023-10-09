#include <iostream>
#include <vector>
#include <unordered_map>

using namespace std;

std::vector<int> twoSum(std::vector<int>& nums, int target) {
    std::unordered_map<int, int> numToIndex;
    std::vector<int> result;

    for (int i = 0; i < nums.size(); ++i) {
        int complement = target - nums[i];

        if (numToIndex.find(complement) != numToIndex.end()) {
            result.push_back(numToIndex[complement]);
            result.push_back(i);
            return result;
        }

        numToIndex[nums[i]] = i;
    }

    return result;
}

int main(){
    ios::sync_with_stdio(0);
    cin.tie(0);
    
    int numTestCases;
    cin >> numTestCases;

    for (int testCase = 0; testCase < numTestCases; testCase++) {
        cout << "Test Case " << testCase + 1 << "\n";
        int target, n;
        cin >> target >> n;

        vector<int> v(n); 

        for (int i = 0; i < n; i++){
            cin >> v[i];
        }

        vector<int> result = twoSum(v, target);

        if (result.size() == 2) {
            cout << "Test Case " << testCase + 1 << ": Indices of numbers that sum to " << target << ": " << result[0] << " " << result[1] << "\n";
        } else {
            cout << "Test Case " << testCase + 1 << ": No two numbers found that sum to " << target << "\n";
        }
    }

    return 0;
}
