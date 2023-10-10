#include <vector>
#include <iostream>
#include <unordered_set>

using namespace std;

bool containsDuplicate(vector<int>& nums) {
    unordered_set<int> seen;
    for (int num : nums) {
        if (seen.count(num) > 0){
            return true;
        }
        seen.insert(num);
    }

    return false;
}

int main(){
    ios::sync_with_stdio(0);
    cin.tie(0);
    
    int numTestCases;
    cin >> numTestCases;
    cout << "numTestCases = " << numTestCases << "\n";
    for (int testCase = 0; testCase < numTestCases; testCase++) {
        cout << "Test Case " << testCase + 1 << "\n";
        int n;
        cin >> n;

        vector<int> v(n); 

        for (int i = 0; i < n; i++){
            cin >> v[i];
        }
        bool result = containsDuplicate(v);
        cout << "Test Case " << testCase + 1 << ": " << result << "\n";
    }
}