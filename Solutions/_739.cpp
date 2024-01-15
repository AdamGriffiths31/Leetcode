#include <vector>
#include <iostream>
#include <tuple>
using namespace std;

vector<int> dailyTemperatures(vector<int> &temperatures)
{
    vector<int> result(temperatures.size(), 0);
    vector<tuple<int, int>> stack;

    for (int i = 0; i < temperatures.size(); i++)
    {
        while (!stack.empty() && get<0>(stack.back()) < temperatures[i])
        {
            int temp, index;
            std::tie(temp, index) = stack.back();
            stack.pop_back();
            result[index] = i - index;
        }
        stack.push_back(make_tuple(temperatures[i], i));
    }
    return result;
}

int main()
{
    auto input = vector<int>{73, 74, 75, 71, 69, 72, 76, 73};
    auto output = dailyTemperatures(input);

    for (auto i : output)
    {
        cout << i << " ";
    }
}