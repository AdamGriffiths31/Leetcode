#include <vector>
#include <iostream>
#include <unordered_map>
#include <map>

using namespace std;

vector<int> topKFrequent(vector<int> &nums, int k)
{
    unordered_map<int, int> occurrences;
    map<int, vector<int>> freq;
    for (int num : nums)
    {
        occurrences[num]++;
    }

    for (const auto &entry : occurrences)
    {
        freq[entry.second].push_back(entry.first);
    }

    vector<int> result;
    for (size_t i = occurrences.size()+1; i >= 0; i--)
    {
        result.insert(result.end(), freq[i].begin(), freq[i].end());
        if (result.size() == k)
        {
            break;
        }
    }

    return result;
}

int main()
{
    auto nums = vector<int>{1, 1, 1, 2, 2, 3};
    auto k = 2;
    auto result = topKFrequent(nums, k);
    for (auto num : result)
    {
        cout << num << " ";
    }
}