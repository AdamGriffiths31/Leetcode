#include <vector>
#include <iostream>

using namespace std;

int findJudge(int n, vector<vector<int>> &trust)
{
    vector<int> incoming(n + 1, 0); 
    vector<int> outgoing(n + 1, 0); 
    for (auto i : trust)
    {
        outgoing[i[0]]++;
        incoming[i[1]]++;
    }

    for (int i = 1; i <= n; i++)
    {
        if (incoming[i] == n - 1 && outgoing[i] == 0)
        {
            return i;
        }
    }

    return -1;
}

int main()
{
    vector<vector<int>> trust = {{1, 2}};
    cout << findJudge(2, trust) << endl;
    return 0;
}