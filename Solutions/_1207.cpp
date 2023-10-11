#include <iostream>
#include <vector>
#include <unordered_map>
#include <unordered_set>

using namespace std;

bool uniqueOccurrences(vector<int>& arr) {
    unordered_map<int, int> occurrences;
    unordered_set<int> uniqueCounts;
    
    for (int num : arr) {
        occurrences[num]++;
        if (occurrences[num] < 0) {
            return false;
        }
    }
    for (const auto& entry : occurrences) {
        if (uniqueCounts.count(entry.second) > 0) {
            return false; 
        }
        uniqueCounts.insert(entry.second);
    }

    return true;
}

int main(){
    vector<int> arr = {1, 2, 2, 1, 1, 3};
    bool result = uniqueOccurrences(arr);
    cout << "result = " << result << "\n";

    arr = {1,2};
    result = uniqueOccurrences(arr);
    cout << "result = " << result << "\n";

    arr = {-3,0,1,-3,1,1,1,-3,10,0};
    result = uniqueOccurrences(arr);
    cout << "result = " << result << "\n";
    
    return 0;
}