#include <iostream>
#include <vector>

using namespace std;

void backtrack(vector<string>& result, int openN, int closeN, int n, string str)
{
    if (openN == n && closeN == n)
    {
        result.push_back(str);
        return;
    }
    if (openN < n)
    {
        backtrack(result, openN + 1, closeN, n, str + "(");
    }
    if (openN > closeN)
    {
        backtrack(result, openN, closeN + 1, n, str + ")");
    }
}

vector<string> generateParenthesis(int n)
{
    vector<string> result;

    backtrack(result, 0, 0, n, "");
    return result;
}

int main()
{
    auto output = generateParenthesis(3);
    cout << "Output: " << endl;
    for (auto s : output)
    {
        cout << s << endl;
    }
}