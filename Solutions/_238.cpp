#include <vector>
#include <iostream>

using namespace std;

vector<int> productExceptSelf(vector<int> &nums)
{
    vector<int> result(nums.size(), 1);
    int prefix = 1;
    for (size_t i = 0; i < nums.size(); i++)
    {
        result[i] = prefix;
        prefix *= nums[i];
    }
    int postfix = 1;
    for (int i = nums.size() - 1; i >= 0; i--)
    {
        result[i] *= postfix;
        postfix *= nums[i];
    }
    return result;
}

int main()
{
    auto input = vector<int>{1, 2, 3, 4};
    auto result = productExceptSelf(input);
    for (auto num : result)
    {
        cout << num << " ";
    }
}