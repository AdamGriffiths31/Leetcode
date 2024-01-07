#include <string>
#include <vector>
#include <iostream>

using namespace std;

bool isValid(string s)
{
    vector<char> stack;
    for (char c : s)
    {
        if (c == '(')
        {
            stack.push_back(')');
        }
        else if (c == '[')
        {
            stack.push_back(']');
        }
        else if (c == '{')
        {
            stack.push_back('}');
        }
        else if (stack.empty() || stack.back() != c)
        {
            return false;
        }
        else
        {
            stack.pop_back();
        }
    }
    return stack.empty();
}

int main()
{
    auto input = string{"(()))[]{}"};
    auto result = isValid(input);
    cout << result << "\n";
}