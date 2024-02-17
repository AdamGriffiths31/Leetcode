#include <iostream>
#include <vector>
#include <algorithm>

using namespace std;

vector<vector<int>> threeSum(vector<int> &nums)
{
    sort(nums.begin(), nums.end());
    vector<vector<int>> result;

    for (int i = 0; i < nums.size(); i++)
    {
        if (i > 0 && nums[i] == nums[i - 1])
        {
            continue;
        }
        int left = i + 1;
        int right = nums.size() - 1;

        while (left < right)
        {
            int threeSum = nums[i] + nums[left] + nums[right];
            if (threeSum > 0)
            {
                right--;
            }
            else if (threeSum < 0)
            {
                left++;
            }
            else
            {
                result.push_back({nums[i], nums[left], nums[right]});
                left++;
                while (left < right && nums[left] == nums[left - 1])
                {
                    left++;
                }
            }
        }
    }

    return result;
}

int main()
{
    vector<int> nums = {-1, 0, 1, 2, -1, -4};
    vector<vector<int>> result = threeSum(nums);
    for (auto vec : result)
    {
        for (int val : vec)
        {
            cout << val << " ";
        }
        cout << endl;
    }
    return 0;
}