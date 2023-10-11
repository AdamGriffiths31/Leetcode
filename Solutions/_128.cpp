#include <iostream>
#include <vector>
#include <unordered_map>
#include <unordered_set>

using namespace std;

int longestConsecutive(vector<int>& nums) {
    unordered_set<int> numSet;
    for (int num : nums) {
        numSet.insert(num);
    }

    int longestStreak = 0;
    for (int num : numSet) {
        if (numSet.count(num - 1) == 0) {
            int currentNum = num;
            int currentStreak = 1;

            while (numSet.count(currentNum + 1) > 0) {
                currentNum++;
                currentStreak++;
            }

            longestStreak = max(longestStreak, currentStreak);
        }
    }
    return longestStreak;
}

int main(){
    vector<int> arr = {100,4,200,1,3,2};
    int result = longestConsecutive(arr);
    cout << "result = " << result << "\n";
    
    arr = {0,3,7,2,5,8,4,6,0,1};
    result = longestConsecutive(arr);
    cout << "result = " << result << "\n";

    return 0;
}